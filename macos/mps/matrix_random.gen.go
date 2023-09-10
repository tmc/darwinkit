// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MatrixRandom] class.
var MatrixRandomClass = _MatrixRandomClass{objc.GetClass("MPSMatrixRandom")}

type _MatrixRandomClass struct {
	objc.Class
}

// An interface definition for the [MatrixRandom] class.
type IMatrixRandom interface {
	IKernel
	EncodeToCommandBufferDestinationMatrix(commandBuffer metal.PCommandBuffer, destinationMatrix IMatrix)
	EncodeToCommandBufferObjectDestinationMatrix(commandBufferObject objc.IObject, destinationMatrix IMatrix)
	DistributionType() MatrixRandomDistribution
	DestinationDataType() DataType
	BatchStart() uint
	SetBatchStart(value uint)
	BatchSize() uint
	SetBatchSize(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom?language=objc
type MatrixRandom struct {
	Kernel
}

func MatrixRandomFrom(ptr unsafe.Pointer) MatrixRandom {
	return MatrixRandom{
		Kernel: KernelFrom(ptr),
	}
}

func (mc _MatrixRandomClass) Alloc() MatrixRandom {
	rv := objc.Call[MatrixRandom](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MatrixRandomClass) New() MatrixRandom {
	rv := objc.Call[MatrixRandom](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixRandom() MatrixRandom {
	return MatrixRandomClass.New()
}

func (m_ MatrixRandom) Init() MatrixRandom {
	rv := objc.Call[MatrixRandom](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixRandom) InitWithDevice(device metal.PDevice) MatrixRandom {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixRandom](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixRandomWithDevice(device metal.PDevice) MatrixRandom {
	instance := MatrixRandomClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixRandom) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixRandom {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixRandom](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixRandom_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixRandom {
	instance := MatrixRandomClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3325839-encodetocommandbuffer?language=objc
func (m_ MatrixRandom) EncodeToCommandBufferDestinationMatrix(commandBuffer metal.PCommandBuffer, destinationMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:destinationMatrix:"), po0, destinationMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3325839-encodetocommandbuffer?language=objc
func (m_ MatrixRandom) EncodeToCommandBufferObjectDestinationMatrix(commandBufferObject objc.IObject, destinationMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:destinationMatrix:"), commandBufferObject, destinationMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242851-encodetocommandbuffer?language=objc
func (m_ MatrixRandom) EncodeToCommandBufferDestinationVector(commandBuffer metal.PCommandBuffer, destinationVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:destinationVector:"), po0, destinationVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242851-encodetocommandbuffer?language=objc
func (m_ MatrixRandom) EncodeToCommandBufferObjectDestinationVector(commandBufferObject objc.IObject, destinationVector IVector) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:destinationVector:"), commandBufferObject, destinationVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242848-batchstart?language=objc
func (m_ MatrixRandom) BatchStart() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchStart"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242848-batchstart?language=objc
func (m_ MatrixRandom) SetBatchStart(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchStart:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242847-batchsize?language=objc
func (m_ MatrixRandom) BatchSize() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242847-batchsize?language=objc
func (m_ MatrixRandom) SetBatchSize(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchSize:"), value)
}
