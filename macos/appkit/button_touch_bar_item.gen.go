// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ButtonTouchBarItem] class.
var ButtonTouchBarItemClass = _ButtonTouchBarItemClass{objc.GetClass("NSButtonTouchBarItem")}

type _ButtonTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [ButtonTouchBarItem] class.
type IButtonTouchBarItem interface {
	ITouchBarItem
	SetCustomizationLabel(value string)
	IsEnabled() bool
	SetEnabled(value bool)
	Target() objc.Object
	SetTarget(value objc.IObject)
	Title() string
	SetTitle(value string)
	Image() Image
	SetImage(value IImage)
	Action() objc.Selector
	SetAction(value objc.Selector)
	BezelColor() Color
	SetBezelColor(value IColor)
}

// A bar item that provides a button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem?language=objc
type ButtonTouchBarItem struct {
	TouchBarItem
}

func ButtonTouchBarItemFrom(ptr unsafe.Pointer) ButtonTouchBarItem {
	return ButtonTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierTitleTargetAction(identifier TouchBarItemIdentifier, title string, target objc.IObject, action objc.Selector) ButtonTouchBarItem {
	rv := objc.Call[ButtonTouchBarItem](bc, objc.Sel("buttonTouchBarItemWithIdentifier:title:target:action:"), identifier, title, target, action)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237180-buttontouchbaritemwithidentifier?language=objc
func ButtonTouchBarItem_ButtonTouchBarItemWithIdentifierTitleTargetAction(identifier TouchBarItemIdentifier, title string, target objc.IObject, action objc.Selector) ButtonTouchBarItem {
	return ButtonTouchBarItemClass.ButtonTouchBarItemWithIdentifierTitleTargetAction(identifier, title, target, action)
}

func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier TouchBarItemIdentifier, title string, image IImage, target objc.IObject, action objc.Selector) ButtonTouchBarItem {
	rv := objc.Call[ButtonTouchBarItem](bc, objc.Sel("buttonTouchBarItemWithIdentifier:title:image:target:action:"), identifier, title, image, target, action)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237179-buttontouchbaritemwithidentifier?language=objc
func ButtonTouchBarItem_ButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier TouchBarItemIdentifier, title string, image IImage, target objc.IObject, action objc.Selector) ButtonTouchBarItem {
	return ButtonTouchBarItemClass.ButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier, title, image, target, action)
}

func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierImageTargetAction(identifier TouchBarItemIdentifier, image IImage, target objc.IObject, action objc.Selector) ButtonTouchBarItem {
	rv := objc.Call[ButtonTouchBarItem](bc, objc.Sel("buttonTouchBarItemWithIdentifier:image:target:action:"), identifier, image, target, action)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237178-buttontouchbaritemwithidentifier?language=objc
func ButtonTouchBarItem_ButtonTouchBarItemWithIdentifierImageTargetAction(identifier TouchBarItemIdentifier, image IImage, target objc.IObject, action objc.Selector) ButtonTouchBarItem {
	return ButtonTouchBarItemClass.ButtonTouchBarItemWithIdentifierImageTargetAction(identifier, image, target, action)
}

func (bc _ButtonTouchBarItemClass) Alloc() ButtonTouchBarItem {
	rv := objc.Call[ButtonTouchBarItem](bc, objc.Sel("alloc"))
	return rv
}

func (bc _ButtonTouchBarItemClass) New() ButtonTouchBarItem {
	rv := objc.Call[ButtonTouchBarItem](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewButtonTouchBarItem() ButtonTouchBarItem {
	return ButtonTouchBarItemClass.New()
}

func (b_ ButtonTouchBarItem) Init() ButtonTouchBarItem {
	rv := objc.Call[ButtonTouchBarItem](b_, objc.Sel("init"))
	return rv
}

func (b_ ButtonTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) ButtonTouchBarItem {
	rv := objc.Call[ButtonTouchBarItem](b_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func NewButtonTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) ButtonTouchBarItem {
	instance := ButtonTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237181-customizationlabel?language=objc
func (b_ ButtonTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setCustomizationLabel:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3353052-enabled?language=objc
func (b_ ButtonTouchBarItem) IsEnabled() bool {
	rv := objc.Call[bool](b_, objc.Sel("isEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3353052-enabled?language=objc
func (b_ ButtonTouchBarItem) SetEnabled(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setEnabled:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237184-target?language=objc
func (b_ ButtonTouchBarItem) Target() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("target"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237184-target?language=objc
func (b_ ButtonTouchBarItem) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setTarget:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237185-title?language=objc
func (b_ ButtonTouchBarItem) Title() string {
	rv := objc.Call[string](b_, objc.Sel("title"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237185-title?language=objc
func (b_ ButtonTouchBarItem) SetTitle(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setTitle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237182-image?language=objc
func (b_ ButtonTouchBarItem) Image() Image {
	rv := objc.Call[Image](b_, objc.Sel("image"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237182-image?language=objc
func (b_ ButtonTouchBarItem) SetImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setImage:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237176-action?language=objc
func (b_ ButtonTouchBarItem) Action() objc.Selector {
	rv := objc.Call[objc.Selector](b_, objc.Sel("action"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237176-action?language=objc
func (b_ ButtonTouchBarItem) SetAction(value objc.Selector) {
	objc.Call[objc.Void](b_, objc.Sel("setAction:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237177-bezelcolor?language=objc
func (b_ ButtonTouchBarItem) BezelColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("bezelColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/3237177-bezelcolor?language=objc
func (b_ ButtonTouchBarItem) SetBezelColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setBezelColor:"), value)
}
