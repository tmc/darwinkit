// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram?language=objc
type PAreaHistogram interface {
	// optional
	SetScale(value float32)
	HasSetScale() bool

	// optional
	Scale() float32
	HasScale() bool

	// optional
	SetCount(value int)
	HasSetCount() bool

	// optional
	Count() int
	HasCount() bool
}

// ensure impl type implements protocol interface
var _ PAreaHistogram = (*AreaHistogramObject)(nil)

// A concrete type for the [PAreaHistogram] protocol.
type AreaHistogramObject struct {
	objc.Object
}

func (a_ AreaHistogramObject) HasSetScale() bool {
	return a_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547093-scale?language=objc
func (a_ AreaHistogramObject) SetScale(value float32) {
	objc.Call[objc.Void](a_, objc.Sel("setScale:"), value)
}

func (a_ AreaHistogramObject) HasScale() bool {
	return a_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547093-scale?language=objc
func (a_ AreaHistogramObject) Scale() float32 {
	rv := objc.Call[float32](a_, objc.Sel("scale"))
	return rv
}

func (a_ AreaHistogramObject) HasSetCount() bool {
	return a_.RespondsToSelector(objc.Sel("setCount:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547092-count?language=objc
func (a_ AreaHistogramObject) SetCount(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setCount:"), value)
}

func (a_ AreaHistogramObject) HasCount() bool {
	return a_.RespondsToSelector(objc.Sel("count"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547092-count?language=objc
func (a_ AreaHistogramObject) Count() int {
	rv := objc.Call[int](a_, objc.Sel("count"))
	return rv
}
