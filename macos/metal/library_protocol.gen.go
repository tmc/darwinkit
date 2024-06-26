// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A collection of Metal shader functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary?language=objc
type PLibrary interface {
	// optional
	NewFunctionWithNameConstantValuesCompletionHandler(name string, constantValues FunctionConstantValues, completionHandler func(function FunctionObject, error foundation.Error))
	HasNewFunctionWithNameConstantValuesCompletionHandler() bool

	// optional
	NewFunctionWithName(functionName string) FunctionObject
	HasNewFunctionWithName() bool

	// optional
	NewFunctionWithDescriptorCompletionHandler(descriptor FunctionDescriptor, completionHandler func(function FunctionObject, error foundation.Error))
	HasNewFunctionWithDescriptorCompletionHandler() bool

	// optional
	NewIntersectionFunctionWithDescriptorCompletionHandler(descriptor IntersectionFunctionDescriptor, completionHandler func(function FunctionObject, error foundation.Error))
	HasNewIntersectionFunctionWithDescriptorCompletionHandler() bool

	// optional
	NewIntersectionFunctionWithDescriptorError(descriptor IntersectionFunctionDescriptor, error unsafe.Pointer) FunctionObject
	HasNewIntersectionFunctionWithDescriptorError() bool

	// optional
	NewFunctionWithDescriptorError(descriptor FunctionDescriptor, error unsafe.Pointer) FunctionObject
	HasNewFunctionWithDescriptorError() bool

	// optional
	NewFunctionWithNameConstantValuesError(name string, constantValues FunctionConstantValues, error unsafe.Pointer) FunctionObject
	HasNewFunctionWithNameConstantValuesError() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	Device() DeviceObject
	HasDevice() bool

	// optional
	FunctionNames() []string
	HasFunctionNames() bool

	// optional
	Type() LibraryType
	HasType() bool

	// optional
	InstallName() string
	HasInstallName() bool
}

// ensure impl type implements protocol interface
var _ PLibrary = (*LibraryObject)(nil)

// A concrete type for the [PLibrary] protocol.
type LibraryObject struct {
	objc.Object
}

func (l_ LibraryObject) HasNewFunctionWithNameConstantValuesCompletionHandler() bool {
	return l_.RespondsToSelector(objc.Sel("newFunctionWithName:constantValues:completionHandler:"))
}

// Asynchronously creates a specialized shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1640053-newfunctionwithname?language=objc
func (l_ LibraryObject) NewFunctionWithNameConstantValuesCompletionHandler(name string, constantValues FunctionConstantValues, completionHandler func(function FunctionObject, error foundation.Error)) {
	objc.Call[objc.Void](l_, objc.Sel("newFunctionWithName:constantValues:completionHandler:"), name, constantValues, completionHandler)
}

func (l_ LibraryObject) HasNewFunctionWithName() bool {
	return l_.RespondsToSelector(objc.Sel("newFunctionWithName:"))
}

// Creates an object that represents a shader function in the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1515524-newfunctionwithname?language=objc
func (l_ LibraryObject) NewFunctionWithName(functionName string) FunctionObject {
	rv := objc.Call[FunctionObject](l_, objc.Sel("newFunctionWithName:"), functionName)
	return rv
}

func (l_ LibraryObject) HasNewFunctionWithDescriptorCompletionHandler() bool {
	return l_.RespondsToSelector(objc.Sel("newFunctionWithDescriptor:completionHandler:"))
}

// Asynchronously creates an object representing a shader function, using the specified descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3554040-newfunctionwithdescriptor?language=objc
func (l_ LibraryObject) NewFunctionWithDescriptorCompletionHandler(descriptor FunctionDescriptor, completionHandler func(function FunctionObject, error foundation.Error)) {
	objc.Call[objc.Void](l_, objc.Sel("newFunctionWithDescriptor:completionHandler:"), descriptor, completionHandler)
}

func (l_ LibraryObject) HasNewIntersectionFunctionWithDescriptorCompletionHandler() bool {
	return l_.RespondsToSelector(objc.Sel("newIntersectionFunctionWithDescriptor:completionHandler:"))
}

// Asynchronously creates an object representing a ray-tracing intersection function, using the specified descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3578134-newintersectionfunctionwithdescr?language=objc
func (l_ LibraryObject) NewIntersectionFunctionWithDescriptorCompletionHandler(descriptor IntersectionFunctionDescriptor, completionHandler func(function FunctionObject, error foundation.Error)) {
	objc.Call[objc.Void](l_, objc.Sel("newIntersectionFunctionWithDescriptor:completionHandler:"), descriptor, completionHandler)
}

func (l_ LibraryObject) HasNewIntersectionFunctionWithDescriptorError() bool {
	return l_.RespondsToSelector(objc.Sel("newIntersectionFunctionWithDescriptor:error:"))
}

// Synchronously creates an object representing a ray-tracing intersection function, using the specified descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3578135-newintersectionfunctionwithdescr?language=objc
func (l_ LibraryObject) NewIntersectionFunctionWithDescriptorError(descriptor IntersectionFunctionDescriptor, error unsafe.Pointer) FunctionObject {
	rv := objc.Call[FunctionObject](l_, objc.Sel("newIntersectionFunctionWithDescriptor:error:"), descriptor, error)
	return rv
}

func (l_ LibraryObject) HasNewFunctionWithDescriptorError() bool {
	return l_.RespondsToSelector(objc.Sel("newFunctionWithDescriptor:error:"))
}

// Synchronously creates an object representing a shader function, using the specified descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3554041-newfunctionwithdescriptor?language=objc
func (l_ LibraryObject) NewFunctionWithDescriptorError(descriptor FunctionDescriptor, error unsafe.Pointer) FunctionObject {
	rv := objc.Call[FunctionObject](l_, objc.Sel("newFunctionWithDescriptor:error:"), descriptor, error)
	return rv
}

func (l_ LibraryObject) HasNewFunctionWithNameConstantValuesError() bool {
	return l_.RespondsToSelector(objc.Sel("newFunctionWithName:constantValues:error:"))
}

// Synchronously creates a specialized shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1640020-newfunctionwithname?language=objc
func (l_ LibraryObject) NewFunctionWithNameConstantValuesError(name string, constantValues FunctionConstantValues, error unsafe.Pointer) FunctionObject {
	rv := objc.Call[FunctionObject](l_, objc.Sel("newFunctionWithName:constantValues:error:"), name, constantValues, error)
	return rv
}

func (l_ LibraryObject) HasSetLabel() bool {
	return l_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1516253-label?language=objc
func (l_ LibraryObject) SetLabel(value string) {
	objc.Call[objc.Void](l_, objc.Sel("setLabel:"), value)
}

func (l_ LibraryObject) HasLabel() bool {
	return l_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1516253-label?language=objc
func (l_ LibraryObject) Label() string {
	rv := objc.Call[string](l_, objc.Sel("label"))
	return rv
}

func (l_ LibraryObject) HasDevice() bool {
	return l_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1515661-device?language=objc
func (l_ LibraryObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](l_, objc.Sel("device"))
	return rv
}

func (l_ LibraryObject) HasFunctionNames() bool {
	return l_.RespondsToSelector(objc.Sel("functionNames"))
}

// The names of all public functions in the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1515651-functionnames?language=objc
func (l_ LibraryObject) FunctionNames() []string {
	rv := objc.Call[[]string](l_, objc.Sel("functionNames"))
	return rv
}

func (l_ LibraryObject) HasType() bool {
	return l_.RespondsToSelector(objc.Sel("type"))
}

// The library’s basic type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3554042-type?language=objc
func (l_ LibraryObject) Type() LibraryType {
	rv := objc.Call[LibraryType](l_, objc.Sel("type"))
	return rv
}

func (l_ LibraryObject) HasInstallName() bool {
	return l_.RespondsToSelector(objc.Sel("installName"))
}

// The installation name for a dynamic library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3554039-installname?language=objc
func (l_ LibraryObject) InstallName() string {
	rv := objc.Call[string](l_, objc.Sel("installName"))
	return rv
}
