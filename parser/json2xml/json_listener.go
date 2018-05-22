// Code generated from JSON.g4 by ANTLR 4.7.1. DO NOT EDIT.

package json2xml // JSON
import "github.com/antlr/antlr4/runtime/Go/antlr"

// JSONListener is a complete listener for a parse tree produced by JSONParser.
type JSONListener interface {
	antlr.ParseTreeListener

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterAnObject is called when entering the AnObject production.
	EnterAnObject(c *AnObjectContext)

	// EnterNullObject is called when entering the NullObject production.
	EnterNullObject(c *NullObjectContext)

	// EnterArrayOfValues is called when entering the ArrayOfValues production.
	EnterArrayOfValues(c *ArrayOfValuesContext)

	// EnterNullArray is called when entering the NullArray production.
	EnterNullArray(c *NullArrayContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterAtom is called when entering the Atom production.
	EnterAtom(c *AtomContext)

	// EnterObjectValue is called when entering the ObjectValue production.
	EnterObjectValue(c *ObjectValueContext)

	// EnterArrayValue is called when entering the ArrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitAnObject is called when exiting the AnObject production.
	ExitAnObject(c *AnObjectContext)

	// ExitNullObject is called when exiting the NullObject production.
	ExitNullObject(c *NullObjectContext)

	// ExitArrayOfValues is called when exiting the ArrayOfValues production.
	ExitArrayOfValues(c *ArrayOfValuesContext)

	// ExitNullArray is called when exiting the NullArray production.
	ExitNullArray(c *NullArrayContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitAtom is called when exiting the Atom production.
	ExitAtom(c *AtomContext)

	// ExitObjectValue is called when exiting the ObjectValue production.
	ExitObjectValue(c *ObjectValueContext)

	// ExitArrayValue is called when exiting the ArrayValue production.
	ExitArrayValue(c *ArrayValueContext)
}
