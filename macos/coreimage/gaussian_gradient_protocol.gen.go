// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a Gaussian gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient?language=objc
type PGaussianGradient interface {
	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetColor0(value Color)
	HasSetColor0() bool

	// optional
	Color0() Color
	HasColor0() bool

	// optional
	SetRadius(value float32)
	HasSetRadius() bool

	// optional
	Radius() float32
	HasRadius() bool

	// optional
	SetColor1(value Color)
	HasSetColor1() bool

	// optional
	Color1() Color
	HasColor1() bool
}

// ensure impl type implements protocol interface
var _ PGaussianGradient = (*GaussianGradientObject)(nil)

// A concrete type for the [PGaussianGradient] protocol.
type GaussianGradientObject struct {
	objc.Object
}

func (g_ GaussianGradientObject) HasSetCenter() bool {
	return g_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228467-center?language=objc
func (g_ GaussianGradientObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setCenter:"), value)
}

func (g_ GaussianGradientObject) HasCenter() bool {
	return g_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228467-center?language=objc
func (g_ GaussianGradientObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("center"))
	return rv
}

func (g_ GaussianGradientObject) HasSetColor0() bool {
	return g_.RespondsToSelector(objc.Sel("setColor0:"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228468-color0?language=objc
func (g_ GaussianGradientObject) SetColor0(value Color) {
	objc.Call[objc.Void](g_, objc.Sel("setColor0:"), value)
}

func (g_ GaussianGradientObject) HasColor0() bool {
	return g_.RespondsToSelector(objc.Sel("color0"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228468-color0?language=objc
func (g_ GaussianGradientObject) Color0() Color {
	rv := objc.Call[Color](g_, objc.Sel("color0"))
	return rv
}

func (g_ GaussianGradientObject) HasSetRadius() bool {
	return g_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the Gaussian distribution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228470-radius?language=objc
func (g_ GaussianGradientObject) SetRadius(value float32) {
	objc.Call[objc.Void](g_, objc.Sel("setRadius:"), value)
}

func (g_ GaussianGradientObject) HasRadius() bool {
	return g_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the Gaussian distribution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228470-radius?language=objc
func (g_ GaussianGradientObject) Radius() float32 {
	rv := objc.Call[float32](g_, objc.Sel("radius"))
	return rv
}

func (g_ GaussianGradientObject) HasSetColor1() bool {
	return g_.RespondsToSelector(objc.Sel("setColor1:"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228469-color1?language=objc
func (g_ GaussianGradientObject) SetColor1(value Color) {
	objc.Call[objc.Void](g_, objc.Sel("setColor1:"), value)
}

func (g_ GaussianGradientObject) HasColor1() bool {
	return g_.RespondsToSelector(objc.Sel("color1"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228469-color1?language=objc
func (g_ GaussianGradientObject) Color1() Color {
	rv := objc.Call[Color](g_, objc.Sel("color1"))
	return rv
}
