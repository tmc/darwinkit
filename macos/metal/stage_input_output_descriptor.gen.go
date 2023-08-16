// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StageInputOutputDescriptor] class.
var StageInputOutputDescriptorClass = _StageInputOutputDescriptorClass{objc.GetClass("MTLStageInputOutputDescriptor")}

type _StageInputOutputDescriptorClass struct {
	objc.Class
}

// An interface definition for the [StageInputOutputDescriptor] class.
type IStageInputOutputDescriptor interface {
	objc.IObject
	Reset()
	IndexType() IndexType
	SetIndexType(value IndexType)
	IndexBufferIndex() uint
	SetIndexBufferIndex(value uint)
	Layouts() BufferLayoutDescriptorArray
	Attributes() AttributeDescriptorArray
}

// An object that describes the input and output data of a function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor?language=objc
type StageInputOutputDescriptor struct {
	objc.Object
}

func StageInputOutputDescriptorFrom(ptr unsafe.Pointer) StageInputOutputDescriptor {
	return StageInputOutputDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StageInputOutputDescriptorClass) Alloc() StageInputOutputDescriptor {
	rv := objc.Call[StageInputOutputDescriptor](sc, objc.Sel("alloc"))
	return rv
}

func StageInputOutputDescriptor_Alloc() StageInputOutputDescriptor {
	return StageInputOutputDescriptorClass.Alloc()
}

func (sc _StageInputOutputDescriptorClass) New() StageInputOutputDescriptor {
	rv := objc.Call[StageInputOutputDescriptor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStageInputOutputDescriptor() StageInputOutputDescriptor {
	return StageInputOutputDescriptorClass.New()
}

func (s_ StageInputOutputDescriptor) Init() StageInputOutputDescriptor {
	rv := objc.Call[StageInputOutputDescriptor](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097294-stageinputoutputdescriptor?language=objc
func (sc _StageInputOutputDescriptorClass) StageInputOutputDescriptor() StageInputOutputDescriptor {
	rv := objc.Call[StageInputOutputDescriptor](sc, objc.Sel("stageInputOutputDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097294-stageinputoutputdescriptor?language=objc
func StageInputOutputDescriptor_StageInputOutputDescriptor() StageInputOutputDescriptor {
	return StageInputOutputDescriptorClass.StageInputOutputDescriptor()
}

// Resets the default state for the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097185-reset?language=objc
func (s_ StageInputOutputDescriptor) Reset() {
	objc.Call[objc.Void](s_, objc.Sel("reset"))
}

// The data type of the indices stored in the index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097184-indextype?language=objc
func (s_ StageInputOutputDescriptor) IndexType() IndexType {
	rv := objc.Call[IndexType](s_, objc.Sel("indexType"))
	return rv
}

// The data type of the indices stored in the index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097184-indextype?language=objc
func (s_ StageInputOutputDescriptor) SetIndexType(value IndexType) {
	objc.Call[objc.Void](s_, objc.Sel("setIndexType:"), value)
}

// The location of the index buffer for a compute function using indexed thread addressing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097237-indexbufferindex?language=objc
func (s_ StageInputOutputDescriptor) IndexBufferIndex() uint {
	rv := objc.Call[uint](s_, objc.Sel("indexBufferIndex"))
	return rv
}

// The location of the index buffer for a compute function using indexed thread addressing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097237-indexbufferindex?language=objc
func (s_ StageInputOutputDescriptor) SetIndexBufferIndex(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setIndexBufferIndex:"), value)
}

// An array that describes how data is fetched for the function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097202-layouts?language=objc
func (s_ StageInputOutputDescriptor) Layouts() BufferLayoutDescriptorArray {
	rv := objc.Call[BufferLayoutDescriptorArray](s_, objc.Sel("layouts"))
	return rv
}

// An array that describes where and how to fetch data for the function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinputoutputdescriptor/2097206-attributes?language=objc
func (s_ StageInputOutputDescriptor) Attributes() AttributeDescriptorArray {
	rv := objc.Call[AttributeDescriptorArray](s_, objc.Sel("attributes"))
	return rv
}