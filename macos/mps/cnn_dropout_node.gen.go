// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDropoutNode] class.
var CNNDropoutNodeClass = _CNNDropoutNodeClass{objc.GetClass("MPSCNNDropoutNode")}

type _CNNDropoutNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNDropoutNode] class.
type ICNNDropoutNode interface {
	INNFilterNode
	KeepProbability() float64
	MaskStrideInPixels() metal.Size
	Seed() uint
}

// A representation of a dropout filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode?language=objc
type CNNDropoutNode struct {
	NNFilterNode
}

func CNNDropoutNodeFrom(ptr unsafe.Pointer) CNNDropoutNode {
	return CNNDropoutNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNDropoutNodeClass) NodeWithSourceKeepProbability(source INNImageNode, keepProbability float64) CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](cc, objc.Sel("nodeWithSource:keepProbability:"), objc.Ptr(source), keepProbability)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947975-nodewithsource?language=objc
func CNNDropoutNode_NodeWithSourceKeepProbability(source INNImageNode, keepProbability float64) CNNDropoutNode {
	return CNNDropoutNodeClass.NodeWithSourceKeepProbability(source, keepProbability)
}

func (c_ CNNDropoutNode) InitWithSourceKeepProbability(source INNImageNode, keepProbability float64) CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](c_, objc.Sel("initWithSource:keepProbability:"), objc.Ptr(source), keepProbability)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2948000-initwithsource?language=objc
func NewCNNDropoutNodeWithSourceKeepProbability(source INNImageNode, keepProbability float64) CNNDropoutNode {
	instance := CNNDropoutNodeClass.Alloc().InitWithSourceKeepProbability(source, keepProbability)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutNode) InitWithSourceKeepProbabilitySeedMaskStrideInPixels(source INNImageNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](c_, objc.Sel("initWithSource:keepProbability:seed:maskStrideInPixels:"), objc.Ptr(source), keepProbability, seed, maskStrideInPixels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947990-initwithsource?language=objc
func NewCNNDropoutNodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source INNImageNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutNode {
	instance := CNNDropoutNodeClass.Alloc().InitWithSourceKeepProbabilitySeedMaskStrideInPixels(source, keepProbability, seed, maskStrideInPixels)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutNode) InitWithSource(source INNImageNode) CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](c_, objc.Sel("initWithSource:"), objc.Ptr(source))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947969-initwithsource?language=objc
func NewCNNDropoutNodeWithSource(source INNImageNode) CNNDropoutNode {
	instance := CNNDropoutNodeClass.Alloc().InitWithSource(source)
	instance.Autorelease()
	return instance
}

func (cc _CNNDropoutNodeClass) NodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source INNImageNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](cc, objc.Sel("nodeWithSource:keepProbability:seed:maskStrideInPixels:"), objc.Ptr(source), keepProbability, seed, maskStrideInPixels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947989-nodewithsource?language=objc
func CNNDropoutNode_NodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source INNImageNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutNode {
	return CNNDropoutNodeClass.NodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source, keepProbability, seed, maskStrideInPixels)
}

func (cc _CNNDropoutNodeClass) NodeWithSource(source INNImageNode) CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(source))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2948007-nodewithsource?language=objc
func CNNDropoutNode_NodeWithSource(source INNImageNode) CNNDropoutNode {
	return CNNDropoutNodeClass.NodeWithSource(source)
}

func (cc _CNNDropoutNodeClass) Alloc() CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNDropoutNodeClass) New() CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDropoutNode() CNNDropoutNode {
	return CNNDropoutNodeClass.New()
}

func (c_ CNNDropoutNode) Init() CNNDropoutNode {
	rv := objc.Call[CNNDropoutNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947982-keepprobability?language=objc
func (c_ CNNDropoutNode) KeepProbability() float64 {
	rv := objc.Call[float64](c_, objc.Sel("keepProbability"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947998-maskstrideinpixels?language=objc
func (c_ CNNDropoutNode) MaskStrideInPixels() metal.Size {
	rv := objc.Call[metal.Size](c_, objc.Sel("maskStrideInPixels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2948030-seed?language=objc
func (c_ CNNDropoutNode) Seed() uint {
	rv := objc.Call[uint](c_, objc.Sel("seed"))
	return rv
}