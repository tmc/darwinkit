// Code generated by DarwinKit. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ModelConfiguration] class.
var ModelConfigurationClass = _ModelConfigurationClass{objc.GetClass("MLModelConfiguration")}

type _ModelConfigurationClass struct {
	objc.Class
}

// An interface definition for the [ModelConfiguration] class.
type IModelConfiguration interface {
	objc.IObject
	PreferredMetalDevice() metal.DeviceObject
	SetPreferredMetalDevice(value metal.PDevice)
	SetPreferredMetalDeviceObject(valueObject objc.IObject)
	Parameters() foundation.Dictionary
	SetParameters(value foundation.Dictionary)
	AllowLowPrecisionAccumulationOnGPU() bool
	SetAllowLowPrecisionAccumulationOnGPU(value bool)
	ComputeUnits() ComputeUnits
	SetComputeUnits(value ComputeUnits)
}

// The settings for creating or updating a machine learning model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration?language=objc
type ModelConfiguration struct {
	objc.Object
}

func ModelConfigurationFrom(ptr unsafe.Pointer) ModelConfiguration {
	return ModelConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _ModelConfigurationClass) Alloc() ModelConfiguration {
	rv := objc.Call[ModelConfiguration](mc, objc.Sel("alloc"))
	return rv
}

func (mc _ModelConfigurationClass) New() ModelConfiguration {
	rv := objc.Call[ModelConfiguration](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModelConfiguration() ModelConfiguration {
	return ModelConfigurationClass.New()
}

func (m_ ModelConfiguration) Init() ModelConfiguration {
	rv := objc.Call[ModelConfiguration](m_, objc.Sel("init"))
	return rv
}

// The metal device you prefer this model use to make predictions (inference) and update the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3222915-preferredmetaldevice?language=objc
func (m_ ModelConfiguration) PreferredMetalDevice() metal.DeviceObject {
	rv := objc.Call[metal.DeviceObject](m_, objc.Sel("preferredMetalDevice"))
	return rv
}

// The metal device you prefer this model use to make predictions (inference) and update the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3222915-preferredmetaldevice?language=objc
func (m_ ModelConfiguration) SetPreferredMetalDevice(value metal.PDevice) {
	po0 := objc.WrapAsProtocol("MTLDevice", value)
	objc.Call[objc.Void](m_, objc.Sel("setPreferredMetalDevice:"), po0)
}

// The metal device you prefer this model use to make predictions (inference) and update the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3222915-preferredmetaldevice?language=objc
func (m_ ModelConfiguration) SetPreferredMetalDeviceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredMetalDevice:"), valueObject)
}

// A dictionary of configuration settings your app can override when loading a model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3333249-parameters?language=objc
func (m_ ModelConfiguration) Parameters() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](m_, objc.Sel("parameters"))
	return rv
}

// A dictionary of configuration settings your app can override when loading a model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3333249-parameters?language=objc
func (m_ ModelConfiguration) SetParameters(value foundation.Dictionary) {
	objc.Call[objc.Void](m_, objc.Sel("setParameters:"), value)
}

// A Boolean value that determines whether to allow low-precision accumulation on a GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3222914-allowlowprecisionaccumulationong?language=objc
func (m_ ModelConfiguration) AllowLowPrecisionAccumulationOnGPU() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowLowPrecisionAccumulationOnGPU"))
	return rv
}

// A Boolean value that determines whether to allow low-precision accumulation on a GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3222914-allowlowprecisionaccumulationong?language=objc
func (m_ ModelConfiguration) SetAllowLowPrecisionAccumulationOnGPU(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowLowPrecisionAccumulationOnGPU:"), value)
}

// The processing unit or units the model uses to make predictions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3022235-computeunits?language=objc
func (m_ ModelConfiguration) ComputeUnits() ComputeUnits {
	rv := objc.Call[ComputeUnits](m_, objc.Sel("computeUnits"))
	return rv
}

// The processing unit or units the model uses to make predictions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/3022235-computeunits?language=objc
func (m_ ModelConfiguration) SetComputeUnits(value ComputeUnits) {
	objc.Call[objc.Void](m_, objc.Sel("setComputeUnits:"), value)
}
