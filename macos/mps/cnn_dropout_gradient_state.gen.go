// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNDropoutGradientState] class.
var CNNDropoutGradientStateClass = _CNNDropoutGradientStateClass{objc.GetClass("MPSCNNDropoutGradientState")}

type _CNNDropoutGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNDropoutGradientState] class.
type ICNNDropoutGradientState interface {
	INNGradientState
	MaskData() []byte
}

// A class that stores the mask used by dropout and gradient dropout filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientstate?language=objc
type CNNDropoutGradientState struct {
	NNGradientState
}

func CNNDropoutGradientStateFrom(ptr unsafe.Pointer) CNNDropoutGradientState {
	return CNNDropoutGradientState{
		NNGradientState: NNGradientStateFrom(ptr),
	}
}

func (cc _CNNDropoutGradientStateClass) Alloc() CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNDropoutGradientStateClass) New() CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDropoutGradientState() CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.New()
}

func (c_ CNNDropoutGradientState) Init() CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNDropoutGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, resourceList)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNDropoutGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

func (c_ CNNDropoutGradientState) InitWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithDevice:textureDescriptor:"), po0, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942400-initwithdevice?language=objc
func NewCNNDropoutGradientStateWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithDeviceTextureDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutGradientState) InitWithResource(resource metal.PResource) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNDropoutGradientStateWithResource(resource metal.PResource) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNDropoutGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutGradientState) InitWithResources(resources []metal.PResource) CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNDropoutGradientStateWithResources(resources []metal.PResource) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (cc _CNNDropoutGradientStateClass) TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942391-temporarystatewithcommandbuffer?language=objc
func CNNDropoutGradientState_TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.TemporaryStateWithCommandBufferBufferSize(cmdBuf, bufferSize)
}

func (cc _CNNDropoutGradientStateClass) TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942393-temporarystatewithcommandbuffer?language=objc
func CNNDropoutGradientState_TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.TemporaryStateWithCommandBuffer(cmdBuf)
}

func (cc _CNNDropoutGradientStateClass) TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:textureDescriptor:"), po0, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942395-temporarystatewithcommandbuffer?language=objc
func CNNDropoutGradientState_TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf, descriptor)
}

func (c_ CNNDropoutGradientState) InitWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithDevice:resourceList:"), po0, resourceList)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947908-initwithdevice?language=objc
func NewCNNDropoutGradientStateWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithDeviceResourceList(device, resourceList)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientstate/2942527-maskdata?language=objc
func (c_ CNNDropoutGradientState) MaskData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("maskData"))
	return rv
}
