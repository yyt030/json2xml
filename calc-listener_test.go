package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"c2j/parser/calc"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct {
	*calc.BaseCalcListener
	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(c *calc.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case calc.CalcParserMUL:
		l.push(left * right)
	case calc.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *calc.AddSubContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case calc.CalcParserADD:
		l.push(left + right)
	case calc.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *calc.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err)
	}

	l.push(i)
}

func calcInter(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the lexer
	lexer := calc.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := calc.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()
}

func TestCalcListener(t *testing.T) {
	log.Println("result :", calcInter("1+2*3+10"))
}
