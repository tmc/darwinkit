// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Attribute] class.
var AttributeClass = _AttributeClass{objc.GetClass("MTLAttribute")}

type _AttributeClass struct {
	objc.Class
}

// An interface definition for the [Attribute] class.
type IAttribute interface {
	objc.IObject
	AttributeIndex() uint
	IsPatchData() bool
	IsPatchControlPointData() bool
	Name() string
	IsActive() bool
	AttributeType() DataType
}

// An object that describes an attribute defined in the stage-in argument for a shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute?language=objc
type Attribute struct {
	objc.Object
}

func AttributeFrom(ptr unsafe.Pointer) Attribute {
	return Attribute{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AttributeClass) Alloc() Attribute {
	rv := objc.Call[Attribute](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AttributeClass) New() Attribute {
	rv := objc.Call[Attribute](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAttribute() Attribute {
	return AttributeClass.New()
}

func (a_ Attribute) Init() Attribute {
	rv := objc.Call[Attribute](a_, objc.Sel("init"))
	return rv
}

// The index of the attribute, as declared in Metal shader source code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/2097158-attributeindex?language=objc
func (a_ Attribute) AttributeIndex() uint {
	rv := objc.Call[uint](a_, objc.Sel("attributeIndex"))
	return rv
}

// A Boolean value that indicates whether the attribute represents tessellation patch data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/2097157-patchdata?language=objc
func (a_ Attribute) IsPatchData() bool {
	rv := objc.Call[bool](a_, objc.Sel("isPatchData"))
	return rv
}

// A Boolean value that indicates whether the attribute represents control point data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/2097156-patchcontrolpointdata?language=objc
func (a_ Attribute) IsPatchControlPointData() bool {
	rv := objc.Call[bool](a_, objc.Sel("isPatchControlPointData"))
	return rv
}

// The name of the attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/2097161-name?language=objc
func (a_ Attribute) Name() string {
	rv := objc.Call[string](a_, objc.Sel("name"))
	return rv
}

// A Boolean value that indicates whether the attribute is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/2097160-active?language=objc
func (a_ Attribute) IsActive() bool {
	rv := objc.Call[bool](a_, objc.Sel("isActive"))
	return rv
}

// The data type for the attribute, as declared in Metal shader source code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/2097155-attributetype?language=objc
func (a_ Attribute) AttributeType() DataType {
	rv := objc.Call[DataType](a_, objc.Sel("attributeType"))
	return rv
}
