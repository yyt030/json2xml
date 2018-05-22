### 什么是antlr
[antlr](http://www.antlr.org/)(ANother Tool for Language Recognition)是一个强大的语法分析器生成工具，它可用于读取，处理，执行和翻译结构化的文本和二进制文件。目前，该工具广泛应用于学术和工业生产领域，同时也是众多语言，工具和框架的基础。
今天我们就用这个工具实现一个go语言版的json2xml转换器；

### antlr的作用
关于一门语言的语法描述叫做grammar， 该工具能够为该语言生成语法解析器，并自动建立语法分析数AST，同时antlr也能自动生成数的遍历器，极大地降低手工coding 语法解析器的成本；

### 实践开始
言归正传，拿json2xml为例，实现一个工具；

### 安装
以macOS为例
```bash
brew install antlr
```
### 编辑json语言解析语法
```G4
// Derived from http://json.org
grammar JSON;

json:   object
    |   array
    ;

object
    :   '{' pair (',' pair)* '}'    # AnObject
    |   '{' '}'                     # NullObject
    ;

array
    :   '[' value (',' value)* ']'  # ArrayOfValues
    |   '[' ']'                     # NullArray
    ;

pair:   STRING ':' value ;

value
    :   STRING		# String
    |   NUMBER		# Atom
    |   object  	# ObjectValue
    |   array  		# ArrayValue
    |   'true'		# Atom
    |   'false'		# Atom
    |   'null'		# Atom
    ;

LCURLY : '{' ;
LBRACK : '[' ;
STRING :  '"' (ESC | ~["\\])* '"' ;
fragment ESC :   '\\' (["\\/bfnrt] | UNICODE) ;
fragment UNICODE : 'u' HEX HEX HEX HEX ;
fragment HEX : [0-9a-fA-F] ;
NUMBER
    :   '-'? INT '.' INT EXP?   // 1.35, 1.35E-9, 0.3, -4.5
    |   '-'? INT EXP            // 1e10 -3e4
    |   '-'? INT                // -3, 45
    ;
fragment INT :   '0' | '1'..'9' '0'..'9'* ; // no leading zeros
fragment EXP :   [Ee] [+\-]? INT ; // \- since - means "range" inside [...]
WS  :   [ \t\n\r]+ -> skip ;

```
上面是依照antlr4的语法格式编辑的文件
- antlr4文件语法也比较简单：
    - 以grammar关键字开头，名字与文件名相匹配
    - 语法分析器的规则必须以小写的字母开头
    - 词法分析器的规则必须用大写开头
    - | 管道符号分割同一语言规则的若干备选分支，使用圆括号把一些符号组成子规则。
- 涉及到的几个专有名词：
    - 语言： 语言即一个有效语句的集合，语句由词组组成，词组由子词组组成，一次循环类推；
    - 语法： 语法定义语言的语意规则， 语法中每条规则定义一种词组结构；
    - 语法分析树： 以树状的形式代表的语法的层次结构；根结点对应语法规则的名字，叶子节点代表语句中的符号或者词法符号。
    - 词法分析器： 将输入的字符序列分解成一系列的词法符号。一个词法分析器负责分析词法；
    - 语法分析器： 检查语句结构是否符合语法规范或者是否合法。分析的过程类似走迷宫，一般都是通过对比匹配完成。
    - 自顶向下语法分析器： 是语法分析器的一种实现，每条规则都对应语法分析器中的一个函数；
    - 前向预测： 语法分析器使用前向预测来进行决策判断，具体指将输入的符号于每个备选分支的起始字符进行比较；

### 生成解析基础代码
```bash
# antlr4 -Dlanguage=Go -package json2xml JSON.g4
```
> 使用antlr生成目标语言为Go， package名为json2xml的基础代码

生成的文件如下：
```
$ tree
├── JSON.g4
├── JSON.interp             # 语法解析中间文件
├── JSON.tokens             # 语法分析tokens流文件
├── JSONLexer.interp        # 词法分析中间文件
├── JSONLexer.tokens        # 词法分析tokens流文件
├── json_base_listener.go   # 默认是listener模式文件
├── json_lexer.go           # 词法分析器
├── json_listener.go        # 抽象listener接口文件
├── json_parser.go          # parser解析器文件

```
### 实现解析器（listener例子）
```Go
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
	xml map[antlr.Tree]string
}

func NewJ2xConvert() *j2xConvert {
	return &j2xConvert{
		&json2xml.BaseJSONListener{},
		make(map[antlr.Tree]string),
	}
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

```
> 上面代码比较简单，看注释就好；

一般流程如下：
- 新建输入流
- 新建词法分析器
- 生成token流，存储词法分析器生成的词法符号tokens
- 新建语法分析器parser，处理tokens
- 然后针对语法规则，开始语法分析
- 最后通过默认提供的Walker，进行AST的遍历

其中针对中间生成的参数和结果如何存放？好办，直接定义个map，map键以Tree存放；
```bash
xml map[antlr.Tree]string
```
### Listener和Visitor
antlr生成的代码有两种默认，默认是listener实现，要生成visitor，需要另加参数-visitor。
这两种机制的区别在于，监听器的方法会被antlr提供的遍历器对象自动调用，而visitor模式的方法中，必须显示调用visit方法来访问子节点。如果忘记调用的话，对应的子树就不会被访问。

### 总结
antlr是一个强大的工具，能让常见的语法解析工作事半功倍，效率极高。同时，该工具使语法分析过程和程序本身高度分离，提供足够的灵活性和可操控性。