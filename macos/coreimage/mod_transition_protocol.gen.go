// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a mod transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition?language=objc
type PModTransition interface {
	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetCompression(value float64)
	HasSetCompression() bool

	// optional
	Compression() float64
	HasCompression() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool
}

// ensure impl type implements protocol interface
var _ PModTransition = (*ModTransitionObject)(nil)

// A concrete type for the [PModTransition] protocol.
type ModTransitionObject struct {
	objc.Object
}

func (m_ ModTransitionObject) HasSetCenter() bool {
	return m_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228570-center?language=objc
func (m_ ModTransitionObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](m_, objc.Sel("setCenter:"), value)
}

func (m_ ModTransitionObject) HasCenter() bool {
	return m_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228570-center?language=objc
func (m_ ModTransitionObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](m_, objc.Sel("center"))
	return rv
}

func (m_ ModTransitionObject) HasSetRadius() bool {
	return m_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the undistorted mod holes in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228572-radius?language=objc
func (m_ ModTransitionObject) SetRadius(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setRadius:"), value)
}

func (m_ ModTransitionObject) HasRadius() bool {
	return m_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the undistorted mod holes in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228572-radius?language=objc
func (m_ ModTransitionObject) Radius() float64 {
	rv := objc.Call[float64](m_, objc.Sel("radius"))
	return rv
}

func (m_ ModTransitionObject) HasSetCompression() bool {
	return m_.RespondsToSelector(objc.Sel("setCompression:"))
}

// The amount of stretching applied to the mod hole pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228571-compression?language=objc
func (m_ ModTransitionObject) SetCompression(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setCompression:"), value)
}

func (m_ ModTransitionObject) HasCompression() bool {
	return m_.RespondsToSelector(objc.Sel("compression"))
}

// The amount of stretching applied to the mod hole pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228571-compression?language=objc
func (m_ ModTransitionObject) Compression() float64 {
	rv := objc.Call[float64](m_, objc.Sel("compression"))
	return rv
}

func (m_ ModTransitionObject) HasSetAngle() bool {
	return m_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the mod hole pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228569-angle?language=objc
func (m_ ModTransitionObject) SetAngle(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setAngle:"), value)
}

func (m_ ModTransitionObject) HasAngle() bool {
	return m_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the mod hole pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimodtransition/3228569-angle?language=objc
func (m_ ModTransitionObject) Angle() float64 {
	rv := objc.Call[float64](m_, objc.Sel("angle"))
	return rv
}