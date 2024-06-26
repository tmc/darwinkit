// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ExistsCommand] class.
var ExistsCommandClass = _ExistsCommandClass{objc.GetClass("NSExistsCommand")}

type _ExistsCommandClass struct {
	objc.Class
}

// An interface definition for the [ExistsCommand] class.
type IExistsCommand interface {
	IScriptCommand
}

// A command that determines whether a scriptable object exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexistscommand?language=objc
type ExistsCommand struct {
	ScriptCommand
}

func ExistsCommandFrom(ptr unsafe.Pointer) ExistsCommand {
	return ExistsCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (ec _ExistsCommandClass) Alloc() ExistsCommand {
	rv := objc.Call[ExistsCommand](ec, objc.Sel("alloc"))
	return rv
}

func (ec _ExistsCommandClass) New() ExistsCommand {
	rv := objc.Call[ExistsCommand](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExistsCommand() ExistsCommand {
	return ExistsCommandClass.New()
}

func (e_ ExistsCommand) Init() ExistsCommand {
	rv := objc.Call[ExistsCommand](e_, objc.Sel("init"))
	return rv
}

func (e_ ExistsCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) ExistsCommand {
	rv := objc.Call[ExistsCommand](e_, objc.Sel("initWithCommandDescription:"), commandDef)
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func NewExistsCommandWithCommandDescription(commandDef IScriptCommandDescription) ExistsCommand {
	instance := ExistsCommandClass.Alloc().InitWithCommandDescription(commandDef)
	instance.Autorelease()
	return instance
}
