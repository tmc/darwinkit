// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NNDefaultPadding] class.
var NNDefaultPaddingClass = _NNDefaultPaddingClass{objc.GetClass("MPSNNDefaultPadding")}

type _NNDefaultPaddingClass struct {
	objc.Class
}

// An interface definition for the [NNDefaultPadding] class.
type INNDefaultPadding interface {
	objc.IObject
	Label() string
}

// A class that provides predefined padding policies for common tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding?language=objc
type NNDefaultPadding struct {
	objc.Object
}

func NNDefaultPaddingFrom(ptr unsafe.Pointer) NNDefaultPadding {
	return NNDefaultPadding{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NNDefaultPaddingClass) PaddingForTensorflowAveragePoolingValidOnly() NNDefaultPadding {
	rv := objc.Call[NNDefaultPadding](nc, objc.Sel("paddingForTensorflowAveragePoolingValidOnly"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2947962-paddingfortensorflowaveragepooli?language=objc
func NNDefaultPadding_PaddingForTensorflowAveragePoolingValidOnly() NNDefaultPadding {
	return NNDefaultPaddingClass.PaddingForTensorflowAveragePoolingValidOnly()
}

func (nc _NNDefaultPaddingClass) PaddingWithMethod(method NNPaddingMethod) NNDefaultPadding {
	rv := objc.Call[NNDefaultPadding](nc, objc.Sel("paddingWithMethod:"), method)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2867160-paddingwithmethod?language=objc
func NNDefaultPadding_PaddingWithMethod(method NNPaddingMethod) NNDefaultPadding {
	return NNDefaultPaddingClass.PaddingWithMethod(method)
}

func (nc _NNDefaultPaddingClass) PaddingForTensorflowAveragePooling() NNDefaultPadding {
	rv := objc.Call[NNDefaultPadding](nc, objc.Sel("paddingForTensorflowAveragePooling"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2867164-paddingfortensorflowaveragepooli?language=objc
func NNDefaultPadding_PaddingForTensorflowAveragePooling() NNDefaultPadding {
	return NNDefaultPaddingClass.PaddingForTensorflowAveragePooling()
}

func (nc _NNDefaultPaddingClass) Alloc() NNDefaultPadding {
	rv := objc.Call[NNDefaultPadding](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNDefaultPaddingClass) New() NNDefaultPadding {
	rv := objc.Call[NNDefaultPadding](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNDefaultPadding() NNDefaultPadding {
	return NNDefaultPaddingClass.New()
}

func (n_ NNDefaultPadding) Init() NNDefaultPadding {
	rv := objc.Call[NNDefaultPadding](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2889871-label?language=objc
func (n_ NNDefaultPadding) Label() string {
	rv := objc.Call[string](n_, objc.Sel("label"))
	return rv
}
