// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a fourfold rotated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile?language=objc
type PFourfoldRotatedTile interface {
	// optional
	SetAngle(value float32)
	HasSetAngle() bool

	// optional
	Angle() float32
	HasAngle() bool

	// optional
	SetWidth(value float32)
	HasSetWidth() bool

	// optional
	Width() float32
	HasWidth() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PFourfoldRotatedTile = (*FourfoldRotatedTileObject)(nil)

// A concrete type for the [PFourfoldRotatedTile] protocol.
type FourfoldRotatedTileObject struct {
	objc.Object
}

func (f_ FourfoldRotatedTileObject) HasSetAngle() bool {
	return f_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228450-angle?language=objc
func (f_ FourfoldRotatedTileObject) SetAngle(value float32) {
	objc.Call[objc.Void](f_, objc.Sel("setAngle:"), value)
}

func (f_ FourfoldRotatedTileObject) HasAngle() bool {
	return f_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228450-angle?language=objc
func (f_ FourfoldRotatedTileObject) Angle() float32 {
	rv := objc.Call[float32](f_, objc.Sel("angle"))
	return rv
}

func (f_ FourfoldRotatedTileObject) HasSetWidth() bool {
	return f_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228453-width?language=objc
func (f_ FourfoldRotatedTileObject) SetWidth(value float32) {
	objc.Call[objc.Void](f_, objc.Sel("setWidth:"), value)
}

func (f_ FourfoldRotatedTileObject) HasWidth() bool {
	return f_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228453-width?language=objc
func (f_ FourfoldRotatedTileObject) Width() float32 {
	rv := objc.Call[float32](f_, objc.Sel("width"))
	return rv
}

func (f_ FourfoldRotatedTileObject) HasSetCenter() bool {
	return f_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228451-center?language=objc
func (f_ FourfoldRotatedTileObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setCenter:"), value)
}

func (f_ FourfoldRotatedTileObject) HasCenter() bool {
	return f_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228451-center?language=objc
func (f_ FourfoldRotatedTileObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("center"))
	return rv
}

func (f_ FourfoldRotatedTileObject) HasSetInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228452-inputimage?language=objc
func (f_ FourfoldRotatedTileObject) SetInputImage(value Image) {
	objc.Call[objc.Void](f_, objc.Sel("setInputImage:"), value)
}

func (f_ FourfoldRotatedTileObject) HasInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldrotatedtile/3228452-inputimage?language=objc
func (f_ FourfoldRotatedTileObject) InputImage() Image {
	rv := objc.Call[Image](f_, objc.Sel("inputImage"))
	return rv
}
