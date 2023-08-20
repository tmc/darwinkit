package typing

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*FunctionType)(nil)

// FunctionType Objective-c interface type
type FunctionType struct {
	Name   string          // objc type name
	GName  string          // Go name, usually is objc type name without prefix 'NS'
	Module *modules.Module // object-c module
}

var Function = &FunctionType{
	Name:   "Function",
	GName:  "Function",
	Module: modules.Get("objc"),
}

func (c *FunctionType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/macdriver/macos/" + c.Module.Package)
}

func (c *FunctionType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	if receiveFromObjc {
		return FullGoName(*c.Module, c.GoStructName(), *currentModule)
	} else {
		return FullGoName(*c.Module, c.GoInterfaceName(), *currentModule)
	}
}

func (c *FunctionType) ObjcName() string {
	return c.Name + "*"
}

func (c *FunctionType) DeclareModule() *modules.Module {
	return c.Module
}

// GoInterfaceName return the go wrapper interface name
func (c *FunctionType) GoInterfaceName() string {
	return "I" + c.GName
}

// GoStructName return the go wrapper struct name
func (c *FunctionType) GoStructName() string {
	return c.GName
}
