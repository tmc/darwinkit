// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NNOptimizerRMSProp] class.
var NNOptimizerRMSPropClass = _NNOptimizerRMSPropClass{objc.GetClass("MPSNNOptimizerRMSProp")}

type _NNOptimizerRMSPropClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizerRMSProp] class.
type INNOptimizerRMSProp interface {
	INNOptimizer
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix)
	Decay() float64
	Epsilon() float32
}

// An optimization layer that performs a root mean square propagation update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop?language=objc
type NNOptimizerRMSProp struct {
	NNOptimizer
}

func NNOptimizerRMSPropFrom(ptr unsafe.Pointer) NNOptimizerRMSProp {
	return NNOptimizerRMSProp{
		NNOptimizer: NNOptimizerFrom(ptr),
	}
}

func (n_ NNOptimizerRMSProp) InitWithDeviceDecayEpsilonOptimizerDescriptor(device metal.PDevice, decay float64, epsilon float32, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerRMSProp {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("initWithDevice:decay:epsilon:optimizerDescriptor:"), po0, decay, epsilon, optimizerDescriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966737-initwithdevice?language=objc
func NewNNOptimizerRMSPropWithDeviceDecayEpsilonOptimizerDescriptor(device metal.PDevice, decay float64, epsilon float32, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerRMSProp {
	instance := NNOptimizerRMSPropClass.Alloc().InitWithDeviceDecayEpsilonOptimizerDescriptor(device, decay, epsilon, optimizerDescriptor)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerRMSProp) InitWithDeviceLearningRate(device metal.PDevice, learningRate float32) NNOptimizerRMSProp {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("initWithDevice:learningRate:"), po0, learningRate)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966738-initwithdevice?language=objc
func NewNNOptimizerRMSPropWithDeviceLearningRate(device metal.PDevice, learningRate float32) NNOptimizerRMSProp {
	instance := NNOptimizerRMSPropClass.Alloc().InitWithDeviceLearningRate(device, learningRate)
	instance.Autorelease()
	return instance
}

func (nc _NNOptimizerRMSPropClass) Alloc() NNOptimizerRMSProp {
	rv := objc.Call[NNOptimizerRMSProp](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNOptimizerRMSPropClass) New() NNOptimizerRMSProp {
	rv := objc.Call[NNOptimizerRMSProp](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizerRMSProp() NNOptimizerRMSProp {
	return NNOptimizerRMSPropClass.New()
}

func (n_ NNOptimizerRMSProp) Init() NNOptimizerRMSProp {
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("init"))
	return rv
}

func (n_ NNOptimizerRMSProp) InitWithDevice(device metal.PDevice) NNOptimizerRMSProp {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNOptimizerRMSPropWithDevice(device metal.PDevice) NNOptimizerRMSProp {
	instance := NNOptimizerRMSPropClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3019335-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState, inputSumOfSquaresVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputSumOfSquaresVectors:resultState:"), po0, batchNormalizationState, inputSumOfSquaresVectors, resultState)
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3019335-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferObjectBatchNormalizationStateInputSumOfSquaresVectorsResultState(commandBufferObject objc.IObject, batchNormalizationState ICNNBatchNormalizationState, inputSumOfSquaresVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputSumOfSquaresVectors:resultState:"), commandBufferObject, batchNormalizationState, inputSumOfSquaresVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3013783-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputSumOfSquaresVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputSumOfSquaresVectors:resultState:"), po0, batchNormalizationGradientState, batchNormalizationSourceState, inputSumOfSquaresVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3013783-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferObjectBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(commandBufferObject objc.IObject, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputSumOfSquaresVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputSumOfSquaresVectors:resultState:"), commandBufferObject, batchNormalizationGradientState, batchNormalizationSourceState, inputSumOfSquaresVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3013784-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(commandBuffer metal.PCommandBuffer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputSumOfSquaresVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputSumOfSquaresVectors:resultState:"), po0, convolutionGradientState, convolutionSourceState, inputSumOfSquaresVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3013784-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferObjectConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(commandBufferObject objc.IObject, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputSumOfSquaresVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputSumOfSquaresVectors:resultState:"), commandBufferObject, convolutionGradientState, convolutionSourceState, inputSumOfSquaresVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966735-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(commandBuffer metal.PCommandBuffer, inputGradientVector IVector, inputValuesVector IVector, inputSumOfSquaresVector IVector, resultValuesVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:resultValuesVector:"), po0, inputGradientVector, inputValuesVector, inputSumOfSquaresVector, resultValuesVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966735-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferObjectInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(commandBufferObject objc.IObject, inputGradientVector IVector, inputValuesVector IVector, inputSumOfSquaresVector IVector, resultValuesVector IVector) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:resultValuesVector:"), commandBufferObject, inputGradientVector, inputValuesVector, inputSumOfSquaresVector, resultValuesVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3197827-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:"), po0, inputGradientMatrix, inputValuesMatrix, inputSumOfSquaresMatrix, resultValuesMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3197827-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:"), commandBufferObject, inputGradientMatrix, inputValuesMatrix, inputSumOfSquaresMatrix, resultValuesMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966734-decay?language=objc
func (n_ NNOptimizerRMSProp) Decay() float64 {
	rv := objc.Call[float64](n_, objc.Sel("decay"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966736-epsilon?language=objc
func (n_ NNOptimizerRMSProp) Epsilon() float32 {
	rv := objc.Call[float32](n_, objc.Sel("epsilon"))
	return rv
}
