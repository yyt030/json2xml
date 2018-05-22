// Code generated from JSON.g4 by ANTLR 4.7.1. DO NOT EDIT.

package json2xml // JSON
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJSONListener is a complete listener for a parse tree produced by JSONParser.
type BaseJSONListener struct{}

var _ JSONListener = &BaseJSONListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJSONListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJSONListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJSONListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJSONListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (s *BaseJSONListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseJSONListener) ExitJson(ctx *JsonContext) {}

// EnterAnObject is called when production AnObject is entered.
func (s *BaseJSONListener) EnterAnObject(ctx *AnObjectContext) {}

// ExitAnObject is called when production AnObject is exited.
func (s *BaseJSONListener) ExitAnObject(ctx *AnObjectContext) {}

// EnterNullObject is called when production NullObject is entered.
func (s *BaseJSONListener) EnterNullObject(ctx *NullObjectContext) {}

// ExitNullObject is called when production NullObject is exited.
func (s *BaseJSONListener) ExitNullObject(ctx *NullObjectContext) {}

// EnterArrayOfValues is called when production ArrayOfValues is entered.
func (s *BaseJSONListener) EnterArrayOfValues(ctx *ArrayOfValuesContext) {}

// ExitArrayOfValues is called when production ArrayOfValues is exited.
func (s *BaseJSONListener) ExitArrayOfValues(ctx *ArrayOfValuesContext) {}

// EnterNullArray is called when production NullArray is entered.
func (s *BaseJSONListener) EnterNullArray(ctx *NullArrayContext) {}

// ExitNullArray is called when production NullArray is exited.
func (s *BaseJSONListener) ExitNullArray(ctx *NullArrayContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseJSONListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseJSONListener) ExitPair(ctx *PairContext) {}

// EnterString is called when production String is entered.
func (s *BaseJSONListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseJSONListener) ExitString(ctx *StringContext) {}

// EnterAtom is called when production Atom is entered.
func (s *BaseJSONListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production Atom is exited.
func (s *BaseJSONListener) ExitAtom(ctx *AtomContext) {}

// EnterObjectValue is called when production ObjectValue is entered.
func (s *BaseJSONListener) EnterObjectValue(ctx *ObjectValueContext) {}

// ExitObjectValue is called when production ObjectValue is exited.
func (s *BaseJSONListener) ExitObjectValue(ctx *ObjectValueContext) {}

// EnterArrayValue is called when production ArrayValue is entered.
func (s *BaseJSONListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production ArrayValue is exited.
func (s *BaseJSONListener) ExitArrayValue(ctx *ArrayValueContext) {}
