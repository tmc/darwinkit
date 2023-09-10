// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageallocator?language=objc
type PImageAllocator interface {
	// optional
	ImageForCommandBufferImageDescriptorKernel(cmdBuf metal.CommandBufferObject, descriptor ImageDescriptor, kernel Kernel) Image
	HasImageForCommandBufferImageDescriptorKernel() bool

	// optional
	ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf metal.CommandBufferObject, descriptor ImageDescriptor, kernel Kernel, count uint) *foundation.Array
	HasImageBatchForCommandBufferImageDescriptorKernelCount() bool
}

// ensure impl type implements protocol interface
var _ PImageAllocator = (*ImageAllocatorObject)(nil)

// A concrete type for the [PImageAllocator] protocol.
type ImageAllocatorObject struct {
	objc.Object
}

func (i_ ImageAllocatorObject) HasImageBatchForCommandBufferImageDescriptorKernelCount() bool {
	return i_.RespondsToSelector(objc.Sel("imageBatchForCommandBuffer:imageDescriptor:kernel:count:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageallocator/3020685-imagebatchforcommandbuffer?language=objc
func (i_ ImageAllocatorObject) ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf metal.CommandBufferObject, descriptor ImageDescriptor, kernel Kernel, count uint) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[*foundation.Array](i_, objc.Sel("imageBatchForCommandBuffer:imageDescriptor:kernel:count:"), po0, descriptor, kernel, count)
	return rv
}

func (i_ ImageAllocatorObject) HasImageForCommandBufferImageDescriptorKernel() bool {
	return i_.RespondsToSelector(objc.Sel("imageForCommandBuffer:imageDescriptor:kernel:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageallocator/2866966-imageforcommandbuffer?language=objc
func (i_ ImageAllocatorObject) ImageForCommandBufferImageDescriptorKernel(cmdBuf metal.CommandBufferObject, descriptor ImageDescriptor, kernel Kernel) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[Image](i_, objc.Sel("imageForCommandBuffer:imageDescriptor:kernel:"), po0, descriptor, kernel)
	return rv
}

func (i_ ImageAllocatorObject) HasImageBatchForCommandBufferImageDescriptorKernelCount() bool {
	return i_.RespondsToSelector(objc.Sel("imageBatchForCommandBuffer:imageDescriptor:kernel:count:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageallocator/3020685-imagebatchforcommandbuffer?language=objc
func (i_ ImageAllocatorObject) ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf metal.CommandBufferObject, descriptor ImageDescriptor, kernel Kernel, count uint) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[*foundation.Array](i_, objc.Sel("imageBatchForCommandBuffer:imageDescriptor:kernel:count:"), po0, objc.Ptr(descriptor), objc.Ptr(kernel), count)
	return rv
}
