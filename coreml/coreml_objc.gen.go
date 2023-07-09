package coreml

import (
	core "github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework CoreML
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <CoreML/CoreML.h>

bool coreml_convertObjCBool(BOOL b) {
	if (b) { return true; }
	return false;
}


void* MLArrayBatchProvider_type_Alloc() {
	return [MLArrayBatchProvider
		alloc];
}
void* MLCPUComputeDevice_type_Alloc() {
	return [MLCPUComputeDevice
		alloc];
}
void* MLDictionaryFeatureProvider_type_Alloc() {
	return [MLDictionaryFeatureProvider
		alloc];
}
void* MLFeatureValue_type_Alloc() {
	return [MLFeatureValue
		alloc];
}
void* MLFeatureValue_type_FeatureValueWithString(void* value) {
	return [MLFeatureValue
		featureValueWithString: value];
}
void* MLFeatureValue_type_FeatureValueWithDictionaryError(void* value, void* error) {
	return [MLFeatureValue
		featureValueWithDictionary: value
		error: error];
}
void* MLGPUComputeDevice_type_Alloc() {
	return [MLGPUComputeDevice
		alloc];
}
void* MLModel_type_Alloc() {
	return [MLModel
		alloc];
}
void* MLModel_type_ModelWithContentsOfURLError(void* url, void* error) {
	return [MLModel
		modelWithContentsOfURL: url
		error: error];
}
void* MLModel_type_AvailableComputeDevices() {
	return [MLModel
		availableComputeDevices];
}
void* MLModelAsset_type_Alloc() {
	return [MLModelAsset
		alloc];
}
void* MLModelAsset_type_ModelAssetWithSpecificationDataError(void* specificationData, void* error) {
	return [MLModelAsset
		modelAssetWithSpecificationData: specificationData
		error: error];
}
void* MLModelCollection_type_Alloc() {
	return [MLModelCollection
		alloc];
}
void* MLNeuralEngineComputeDevice_type_Alloc() {
	return [MLNeuralEngineComputeDevice
		alloc];
}


void* MLArrayBatchProvider_inst_InitWithDictionaryError(void *id, void* dictionary, void* error) {
	return [(MLArrayBatchProvider*)id
		initWithDictionary: dictionary
		error: error];
}

void* MLArrayBatchProvider_inst_InitWithFeatureProviderArray(void *id, void* array) {
	return [(MLArrayBatchProvider*)id
		initWithFeatureProviderArray: array];
}

void* MLArrayBatchProvider_inst_Init(void *id) {
	return [(MLArrayBatchProvider*)id
		init];
}

void* MLArrayBatchProvider_inst_Array(void *id) {
	return [(MLArrayBatchProvider*)id
		array];
}

void* MLCPUComputeDevice_inst_Init(void *id) {
	return [(MLCPUComputeDevice*)id
		init];
}

void* MLDictionaryFeatureProvider_inst_InitWithDictionaryError(void *id, void* dictionary, void* error) {
	return [(MLDictionaryFeatureProvider*)id
		initWithDictionary: dictionary
		error: error];
}

void* MLDictionaryFeatureProvider_inst_ObjectForKeyedSubscript(void *id, void* featureName) {
	return [(MLDictionaryFeatureProvider*)id
		objectForKeyedSubscript: featureName];
}

void* MLDictionaryFeatureProvider_inst_Init(void *id) {
	return [(MLDictionaryFeatureProvider*)id
		init];
}

void* MLDictionaryFeatureProvider_inst_Dictionary(void *id) {
	return [(MLDictionaryFeatureProvider*)id
		dictionary];
}

BOOL MLFeatureValue_inst_IsEqualToFeatureValue(void *id, void* value) {
	return [(MLFeatureValue*)id
		isEqualToFeatureValue: value];
}

void* MLFeatureValue_inst_Init(void *id) {
	return [(MLFeatureValue*)id
		init];
}

BOOL MLFeatureValue_inst_IsUndefined(void *id) {
	return [(MLFeatureValue*)id
		isUndefined];
}

void* MLFeatureValue_inst_StringValue(void *id) {
	return [(MLFeatureValue*)id
		stringValue];
}

void* MLFeatureValue_inst_DictionaryValue(void *id) {
	return [(MLFeatureValue*)id
		dictionaryValue];
}

void* MLGPUComputeDevice_inst_Init(void *id) {
	return [(MLGPUComputeDevice*)id
		init];
}

void* MLGPUComputeDevice_inst_MetalDevice(void *id) {
	return [(MLGPUComputeDevice*)id
		metalDevice];
}

void* MLModel_inst_PredictionFromFeaturesError(void *id, void* input, void* error) {
	return [(MLModel*)id
		predictionFromFeatures: input
		error: error];
}

void* MLModel_inst_PredictionsFromBatchError(void *id, void* inputBatch, void* error) {
	return [(MLModel*)id
		predictionsFromBatch: inputBatch
		error: error];
}

void* MLModel_inst_Init(void *id) {
	return [(MLModel*)id
		init];
}

void* MLModelAsset_inst_Init(void *id) {
	return [(MLModelAsset*)id
		init];
}

void* MLModelCollection_inst_Init(void *id) {
	return [(MLModelCollection*)id
		init];
}

void* MLNeuralEngineComputeDevice_inst_Init(void *id) {
	return [(MLNeuralEngineComputeDevice*)id
		init];
}

long MLNeuralEngineComputeDevice_inst_TotalCoreCount(void *id) {
	return [(MLNeuralEngineComputeDevice*)id
		totalCoreCount];
}


BOOL coreml_objc_bool_true = YES;
BOOL coreml_objc_bool_false = NO;

*/
import "C"

func convertObjCBoolToGo(b C.BOOL) bool {
	// NOTE: the prefix here is used to namespace these since the linker will
	// otherwise report a "duplicate symbol" because the C functions have the
	// same name.
	return bool(C.coreml_convertObjCBool(b))
}

func convertToObjCBool(b bool) C.BOOL {
	if b {
		return C.coreml_objc_bool_true
	}
	return C.coreml_objc_bool_false
}

// MLArrayBatchProvider_Alloc is undocumented.
func MLArrayBatchProvider_Alloc() MLArrayBatchProvider {
	ret := C.MLArrayBatchProvider_type_Alloc()

	return MLArrayBatchProvider_FromPointer(ret)
}

// MLCPUComputeDevice_Alloc is undocumented.
func MLCPUComputeDevice_Alloc() MLCPUComputeDevice {
	ret := C.MLCPUComputeDevice_type_Alloc()

	return MLCPUComputeDevice_FromPointer(ret)
}

// MLDictionaryFeatureProvider_Alloc is undocumented.
func MLDictionaryFeatureProvider_Alloc() MLDictionaryFeatureProvider {
	ret := C.MLDictionaryFeatureProvider_type_Alloc()

	return MLDictionaryFeatureProvider_FromPointer(ret)
}

// MLFeatureValue_Alloc is undocumented.
func MLFeatureValue_Alloc() MLFeatureValue {
	ret := C.MLFeatureValue_type_Alloc()

	return MLFeatureValue_FromPointer(ret)
}

// MLFeatureValue_FeatureValueWithString creates a feature value that contains a string.
//
// See https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879343-featurevaluewithstring?language=objc for details.
func MLFeatureValue_FeatureValueWithString(value core.NSStringRef) MLFeatureValue {
	ret := C.MLFeatureValue_type_FeatureValueWithString(
		objc.RefPointer(value),
	)

	return MLFeatureValue_FromPointer(ret)
}

// MLFeatureValue_FeatureValueWithDictionaryError creates a feature value that contains a dictionary of numbers.
//
// See https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879393-featurevaluewithdictionary?language=objc for details.
func MLFeatureValue_FeatureValueWithDictionaryError(value core.NSDictionaryRef, error core.NSErrorRef) MLFeatureValue {
	ret := C.MLFeatureValue_type_FeatureValueWithDictionaryError(
		objc.RefPointer(value),
		objc.RefPointer(error),
	)

	return MLFeatureValue_FromPointer(ret)
}

// MLGPUComputeDevice_Alloc is undocumented.
func MLGPUComputeDevice_Alloc() MLGPUComputeDevice {
	ret := C.MLGPUComputeDevice_type_Alloc()

	return MLGPUComputeDevice_FromPointer(ret)
}

// MLModel_Alloc is undocumented.
func MLModel_Alloc() MLModel {
	ret := C.MLModel_type_Alloc()

	return MLModel_FromPointer(ret)
}

// MLModel_ModelWithContentsOfURLError creates a Core ML model instance from a compiled model file.
//
// See https://developer.apple.com/documentation/coreml/mlmodel/2880279-modelwithcontentsofurl?language=objc for details.
func MLModel_ModelWithContentsOfURLError(url core.NSURLRef, error core.NSErrorRef) MLModel {
	ret := C.MLModel_type_ModelWithContentsOfURLError(
		objc.RefPointer(url),
		objc.RefPointer(error),
	)

	return MLModel_FromPointer(ret)
}

// MLModel_AvailableComputeDevices is undocumented.
//
// See https://developer.apple.com/documentation/coreml/mlmodel/4230952-availablecomputedevices?language=objc for details.
func MLModel_AvailableComputeDevices() core.NSArray {
	ret := C.MLModel_type_AvailableComputeDevices()

	return core.NSArray_FromPointer(ret)
}

// MLModelAsset_Alloc is undocumented.
func MLModelAsset_Alloc() MLModelAsset {
	ret := C.MLModelAsset_type_Alloc()

	return MLModelAsset_FromPointer(ret)
}

// MLModelAsset_ModelAssetWithSpecificationDataError creates a model asset from an in-memory model specification.
//
// See https://developer.apple.com/documentation/coreml/mlmodelasset/3950977-modelassetwithspecificationdata?language=objc for details.
func MLModelAsset_ModelAssetWithSpecificationDataError(specificationData core.NSDataRef, error core.NSErrorRef) MLModelAsset {
	ret := C.MLModelAsset_type_ModelAssetWithSpecificationDataError(
		objc.RefPointer(specificationData),
		objc.RefPointer(error),
	)

	return MLModelAsset_FromPointer(ret)
}

// MLModelCollection_Alloc is undocumented.
func MLModelCollection_Alloc() MLModelCollection {
	ret := C.MLModelCollection_type_Alloc()

	return MLModelCollection_FromPointer(ret)
}

// MLNeuralEngineComputeDevice_Alloc is undocumented.
func MLNeuralEngineComputeDevice_Alloc() MLNeuralEngineComputeDevice {
	ret := C.MLNeuralEngineComputeDevice_type_Alloc()

	return MLNeuralEngineComputeDevice_FromPointer(ret)
}

type MLArrayBatchProviderRef interface {
	Pointer() uintptr
	Init_AsMLArrayBatchProvider() MLArrayBatchProvider
}

type gen_MLArrayBatchProvider struct {
	objc.Object
}

func MLArrayBatchProvider_FromPointer(ptr unsafe.Pointer) MLArrayBatchProvider {
	return MLArrayBatchProvider{gen_MLArrayBatchProvider{
		objc.Object_FromPointer(ptr),
	}}
}

func MLArrayBatchProvider_FromRef(ref objc.Ref) MLArrayBatchProvider {
	return MLArrayBatchProvider_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithDictionaryError creates a batch provider based on feature names and their associated arrays of data.
//
// See https://developer.apple.com/documentation/coreml/mlarraybatchprovider/2962853-initwithdictionary?language=objc for details.
func (x gen_MLArrayBatchProvider) InitWithDictionaryError(
	dictionary core.NSDictionaryRef,
	error core.NSErrorRef,
) MLArrayBatchProvider {
	ret := C.MLArrayBatchProvider_inst_InitWithDictionaryError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(dictionary),
		objc.RefPointer(error),
	)

	return MLArrayBatchProvider_FromPointer(ret)
}

// InitWithFeatureProviderArray creates the batch provider based on the array of feature providers.
//
// See https://developer.apple.com/documentation/coreml/mlarraybatchprovider/2962854-initwithfeatureproviderarray?language=objc for details.
func (x gen_MLArrayBatchProvider) InitWithFeatureProviderArray(
	array core.NSArrayRef,
) MLArrayBatchProvider {
	ret := C.MLArrayBatchProvider_inst_InitWithFeatureProviderArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(array),
	)

	return MLArrayBatchProvider_FromPointer(ret)
}

// Init initializes a new instance of the MLArrayBatchProvider class.
func (x gen_MLArrayBatchProvider) Init() MLArrayBatchProvider {
	ret := C.MLArrayBatchProvider_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLArrayBatchProvider_FromPointer(ret)
}

// Init_AsMLArrayBatchProvider is a typed version of Init.
func (x gen_MLArrayBatchProvider) Init_AsMLArrayBatchProvider() MLArrayBatchProvider {
	ret := C.MLArrayBatchProvider_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLArrayBatchProvider_FromPointer(ret)
}

// Array returns the array of feature providers.
//
// See https://developer.apple.com/documentation/coreml/mlarraybatchprovider/2962852-array?language=objc for details.
func (x gen_MLArrayBatchProvider) Array() core.NSArray {
	ret := C.MLArrayBatchProvider_inst_Array(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

type MLCPUComputeDeviceRef interface {
	Pointer() uintptr
	Init_AsMLCPUComputeDevice() MLCPUComputeDevice
}

type gen_MLCPUComputeDevice struct {
	objc.Object
}

func MLCPUComputeDevice_FromPointer(ptr unsafe.Pointer) MLCPUComputeDevice {
	return MLCPUComputeDevice{gen_MLCPUComputeDevice{
		objc.Object_FromPointer(ptr),
	}}
}

func MLCPUComputeDevice_FromRef(ref objc.Ref) MLCPUComputeDevice {
	return MLCPUComputeDevice_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the MLCPUComputeDevice class.
func (x gen_MLCPUComputeDevice) Init() MLCPUComputeDevice {
	ret := C.MLCPUComputeDevice_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLCPUComputeDevice_FromPointer(ret)
}

// Init_AsMLCPUComputeDevice is a typed version of Init.
func (x gen_MLCPUComputeDevice) Init_AsMLCPUComputeDevice() MLCPUComputeDevice {
	ret := C.MLCPUComputeDevice_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLCPUComputeDevice_FromPointer(ret)
}

type MLDictionaryFeatureProviderRef interface {
	Pointer() uintptr
	Init_AsMLDictionaryFeatureProvider() MLDictionaryFeatureProvider
}

type gen_MLDictionaryFeatureProvider struct {
	objc.Object
}

func MLDictionaryFeatureProvider_FromPointer(ptr unsafe.Pointer) MLDictionaryFeatureProvider {
	return MLDictionaryFeatureProvider{gen_MLDictionaryFeatureProvider{
		objc.Object_FromPointer(ptr),
	}}
}

func MLDictionaryFeatureProvider_FromRef(ref objc.Ref) MLDictionaryFeatureProvider {
	return MLDictionaryFeatureProvider_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithDictionaryError creates the feature provider based on a dictionary.
//
// See https://developer.apple.com/documentation/coreml/mldictionaryfeatureprovider/2879366-initwithdictionary?language=objc for details.
func (x gen_MLDictionaryFeatureProvider) InitWithDictionaryError(
	dictionary core.NSDictionaryRef,
	error core.NSErrorRef,
) MLDictionaryFeatureProvider {
	ret := C.MLDictionaryFeatureProvider_inst_InitWithDictionaryError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(dictionary),
		objc.RefPointer(error),
	)

	return MLDictionaryFeatureProvider_FromPointer(ret)
}

// ObjectForKeyedSubscript subscript interface for the feature provider to pass through to the dictionary.
//
// See https://developer.apple.com/documentation/coreml/mldictionaryfeatureprovider/2881954-objectforkeyedsubscript?language=objc for details.
func (x gen_MLDictionaryFeatureProvider) ObjectForKeyedSubscript(
	featureName core.NSStringRef,
) MLFeatureValue {
	ret := C.MLDictionaryFeatureProvider_inst_ObjectForKeyedSubscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(featureName),
	)

	return MLFeatureValue_FromPointer(ret)
}

// Init initializes a new instance of the MLDictionaryFeatureProvider class.
func (x gen_MLDictionaryFeatureProvider) Init() MLDictionaryFeatureProvider {
	ret := C.MLDictionaryFeatureProvider_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLDictionaryFeatureProvider_FromPointer(ret)
}

// Init_AsMLDictionaryFeatureProvider is a typed version of Init.
func (x gen_MLDictionaryFeatureProvider) Init_AsMLDictionaryFeatureProvider() MLDictionaryFeatureProvider {
	ret := C.MLDictionaryFeatureProvider_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLDictionaryFeatureProvider_FromPointer(ret)
}

// Dictionary returns the backing dictionary.
//
// See https://developer.apple.com/documentation/coreml/mldictionaryfeatureprovider/2879354-dictionary?language=objc for details.
func (x gen_MLDictionaryFeatureProvider) Dictionary() core.NSDictionary {
	ret := C.MLDictionaryFeatureProvider_inst_Dictionary(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

type MLFeatureValueRef interface {
	Pointer() uintptr
	Init_AsMLFeatureValue() MLFeatureValue
}

type gen_MLFeatureValue struct {
	objc.Object
}

func MLFeatureValue_FromPointer(ptr unsafe.Pointer) MLFeatureValue {
	return MLFeatureValue{gen_MLFeatureValue{
		objc.Object_FromPointer(ptr),
	}}
}

func MLFeatureValue_FromRef(ref objc.Ref) MLFeatureValue {
	return MLFeatureValue_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// IsEqualToFeatureValue returns a Boolean value that indicates whether a feature value is equal to another.
//
// See https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879399-isequaltofeaturevalue?language=objc for details.
func (x gen_MLFeatureValue) IsEqualToFeatureValue(
	value MLFeatureValueRef,
) bool {
	ret := C.MLFeatureValue_inst_IsEqualToFeatureValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return convertObjCBoolToGo(ret)
}

// Init initializes a new instance of the MLFeatureValue class.
func (x gen_MLFeatureValue) Init() MLFeatureValue {
	ret := C.MLFeatureValue_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLFeatureValue_FromPointer(ret)
}

// Init_AsMLFeatureValue is a typed version of Init.
func (x gen_MLFeatureValue) Init_AsMLFeatureValue() MLFeatureValue {
	ret := C.MLFeatureValue_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLFeatureValue_FromPointer(ret)
}

// IsUndefined returns a Boolean value that indicates whether the feature value is undefined or missing.
//
// See https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879392-undefined?language=objc for details.
func (x gen_MLFeatureValue) IsUndefined() bool {
	ret := C.MLFeatureValue_inst_IsUndefined(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// StringValue returns the underlying string of the feature value.
//
// See https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879349-stringvalue?language=objc for details.
func (x gen_MLFeatureValue) StringValue() core.NSString {
	ret := C.MLFeatureValue_inst_StringValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// DictionaryValue returns the underlying dictionary of the feature value.
//
// See https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879387-dictionaryvalue?language=objc for details.
func (x gen_MLFeatureValue) DictionaryValue() core.NSDictionary {
	ret := C.MLFeatureValue_inst_DictionaryValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

type MLGPUComputeDeviceRef interface {
	Pointer() uintptr
	Init_AsMLGPUComputeDevice() MLGPUComputeDevice
}

type gen_MLGPUComputeDevice struct {
	objc.Object
}

func MLGPUComputeDevice_FromPointer(ptr unsafe.Pointer) MLGPUComputeDevice {
	return MLGPUComputeDevice{gen_MLGPUComputeDevice{
		objc.Object_FromPointer(ptr),
	}}
}

func MLGPUComputeDevice_FromRef(ref objc.Ref) MLGPUComputeDevice {
	return MLGPUComputeDevice_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the MLGPUComputeDevice class.
func (x gen_MLGPUComputeDevice) Init() MLGPUComputeDevice {
	ret := C.MLGPUComputeDevice_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLGPUComputeDevice_FromPointer(ret)
}

// Init_AsMLGPUComputeDevice is a typed version of Init.
func (x gen_MLGPUComputeDevice) Init_AsMLGPUComputeDevice() MLGPUComputeDevice {
	ret := C.MLGPUComputeDevice_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLGPUComputeDevice_FromPointer(ret)
}

// MetalDevice returns the device that represents the underlying metal device.
//
// See https://developer.apple.com/documentation/coreml/mlgpucomputedevice/4134606-metaldevice?language=objc for details.
func (x gen_MLGPUComputeDevice) MetalDevice() objc.Object {
	ret := C.MLGPUComputeDevice_inst_MetalDevice(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

type MLModelRef interface {
	Pointer() uintptr
	Init_AsMLModel() MLModel
}

type gen_MLModel struct {
	objc.Object
}

func MLModel_FromPointer(ptr unsafe.Pointer) MLModel {
	return MLModel{gen_MLModel{
		objc.Object_FromPointer(ptr),
	}}
}

func MLModel_FromRef(ref objc.Ref) MLModel {
	return MLModel_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// PredictionFromFeaturesError generates a prediction from the feature values within the input feature provider.
//
// See https://developer.apple.com/documentation/coreml/mlmodel/2880280-predictionfromfeatures?language=objc for details.
func (x gen_MLModel) PredictionFromFeaturesError(
	input objc.Ref,
	error core.NSErrorRef,
) objc.Object {
	ret := C.MLModel_inst_PredictionFromFeaturesError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(input),
		objc.RefPointer(error),
	)

	return objc.Object_FromPointer(ret)
}

// PredictionsFromBatchError generates predictions for each input feature provider within the batch provider.
//
// See https://developer.apple.com/documentation/coreml/mlmodel/3088750-predictionsfrombatch?language=objc for details.
func (x gen_MLModel) PredictionsFromBatchError(
	inputBatch objc.Ref,
	error core.NSErrorRef,
) objc.Object {
	ret := C.MLModel_inst_PredictionsFromBatchError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(inputBatch),
		objc.RefPointer(error),
	)

	return objc.Object_FromPointer(ret)
}

// Init initializes a new instance of the MLModel class.
func (x gen_MLModel) Init() MLModel {
	ret := C.MLModel_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLModel_FromPointer(ret)
}

// Init_AsMLModel is a typed version of Init.
func (x gen_MLModel) Init_AsMLModel() MLModel {
	ret := C.MLModel_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLModel_FromPointer(ret)
}

type MLModelAssetRef interface {
	Pointer() uintptr
	Init_AsMLModelAsset() MLModelAsset
}

type gen_MLModelAsset struct {
	objc.Object
}

func MLModelAsset_FromPointer(ptr unsafe.Pointer) MLModelAsset {
	return MLModelAsset{gen_MLModelAsset{
		objc.Object_FromPointer(ptr),
	}}
}

func MLModelAsset_FromRef(ref objc.Ref) MLModelAsset {
	return MLModelAsset_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the MLModelAsset class.
func (x gen_MLModelAsset) Init() MLModelAsset {
	ret := C.MLModelAsset_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLModelAsset_FromPointer(ret)
}

// Init_AsMLModelAsset is a typed version of Init.
func (x gen_MLModelAsset) Init_AsMLModelAsset() MLModelAsset {
	ret := C.MLModelAsset_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLModelAsset_FromPointer(ret)
}

type MLModelCollectionRef interface {
	Pointer() uintptr
	Init_AsMLModelCollection() MLModelCollection
}

type gen_MLModelCollection struct {
	objc.Object
}

func MLModelCollection_FromPointer(ptr unsafe.Pointer) MLModelCollection {
	return MLModelCollection{gen_MLModelCollection{
		objc.Object_FromPointer(ptr),
	}}
}

func MLModelCollection_FromRef(ref objc.Ref) MLModelCollection {
	return MLModelCollection_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the MLModelCollection class.
func (x gen_MLModelCollection) Init() MLModelCollection {
	ret := C.MLModelCollection_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLModelCollection_FromPointer(ret)
}

// Init_AsMLModelCollection is a typed version of Init.
func (x gen_MLModelCollection) Init_AsMLModelCollection() MLModelCollection {
	ret := C.MLModelCollection_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLModelCollection_FromPointer(ret)
}

type MLNeuralEngineComputeDeviceRef interface {
	Pointer() uintptr
	Init_AsMLNeuralEngineComputeDevice() MLNeuralEngineComputeDevice
}

type gen_MLNeuralEngineComputeDevice struct {
	objc.Object
}

func MLNeuralEngineComputeDevice_FromPointer(ptr unsafe.Pointer) MLNeuralEngineComputeDevice {
	return MLNeuralEngineComputeDevice{gen_MLNeuralEngineComputeDevice{
		objc.Object_FromPointer(ptr),
	}}
}

func MLNeuralEngineComputeDevice_FromRef(ref objc.Ref) MLNeuralEngineComputeDevice {
	return MLNeuralEngineComputeDevice_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the MLNeuralEngineComputeDevice class.
func (x gen_MLNeuralEngineComputeDevice) Init() MLNeuralEngineComputeDevice {
	ret := C.MLNeuralEngineComputeDevice_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLNeuralEngineComputeDevice_FromPointer(ret)
}

// Init_AsMLNeuralEngineComputeDevice is a typed version of Init.
func (x gen_MLNeuralEngineComputeDevice) Init_AsMLNeuralEngineComputeDevice() MLNeuralEngineComputeDevice {
	ret := C.MLNeuralEngineComputeDevice_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return MLNeuralEngineComputeDevice_FromPointer(ret)
}

// TotalCoreCount is undocumented.
//
// See https://developer.apple.com/documentation/coreml/mlneuralenginecomputedevice/4278544-totalcorecount?language=objc for details.
func (x gen_MLNeuralEngineComputeDevice) TotalCoreCount() core.NSInteger {
	ret := C.MLNeuralEngineComputeDevice_inst_TotalCoreCount(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}
