// Code generated by DarwinKit. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ParameterDescription] class.
var ParameterDescriptionClass = _ParameterDescriptionClass{objc.GetClass("MLParameterDescription")}

type _ParameterDescriptionClass struct {
	objc.Class
}

// An interface definition for the [ParameterDescription] class.
type IParameterDescription interface {
	objc.IObject
	NumericConstraint() NumericConstraint
	DefaultValue() objc.Object
	Key() ParameterKey
}

// A description of a model parameter that includes a default value and a constraint, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription?language=objc
type ParameterDescription struct {
	objc.Object
}

func ParameterDescriptionFrom(ptr unsafe.Pointer) ParameterDescription {
	return ParameterDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _ParameterDescriptionClass) Alloc() ParameterDescription {
	rv := objc.Call[ParameterDescription](pc, objc.Sel("alloc"))
	return rv
}

func (pc _ParameterDescriptionClass) New() ParameterDescription {
	rv := objc.Call[ParameterDescription](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewParameterDescription() ParameterDescription {
	return ParameterDescriptionClass.New()
}

func (p_ ParameterDescription) Init() ParameterDescription {
	rv := objc.Call[ParameterDescription](p_, objc.Sel("init"))
	return rv
}

// The constraints of this paramter description value, if and only if the value is numerical. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription/3180070-numericconstraint?language=objc
func (p_ ParameterDescription) NumericConstraint() NumericConstraint {
	rv := objc.Call[NumericConstraint](p_, objc.Sel("numericConstraint"))
	return rv
}

// The default value for the parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription/3180068-defaultvalue?language=objc
func (p_ ParameterDescription) DefaultValue() objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("defaultValue"))
	return rv
}

// The key for this parameter description value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription/3180069-key?language=objc
func (p_ ParameterDescription) Key() ParameterKey {
	rv := objc.Call[ParameterKey](p_, objc.Sel("key"))
	return rv
}
