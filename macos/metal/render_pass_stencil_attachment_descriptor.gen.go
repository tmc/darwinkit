// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [RenderPassStencilAttachmentDescriptor] class.
var RenderPassStencilAttachmentDescriptorClass = _RenderPassStencilAttachmentDescriptorClass{objc.GetClass("MTLRenderPassStencilAttachmentDescriptor")}

type _RenderPassStencilAttachmentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RenderPassStencilAttachmentDescriptor] class.
type IRenderPassStencilAttachmentDescriptor interface {
	IRenderPassAttachmentDescriptor
	ClearStencil() uint32
	SetClearStencil(value uint32)
	StencilResolveFilter() MultisampleStencilResolveFilter
	SetStencilResolveFilter(value MultisampleStencilResolveFilter)
}

// A stencil render target that serves as the output destination for stencil pixels generated by a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassstencilattachmentdescriptor?language=objc
type RenderPassStencilAttachmentDescriptor struct {
	RenderPassAttachmentDescriptor
}

func RenderPassStencilAttachmentDescriptorFrom(ptr unsafe.Pointer) RenderPassStencilAttachmentDescriptor {
	return RenderPassStencilAttachmentDescriptor{
		RenderPassAttachmentDescriptor: RenderPassAttachmentDescriptorFrom(ptr),
	}
}

func (rc _RenderPassStencilAttachmentDescriptorClass) Alloc() RenderPassStencilAttachmentDescriptor {
	rv := objc.Call[RenderPassStencilAttachmentDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func (rc _RenderPassStencilAttachmentDescriptorClass) New() RenderPassStencilAttachmentDescriptor {
	rv := objc.Call[RenderPassStencilAttachmentDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPassStencilAttachmentDescriptor() RenderPassStencilAttachmentDescriptor {
	return RenderPassStencilAttachmentDescriptorClass.New()
}

func (r_ RenderPassStencilAttachmentDescriptor) Init() RenderPassStencilAttachmentDescriptor {
	rv := objc.Call[RenderPassStencilAttachmentDescriptor](r_, objc.Sel("init"))
	return rv
}

// The value to use when clearing the stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassstencilattachmentdescriptor/1437931-clearstencil?language=objc
func (r_ RenderPassStencilAttachmentDescriptor) ClearStencil() uint32 {
	rv := objc.Call[uint32](r_, objc.Sel("clearStencil"))
	return rv
}

// The value to use when clearing the stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassstencilattachmentdescriptor/1437931-clearstencil?language=objc
func (r_ RenderPassStencilAttachmentDescriptor) SetClearStencil(value uint32) {
	objc.Call[objc.Void](r_, objc.Sel("setClearStencil:"), value)
}

// The filter used for stencil multisample resolve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassstencilattachmentdescriptor/2967446-stencilresolvefilter?language=objc
func (r_ RenderPassStencilAttachmentDescriptor) StencilResolveFilter() MultisampleStencilResolveFilter {
	rv := objc.Call[MultisampleStencilResolveFilter](r_, objc.Sel("stencilResolveFilter"))
	return rv
}

// The filter used for stencil multisample resolve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassstencilattachmentdescriptor/2967446-stencilresolvefilter?language=objc
func (r_ RenderPassStencilAttachmentDescriptor) SetStencilResolveFilter(value MultisampleStencilResolveFilter) {
	objc.Call[objc.Void](r_, objc.Sel("setStencilResolveFilter:"), value)
}
