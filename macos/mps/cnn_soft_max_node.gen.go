// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNSoftMaxNode] class.
var CNNSoftMaxNodeClass = _CNNSoftMaxNodeClass{objc.GetClass("MPSCNNSoftMaxNode")}

type _CNNSoftMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNSoftMaxNode] class.
type ICNNSoftMaxNode interface {
	INNFilterNode
}

// A representation of a softmax filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxnode?language=objc
type CNNSoftMaxNode struct {
	NNFilterNode
}

func CNNSoftMaxNodeFrom(ptr unsafe.Pointer) CNNSoftMaxNode {
	return CNNSoftMaxNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (c_ CNNSoftMaxNode) InitWithSource(sourceNode INNImageNode) CNNSoftMaxNode {
	rv := objc.Call[CNNSoftMaxNode](c_, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxnode/2866408-initwithsource?language=objc
func NewCNNSoftMaxNodeWithSource(sourceNode INNImageNode) CNNSoftMaxNode {
	instance := CNNSoftMaxNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNSoftMaxNodeClass) NodeWithSource(sourceNode INNImageNode) CNNSoftMaxNode {
	rv := objc.Call[CNNSoftMaxNode](cc, objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxnode/2866455-nodewithsource?language=objc
func CNNSoftMaxNode_NodeWithSource(sourceNode INNImageNode) CNNSoftMaxNode {
	return CNNSoftMaxNodeClass.NodeWithSource(sourceNode)
}

func (cc _CNNSoftMaxNodeClass) Alloc() CNNSoftMaxNode {
	rv := objc.Call[CNNSoftMaxNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNSoftMaxNodeClass) New() CNNSoftMaxNode {
	rv := objc.Call[CNNSoftMaxNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSoftMaxNode() CNNSoftMaxNode {
	return CNNSoftMaxNodeClass.New()
}

func (c_ CNNSoftMaxNode) Init() CNNSoftMaxNode {
	rv := objc.Call[CNNSoftMaxNode](c_, objc.Sel("init"))
	return rv
}
