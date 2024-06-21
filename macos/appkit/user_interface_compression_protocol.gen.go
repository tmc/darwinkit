// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A protocol that describes how a UI control should redisplay when space is restricted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompression?language=objc
type PUserInterfaceCompression interface {
	// optional
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) foundation.Size
	HasMinimumSizeWithPrioritizedCompressionOptions() bool

	// optional
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	HasCompressWithPrioritizedCompressionOptions() bool

	// optional
	ActiveCompressionOptions() UserInterfaceCompressionOptions
	HasActiveCompressionOptions() bool
}

// ensure impl type implements protocol interface
var _ PUserInterfaceCompression = (*UserInterfaceCompressionObject)(nil)

// A concrete type for the [PUserInterfaceCompression] protocol.
type UserInterfaceCompressionObject struct {
	objc.Object
}

func (u_ UserInterfaceCompressionObject) HasMinimumSizeWithPrioritizedCompressionOptions() bool {
	return u_.RespondsToSelector(objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"))
}

// Returns the minimum size a view can achieve by applying the supplied compression options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompression/2909973-minimumsizewithprioritizedcompre?language=objc
func (u_ UserInterfaceCompressionObject) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) foundation.Size {
	rv := objc.Call[foundation.Size](u_, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

func (u_ UserInterfaceCompressionObject) HasCompressWithPrioritizedCompressionOptions() bool {
	return u_.RespondsToSelector(objc.Sel("compressWithPrioritizedCompressionOptions:"))
}

// Compress the view by applying the specified compression options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompression/2909978-compresswithprioritizedcompressi?language=objc
func (u_ UserInterfaceCompressionObject) CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) {
	objc.Call[objc.Void](u_, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

func (u_ UserInterfaceCompressionObject) HasActiveCompressionOptions() bool {
	return u_.RespondsToSelector(objc.Sel("activeCompressionOptions"))
}

// The compression options that are currently applied to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompression/2909984-activecompressionoptions?language=objc
func (u_ UserInterfaceCompressionObject) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](u_, objc.Sel("activeCompressionOptions"))
	return rv
}
