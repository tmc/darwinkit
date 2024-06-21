// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A set of optional methods that a popover delegate can implement to provide additional or custom functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate?language=objc
type PPopoverDelegate interface {
	// optional
	PopoverWillShow(notification foundation.Notification)
	HasPopoverWillShow() bool

	// optional
	PopoverWillClose(notification foundation.Notification)
	HasPopoverWillClose() bool

	// optional
	PopoverDidShow(notification foundation.Notification)
	HasPopoverDidShow() bool

	// optional
	PopoverDidDetach(popover Popover)
	HasPopoverDidDetach() bool

	// optional
	DetachableWindowForPopover(popover Popover) Window
	HasDetachableWindowForPopover() bool

	// optional
	PopoverDidClose(notification foundation.Notification)
	HasPopoverDidClose() bool

	// optional
	PopoverShouldClose(popover Popover) bool
	HasPopoverShouldClose() bool

	// optional
	PopoverShouldDetach(popover Popover) bool
	HasPopoverShouldDetach() bool
}

// A delegate implementation builder for the [PPopoverDelegate] protocol.
type PopoverDelegate struct {
	_PopoverWillShow            func(notification foundation.Notification)
	_PopoverWillClose           func(notification foundation.Notification)
	_PopoverDidShow             func(notification foundation.Notification)
	_PopoverDidDetach           func(popover Popover)
	_DetachableWindowForPopover func(popover Popover) Window
	_PopoverDidClose            func(notification foundation.Notification)
	_PopoverShouldClose         func(popover Popover) bool
	_PopoverShouldDetach        func(popover Popover) bool
}

func (di *PopoverDelegate) HasPopoverWillShow() bool {
	return di._PopoverWillShow != nil
}

// Invoked when the popover will show. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532556-popoverwillshow?language=objc
func (di *PopoverDelegate) SetPopoverWillShow(f func(notification foundation.Notification)) {
	di._PopoverWillShow = f
}

// Invoked when the popover will show. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532556-popoverwillshow?language=objc
func (di *PopoverDelegate) PopoverWillShow(notification foundation.Notification) {
	di._PopoverWillShow(notification)
}
func (di *PopoverDelegate) HasPopoverWillClose() bool {
	return di._PopoverWillClose != nil
}

// Invoked when the popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1535119-popoverwillclose?language=objc
func (di *PopoverDelegate) SetPopoverWillClose(f func(notification foundation.Notification)) {
	di._PopoverWillClose = f
}

// Invoked when the popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1535119-popoverwillclose?language=objc
func (di *PopoverDelegate) PopoverWillClose(notification foundation.Notification) {
	di._PopoverWillClose(notification)
}
func (di *PopoverDelegate) HasPopoverDidShow() bool {
	return di._PopoverDidShow != nil
}

// Invoked when the popover has been shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1533573-popoverdidshow?language=objc
func (di *PopoverDelegate) SetPopoverDidShow(f func(notification foundation.Notification)) {
	di._PopoverDidShow = f
}

// Invoked when the popover has been shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1533573-popoverdidshow?language=objc
func (di *PopoverDelegate) PopoverDidShow(notification foundation.Notification) {
	di._PopoverDidShow(notification)
}
func (di *PopoverDelegate) HasPopoverDidDetach() bool {
	return di._PopoverDidDetach != nil
}

// Indicates that a popover has been released while it's in an implicitly detached state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1524674-popoverdiddetach?language=objc
func (di *PopoverDelegate) SetPopoverDidDetach(f func(popover Popover)) {
	di._PopoverDidDetach = f
}

// Indicates that a popover has been released while it's in an implicitly detached state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1524674-popoverdiddetach?language=objc
func (di *PopoverDelegate) PopoverDidDetach(popover Popover) {
	di._PopoverDidDetach(popover)
}
func (di *PopoverDelegate) HasDetachableWindowForPopover() bool {
	return di._DetachableWindowForPopover != nil
}

// Detaches the popover creating a window containing the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1534822-detachablewindowforpopover?language=objc
func (di *PopoverDelegate) SetDetachableWindowForPopover(f func(popover Popover) Window) {
	di._DetachableWindowForPopover = f
}

// Detaches the popover creating a window containing the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1534822-detachablewindowforpopover?language=objc
func (di *PopoverDelegate) DetachableWindowForPopover(popover Popover) Window {
	return di._DetachableWindowForPopover(popover)
}
func (di *PopoverDelegate) HasPopoverDidClose() bool {
	return di._PopoverDidClose != nil
}

// Invoked when the popover did close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1526581-popoverdidclose?language=objc
func (di *PopoverDelegate) SetPopoverDidClose(f func(notification foundation.Notification)) {
	di._PopoverDidClose = f
}

// Invoked when the popover did close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1526581-popoverdidclose?language=objc
func (di *PopoverDelegate) PopoverDidClose(notification foundation.Notification) {
	di._PopoverDidClose(notification)
}
func (di *PopoverDelegate) HasPopoverShouldClose() bool {
	return di._PopoverShouldClose != nil
}

// Allows a delegate to override a close request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532593-popovershouldclose?language=objc
func (di *PopoverDelegate) SetPopoverShouldClose(f func(popover Popover) bool) {
	di._PopoverShouldClose = f
}

// Allows a delegate to override a close request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532593-popovershouldclose?language=objc
func (di *PopoverDelegate) PopoverShouldClose(popover Popover) bool {
	return di._PopoverShouldClose(popover)
}
func (di *PopoverDelegate) HasPopoverShouldDetach() bool {
	return di._PopoverShouldDetach != nil
}

// Returns a Boolean value that indicates whether a popover should detach from its positioning view and become a separate window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1529911-popovershoulddetach?language=objc
func (di *PopoverDelegate) SetPopoverShouldDetach(f func(popover Popover) bool) {
	di._PopoverShouldDetach = f
}

// Returns a Boolean value that indicates whether a popover should detach from its positioning view and become a separate window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1529911-popovershoulddetach?language=objc
func (di *PopoverDelegate) PopoverShouldDetach(popover Popover) bool {
	return di._PopoverShouldDetach(popover)
}

// ensure impl type implements protocol interface
var _ PPopoverDelegate = (*PopoverDelegateObject)(nil)

// A concrete type for the [PPopoverDelegate] protocol.
type PopoverDelegateObject struct {
	objc.Object
}

func (p_ PopoverDelegateObject) HasPopoverWillShow() bool {
	return p_.RespondsToSelector(objc.Sel("popoverWillShow:"))
}

// Invoked when the popover will show. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532556-popoverwillshow?language=objc
func (p_ PopoverDelegateObject) PopoverWillShow(notification foundation.Notification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverWillShow:"), notification)
}

func (p_ PopoverDelegateObject) HasPopoverWillClose() bool {
	return p_.RespondsToSelector(objc.Sel("popoverWillClose:"))
}

// Invoked when the popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1535119-popoverwillclose?language=objc
func (p_ PopoverDelegateObject) PopoverWillClose(notification foundation.Notification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverWillClose:"), notification)
}

func (p_ PopoverDelegateObject) HasPopoverDidShow() bool {
	return p_.RespondsToSelector(objc.Sel("popoverDidShow:"))
}

// Invoked when the popover has been shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1533573-popoverdidshow?language=objc
func (p_ PopoverDelegateObject) PopoverDidShow(notification foundation.Notification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverDidShow:"), notification)
}

func (p_ PopoverDelegateObject) HasPopoverDidDetach() bool {
	return p_.RespondsToSelector(objc.Sel("popoverDidDetach:"))
}

// Indicates that a popover has been released while it's in an implicitly detached state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1524674-popoverdiddetach?language=objc
func (p_ PopoverDelegateObject) PopoverDidDetach(popover Popover) {
	objc.Call[objc.Void](p_, objc.Sel("popoverDidDetach:"), popover)
}

func (p_ PopoverDelegateObject) HasDetachableWindowForPopover() bool {
	return p_.RespondsToSelector(objc.Sel("detachableWindowForPopover:"))
}

// Detaches the popover creating a window containing the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1534822-detachablewindowforpopover?language=objc
func (p_ PopoverDelegateObject) DetachableWindowForPopover(popover Popover) Window {
	rv := objc.Call[Window](p_, objc.Sel("detachableWindowForPopover:"), popover)
	return rv
}

func (p_ PopoverDelegateObject) HasPopoverDidClose() bool {
	return p_.RespondsToSelector(objc.Sel("popoverDidClose:"))
}

// Invoked when the popover did close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1526581-popoverdidclose?language=objc
func (p_ PopoverDelegateObject) PopoverDidClose(notification foundation.Notification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverDidClose:"), notification)
}

func (p_ PopoverDelegateObject) HasPopoverShouldClose() bool {
	return p_.RespondsToSelector(objc.Sel("popoverShouldClose:"))
}

// Allows a delegate to override a close request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532593-popovershouldclose?language=objc
func (p_ PopoverDelegateObject) PopoverShouldClose(popover Popover) bool {
	rv := objc.Call[bool](p_, objc.Sel("popoverShouldClose:"), popover)
	return rv
}

func (p_ PopoverDelegateObject) HasPopoverShouldDetach() bool {
	return p_.RespondsToSelector(objc.Sel("popoverShouldDetach:"))
}

// Returns a Boolean value that indicates whether a popover should detach from its positioning view and become a separate window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1529911-popovershoulddetach?language=objc
func (p_ PopoverDelegateObject) PopoverShouldDetach(popover Popover) bool {
	rv := objc.Call[bool](p_, objc.Sel("popoverShouldDetach:"), popover)
	return rv
}
