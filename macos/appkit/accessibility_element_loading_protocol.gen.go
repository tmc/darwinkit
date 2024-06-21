// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to support loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelementloading?language=objc
type PAccessibilityElementLoading interface {
	// optional
	AccessibilityElementWithToken(token AccessibilityLoadingToken) objc.Object
	HasAccessibilityElementWithToken() bool

	// optional
	AccessibilityRangeInTargetElementWithToken(token AccessibilityLoadingToken) foundation.Range
	HasAccessibilityRangeInTargetElementWithToken() bool
}

// ensure impl type implements protocol interface
var _ PAccessibilityElementLoading = (*AccessibilityElementLoadingObject)(nil)

// A concrete type for the [PAccessibilityElementLoading] protocol.
type AccessibilityElementLoadingObject struct {
	objc.Object
}

func (a_ AccessibilityElementLoadingObject) HasAccessibilityElementWithToken() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityElementWithToken:"))
}

// Loads the target accessibility element with the specified load token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelementloading/2890815-accessibilityelementwithtoken?language=objc
func (a_ AccessibilityElementLoadingObject) AccessibilityElementWithToken(token AccessibilityLoadingToken) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityElementWithToken:"), token)
	return rv
}

func (a_ AccessibilityElementLoadingObject) HasAccessibilityRangeInTargetElementWithToken() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRangeInTargetElementWithToken:"))
}

// Returns the range that specifies the area of interest in text-based accessibility elements with the specified load token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelementloading/2890818-accessibilityrangeintargetelemen?language=objc
func (a_ AccessibilityElementLoadingObject) AccessibilityRangeInTargetElementWithToken(token AccessibilityLoadingToken) foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityRangeInTargetElementWithToken:"), token)
	return rv
}
