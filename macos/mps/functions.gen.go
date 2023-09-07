// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "MetalPerformanceShaders/MetalPerformanceShaders.h"
// uint StateBatchResourceSize(MPSStateBatch* batch);
// void HintTemporaryMemoryHighWaterMark(void * cmdBuf, uint bytes);
// uint ImageBatchResourceSize(MPSImageBatch* batch);
// void SetHeapCacheDuration(void * cmdBuf, double seconds);
// uint StateBatchIncrementReadCount(MPSStateBatch* batch, long amount);
// void StateBatchSynchronize(MPSStateBatch* batch, void * cmdBuf);
// int32_t GetCustomKernelBatchedDestinationIndex(MPSCustomKernelArgumentCount c);
// int32_t GetCustomKernelMaxBatchSize(MPSCustomKernelArgumentCount c, int32_t MPSMaxTextures);
// uint ImageBatchIncrementReadCount(MPSImageBatch* batch, long amount);
// void * GetPreferredDevice(MPSDeviceOptions options);
// int32_t GetCustomKernelBroadcastSourceIndex(MPSCustomKernelArgumentCount c, int32_t sourceIndex, int32_t MPSMaxTextures);
// MPSImageType GetImageType(void * image);
// MPSIntegerDivisionParams FindIntegerDivisionParams(uint16_t divisor);
// int32_t GetCustomKernelBatchedSourceIndex(MPSCustomKernelArgumentCount c, int32_t sourceIndex, int32_t MPSMaxTextures);
// void ImageBatchSynchronize(MPSImageBatch* batch, void * cmdBuf);
// bool SupportsMTLDevice(void * device);
import "C"
import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// Returns the number of bytes used to allocate the specified state batch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2980728-mpsstatebatchresourcesize?language=objc
func StateBatchResourceSize(batch *foundation.Array) uint {
	rv := C.StateBatchResourceSize(
		// *typing.PointerType
		(*C.MPSStateBatch)(unsafe.Pointer(&batch)),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Triggers Metal Performance Shaders to prefetch a Metal heap of the indicated size into its internal cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990485-mpshinttemporarymemoryhighwaterm?language=objc
func HintTemporaryMemoryHighWaterMark(cmdBuf metal.CommandBufferWrapper, bytes uint) {
	C.HintTemporaryMemoryHighWaterMark(
		// *typing.ProtocolType
		unsafe.Pointer(&cmdBuf),
		// *typing.PrimitiveType
		C.uint(bytes),
	)
}

// Returns the number of bytes used to allocate the specified image batch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2980727-mpsimagebatchresourcesize?language=objc
func ImageBatchResourceSize(batch *foundation.Array) uint {
	rv := C.ImageBatchResourceSize(
		// *typing.PointerType
		(*C.MPSImageBatch)(unsafe.Pointer(&batch)),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Sets the timeout after which unused cached Metal heaps are released. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990486-mpssetheapcacheduration?language=objc
func SetHeapCacheDuration(cmdBuf metal.CommandBufferWrapper, seconds float64) {
	C.SetHeapCacheDuration(
		// *typing.ProtocolType
		unsafe.Pointer(&cmdBuf),
		// *typing.PrimitiveType
		C.double(seconds),
	)
}

// Increments or decrements the read count of a state batch by a specified amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2951920-mpsstatebatchincrementreadcount?language=objc
func StateBatchIncrementReadCount(batch *foundation.Array, amount int) uint {
	rv := C.StateBatchIncrementReadCount(
		// *typing.PointerType
		(*C.MPSStateBatch)(unsafe.Pointer(&batch)),
		// *typing.PrimitiveType
		C.long(amount),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Removes any copy of the specified state batch from the device's caches, and, if needed, invalidates any CPU caches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2953920-mpsstatebatchsynchronize?language=objc
func StateBatchSynchronize(batch *foundation.Array, cmdBuf metal.CommandBufferWrapper) {
	C.StateBatchSynchronize(
		// *typing.PointerType
		(*C.MPSStateBatch)(unsafe.Pointer(&batch)),
		// *typing.ProtocolType
		unsafe.Pointer(&cmdBuf),
	)
}

// Returns the index of the first destination texture argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990481-mpsgetcustomkernelbatcheddestina?language=objc
func GetCustomKernelBatchedDestinationIndex(c CustomKernelArgumentCount) int32 {
	rv := C.GetCustomKernelBatchedDestinationIndex(
		// *typing.StructType
		*(*C.MPSCustomKernelArgumentCount)(unsafe.Pointer(&c)),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns the maximum allowed batch size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990484-mpsgetcustomkernelmaxbatchsize?language=objc
func GetCustomKernelMaxBatchSize(c CustomKernelArgumentCount, MPSMaxTextures int32) int32 {
	rv := C.GetCustomKernelMaxBatchSize(
		// *typing.StructType
		*(*C.MPSCustomKernelArgumentCount)(unsafe.Pointer(&c)),
		// *typing.PrimitiveType
		C.int32_t(MPSMaxTextures),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Increments or decrements the read count of an image batch by a specified amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2951916-mpsimagebatchincrementreadcount?language=objc
func ImageBatchIncrementReadCount(batch *foundation.Array, amount int) uint {
	rv := C.ImageBatchIncrementReadCount(
		// *typing.PointerType
		(*C.MPSImageBatch)(unsafe.Pointer(&batch)),
		// *typing.PrimitiveType
		C.long(amount),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/3088918-mpsgetpreferreddevice?language=objc
func GetPreferredDevice(options DeviceOptions) metal.DeviceWrapper {
	rv := C.GetPreferredDevice(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.MPSDeviceOptions)(options),
	)
	// *typing.ProtocolType
	return metal.DeviceWrapper{objc.ObjectFrom(rv)}
}

// Returns the index of the specified nonbatched source texture argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990483-mpsgetcustomkernelbroadcastsourc?language=objc
func GetCustomKernelBroadcastSourceIndex(c CustomKernelArgumentCount, sourceIndex int32, MPSMaxTextures int32) int32 {
	rv := C.GetCustomKernelBroadcastSourceIndex(
		// *typing.StructType
		*(*C.MPSCustomKernelArgumentCount)(unsafe.Pointer(&c)),
		// *typing.PrimitiveType
		C.int32_t(sourceIndex),
		// *typing.PrimitiveType
		C.int32_t(MPSMaxTextures),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/3131717-mpsgetimagetype?language=objc
func GetImageType(image Image) ImageType {
	rv := C.GetImageType(
		// *typing.ClassType
		unsafe.Pointer(&image),
	)
	// *typing.AliasType
	return *(*ImageType)(unsafe.Pointer(&rv))
}

// Returns the integer division parameters for a specified divisor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2954868-mpsfindintegerdivisionparams?language=objc
func FindIntegerDivisionParams(divisor uint16) IntegerDivisionParams {
	rv := C.FindIntegerDivisionParams(
		// *typing.PrimitiveType
		C.uint16_t(divisor),
	)
	// *typing.StructType
	return *(*IntegerDivisionParams)(unsafe.Pointer(&rv))
}

// Returns the index of the specified batched source texture argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990482-mpsgetcustomkernelbatchedsourcei?language=objc
func GetCustomKernelBatchedSourceIndex(c CustomKernelArgumentCount, sourceIndex int32, MPSMaxTextures int32) int32 {
	rv := C.GetCustomKernelBatchedSourceIndex(
		// *typing.StructType
		*(*C.MPSCustomKernelArgumentCount)(unsafe.Pointer(&c)),
		// *typing.PrimitiveType
		C.int32_t(sourceIndex),
		// *typing.PrimitiveType
		C.int32_t(MPSMaxTextures),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Removes any copy of the specified image batch from the device's caches, and, if needed, invalidates any CPU caches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2953951-mpsimagebatchsynchronize?language=objc
func ImageBatchSynchronize(batch *foundation.Array, cmdBuf metal.CommandBufferWrapper) {
	C.ImageBatchSynchronize(
		// *typing.PointerType
		(*C.MPSImageBatch)(unsafe.Pointer(&batch)),
		// *typing.ProtocolType
		unsafe.Pointer(&cmdBuf),
	)
}

// Determines whether the Metal Performance Shaders framework supports a Metal device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/1618849-mpssupportsmtldevice?language=objc
func SupportsMTLDevice(device metal.DeviceWrapper) bool {
	rv := C.SupportsMTLDevice(
		// *typing.ProtocolType
		unsafe.Pointer(&device),
	)
	// *typing.PrimitiveType
	return bool(rv)
}
