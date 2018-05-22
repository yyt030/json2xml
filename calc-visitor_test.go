package main

import (
	"log"
	"testing"

	"c2j/parser/calc"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type EvalVisitor struct {
	*calc.BaseCalcVisitor
	//stack map[string]int
}

func NewEvalVisitor() *EvalVisitor {
	//visitor := new(EvalVisitor)
	//return visitor
	return &EvalVisitor{
		&calc.BaseCalcVisitor{},
	}
}

func (v *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *EvalVisitor) VisitStart(ctx *calc.StartContext) interface{} {
	log.Println("MMM", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *EvalVisitor) VisitNumber(ctx *calc.NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvalVisitor) VisitMulDiv(ctx *calc.MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvalVisitor) VisitAddSub(ctx *calc.AddSubContext) interface{} {
	log.Println("AddSub")
	return v.VisitChildren(ctx)
}

func TestEvalVisitor(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 + 3 * 2 +2")

	// Create lexer
	lexer := calc.NewCalcLexer(is)

	// Create tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create parser
	p := calc.NewCalcParser(stream)
	p.BuildParseTrees = true

	// Create ast tree
	tree := p.Start()
	log.Println("tree	:", tree.GetText())

	// Finilly create internal
	eval := NewEvalVisitor()
	log.Println("visitor:", eval)
	eval.Visit(tree)
}
