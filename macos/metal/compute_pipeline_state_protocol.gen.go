// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// An object that contains a compiled compute pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate?language=objc
type PComputePipelineState interface {
	// optional
	NewIntersectionFunctionTableWithDescriptor(descriptor IntersectionFunctionTableDescriptor) IntersectionFunctionTableObject
	HasNewIntersectionFunctionTableWithDescriptor() bool

	// optional
	NewVisibleFunctionTableWithDescriptor(descriptor VisibleFunctionTableDescriptor) VisibleFunctionTableObject
	HasNewVisibleFunctionTableWithDescriptor() bool

	// optional
	NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []FunctionObject, error unsafe.Pointer) ComputePipelineStateObject
	HasNewComputePipelineStateWithAdditionalBinaryFunctionsError() bool

	// optional
	Device() DeviceObject
	HasDevice() bool

	// optional
	MaxTotalThreadsPerThreadgroup() uint
	HasMaxTotalThreadsPerThreadgroup() bool

	// optional
	ThreadExecutionWidth() uint
	HasThreadExecutionWidth() bool

	// optional
	StaticThreadgroupMemoryLength() uint
	HasStaticThreadgroupMemoryLength() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	SupportIndirectCommandBuffers() bool
	HasSupportIndirectCommandBuffers() bool
}

// ensure impl type implements protocol interface
var _ PComputePipelineState = (*ComputePipelineStateObject)(nil)

// A concrete type for the [PComputePipelineState] protocol.
type ComputePipelineStateObject struct {
	objc.Object
}

func (c_ ComputePipelineStateObject) HasNewIntersectionFunctionTableWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("newIntersectionFunctionTableWithDescriptor:"))
}

// Creates a new intersection function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3580381-newintersectionfunctiontablewith?language=objc
func (c_ ComputePipelineStateObject) NewIntersectionFunctionTableWithDescriptor(descriptor IntersectionFunctionTableDescriptor) IntersectionFunctionTableObject {
	rv := objc.Call[IntersectionFunctionTableObject](c_, objc.Sel("newIntersectionFunctionTableWithDescriptor:"), descriptor)
	return rv
}

func (c_ ComputePipelineStateObject) HasNewVisibleFunctionTableWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("newVisibleFunctionTableWithDescriptor:"))
}

// Creates a new visible function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3566543-newvisiblefunctiontablewithdescr?language=objc
func (c_ ComputePipelineStateObject) NewVisibleFunctionTableWithDescriptor(descriptor VisibleFunctionTableDescriptor) VisibleFunctionTableObject {
	rv := objc.Call[VisibleFunctionTableObject](c_, objc.Sel("newVisibleFunctionTableWithDescriptor:"), descriptor)
	return rv
}

func (c_ ComputePipelineStateObject) HasFunctionHandleWithFunction() bool {
	return c_.RespondsToSelector(objc.Sel("functionHandleWithFunction:"))
}

// Creates a function handle for a visible function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3553964-functionhandlewithfunction?language=objc
func (c_ ComputePipelineStateObject) FunctionHandleWithFunction(function FunctionObject) FunctionHandleObject {
	po0 := objc.WrapAsProtocol("MTLFunction", function)
	rv := objc.Call[FunctionHandleObject](c_, objc.Sel("functionHandleWithFunction:"), po0)
	return rv
}

func (c_ ComputePipelineStateObject) HasImageblockMemoryLengthForDimensions() bool {
	return c_.RespondsToSelector(objc.Sel("imageblockMemoryLengthForDimensions:"))
}

// Returns the imageblock memory length, in bytes, for the specified imageblock dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2928195-imageblockmemorylengthfordimensi?language=objc
func (c_ ComputePipelineStateObject) ImageblockMemoryLengthForDimensions(imageblockDimensions Size) uint {
	rv := objc.Call[uint](c_, objc.Sel("imageblockMemoryLengthForDimensions:"), imageblockDimensions)
	return rv
}

func (c_ ComputePipelineStateObject) HasNewComputePipelineStateWithAdditionalBinaryFunctionsError() bool {
	return c_.RespondsToSelector(objc.Sel("newComputePipelineStateWithAdditionalBinaryFunctions:error:"))
}

// Creates a new pipeline state object with additional callable functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3580380-newcomputepipelinestatewithaddit?language=objc
func (c_ ComputePipelineStateObject) NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []FunctionObject, error unsafe.Pointer) ComputePipelineStateObject {
	rv := objc.Call[ComputePipelineStateObject](c_, objc.Sel("newComputePipelineStateWithAdditionalBinaryFunctions:error:"), functions, error)
	return rv
}

func (c_ ComputePipelineStateObject) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the pipeline state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/1414925-device?language=objc
func (c_ ComputePipelineStateObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](c_, objc.Sel("device"))
	return rv
}

func (c_ ComputePipelineStateObject) HasMaxTotalThreadsPerThreadgroup() bool {
	return c_.RespondsToSelector(objc.Sel("maxTotalThreadsPerThreadgroup"))
}

// The maximum number of threads in a threadgroup that you can dispatch to the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/1414927-maxtotalthreadsperthreadgroup?language=objc
func (c_ ComputePipelineStateObject) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Call[uint](c_, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}

func (c_ ComputePipelineStateObject) HasThreadExecutionWidth() bool {
	return c_.RespondsToSelector(objc.Sel("threadExecutionWidth"))
}

// The number of threads that the GPU executes simultaneously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/1414911-threadexecutionwidth?language=objc
func (c_ ComputePipelineStateObject) ThreadExecutionWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("threadExecutionWidth"))
	return rv
}

func (c_ ComputePipelineStateObject) HasStaticThreadgroupMemoryLength() bool {
	return c_.RespondsToSelector(objc.Sel("staticThreadgroupMemoryLength"))
}

// The length, in bytes, of statically allocated threadgroup memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2877435-staticthreadgroupmemorylength?language=objc
func (c_ ComputePipelineStateObject) StaticThreadgroupMemoryLength() uint {
	rv := objc.Call[uint](c_, objc.Sel("staticThreadgroupMemoryLength"))
	return rv
}

func (c_ ComputePipelineStateObject) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the compute pipeline state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2880323-label?language=objc
func (c_ ComputePipelineStateObject) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ ComputePipelineStateObject) HasSupportIndirectCommandBuffers() bool {
	return c_.RespondsToSelector(objc.Sel("supportIndirectCommandBuffers"))
}

// A Boolean value that indicates whether the pipeline supports indirect command buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2966562-supportindirectcommandbuffers?language=objc
func (c_ ComputePipelineStateObject) SupportIndirectCommandBuffers() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}
