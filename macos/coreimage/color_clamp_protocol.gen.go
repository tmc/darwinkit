// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a color clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp?language=objc
type PColorClamp interface {
	// optional
	SetMinComponents(value Vector)
	HasSetMinComponents() bool

	// optional
	MinComponents() Vector
	HasMinComponents() bool

	// optional
	SetMaxComponents(value Vector)
	HasSetMaxComponents() bool

	// optional
	MaxComponents() Vector
	HasMaxComponents() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PColorClamp = (*ColorClampObject)(nil)

// A concrete type for the [PColorClamp] protocol.
type ColorClampObject struct {
	objc.Object
}

func (c_ ColorClampObject) HasSetMinComponents() bool {
	return c_.RespondsToSelector(objc.Sel("setMinComponents:"))
}

// A vector containing the lower clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228122-mincomponents?language=objc
func (c_ ColorClampObject) SetMinComponents(value Vector) {
	objc.Call[objc.Void](c_, objc.Sel("setMinComponents:"), value)
}

func (c_ ColorClampObject) HasMinComponents() bool {
	return c_.RespondsToSelector(objc.Sel("minComponents"))
}

// A vector containing the lower clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228122-mincomponents?language=objc
func (c_ ColorClampObject) MinComponents() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("minComponents"))
	return rv
}

func (c_ ColorClampObject) HasSetMaxComponents() bool {
	return c_.RespondsToSelector(objc.Sel("setMaxComponents:"))
}

// A vector containing the higher clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228121-maxcomponents?language=objc
func (c_ ColorClampObject) SetMaxComponents(value Vector) {
	objc.Call[objc.Void](c_, objc.Sel("setMaxComponents:"), value)
}

func (c_ ColorClampObject) HasMaxComponents() bool {
	return c_.RespondsToSelector(objc.Sel("maxComponents"))
}

// A vector containing the higher clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228121-maxcomponents?language=objc
func (c_ ColorClampObject) MaxComponents() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("maxComponents"))
	return rv
}

func (c_ ColorClampObject) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228120-inputimage?language=objc
func (c_ ColorClampObject) SetInputImage(value Image) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), value)
}

func (c_ ColorClampObject) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228120-inputimage?language=objc
func (c_ ColorClampObject) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}
