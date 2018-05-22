package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"c2j/parser/json2xml"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type j2xConvert struct {
	*json2xml.BaseJSONListener
	xml   map[antlr.Tree]string
	stack []antlr.ParseTree
}

func NewJ2xConvert() *j2xConvert {
	return &j2xConvert{
		&json2xml.BaseJSONListener{},
		make(map[antlr.Tree]string),
		make([]antlr.ParseTree, 0),
	}
}

func (j *j2xConvert) push(ctx antlr.ParseTree) {
	j.stack = append(j.stack, ctx)
}

func (j *j2xConvert) pop() antlr.ParseTree {
	value := j.stack[len(j.stack)-1]
	j.stack = j.stack[:len(j.stack)-1]
	return value
}

func (j *j2xConvert) setXML(ctx antlr.Tree, s string) {
	j.xml[ctx] = s
}

func (j *j2xConvert) getXML(ctx antlr.Tree) string {
	return j.xml[ctx]
}

// j2xConvert methods
func (j *j2xConvert) ExitJson(ctx *json2xml.JsonContext) {
	j.setXML(ctx, j.getXML(ctx.GetChild(0)));
}

func (j *j2xConvert) stripQuotes(s string) string {
	if s == "" || ! strings.Contains(s, "\"") {
		return s
	}
	return s[1 : len(s)-1]
}

func (j *j2xConvert) ExitAnObject(ctx *json2xml.AnObjectContext) {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for _, p := range ctx.AllPair() {
		sb.WriteString(j.getXML(p))
	}
	j.setXML(ctx, sb.String())
}

func (j *j2xConvert) ExitNullObject(ctx *json2xml.NullObjectContext) {
	j.setXML(ctx, "")
}

func (j *j2xConvert) ExitArrayOfValues(ctx *json2xml.ArrayOfValuesContext) {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for _, p := range ctx.AllValue() {
		//log.Println(">>>", p.GetText(), j.getXML(p))
		//s := fmt.Sprintf("<element>%s<elment>", )
		sb.WriteString("<element>")
		sb.WriteString(j.getXML(p))
		sb.WriteString("</element>")
		sb.WriteString("\n")
	}
	j.setXML(ctx, sb.String())
}

func (j *j2xConvert) ExitNullArray(ctx *json2xml.NullArrayContext) {
	j.setXML(ctx, "")
}

func (j *j2xConvert) ExitPair(ctx *json2xml.PairContext) {
	tag := j.stripQuotes(ctx.STRING().GetText())
	v := ctx.Value()
	r := fmt.Sprintf("<%s>%s</%s>\n", tag, j.getXML(v), tag)
	j.setXML(ctx, r)
}

func (j *j2xConvert) ExitObjectValue(ctx *json2xml.ObjectValueContext) {
	j.setXML(ctx, j.getXML(ctx.Object()))
	//log.Println("ExitObjectValue:", ctx.Object())
}

func (j *j2xConvert) ExitArrayValue(ctx *json2xml.ArrayValueContext) {
	j.setXML(ctx, j.getXML(ctx.Array()))
}

func (j *j2xConvert) ExitAtom(ctx *json2xml.AtomContext) {
	j.setXML(ctx, ctx.GetText())
}

func (j *j2xConvert) ExitString(ctx *json2xml.StringContext) {
	j.setXML(ctx, j.stripQuotes(ctx.GetText()))
}

func TestJSON2XMLVisitor(t *testing.T) {
	f, err := os.Open("testdata/json2xml/t.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Setup the input
	is := antlr.NewInputStream(string(content))

	// Create lexter
	lexer := json2xml.NewJSONLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// Create parser and tree
	p := json2xml.NewJSONParser(stream)
	p.BuildParseTrees = true
	tree := p.Json()

	// Finally AST tree
	j2x := NewJ2xConvert()
	antlr.ParseTreeWalkerDefault.Walk(j2x, tree)
	log.Println(j2x.getXML(tree))
}
