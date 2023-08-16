// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewGridLayout] class.
var CollectionViewGridLayoutClass = _CollectionViewGridLayoutClass{objc.GetClass("NSCollectionViewGridLayout")}

type _CollectionViewGridLayoutClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewGridLayout] class.
type ICollectionViewGridLayout interface {
	ICollectionViewLayout
	BackgroundColors() []Color
	SetBackgroundColors(value []IColor)
	MinimumItemSize() foundation.Size
	SetMinimumItemSize(value foundation.Size)
	MaximumNumberOfRows() uint
	SetMaximumNumberOfRows(value uint)
	MaximumItemSize() foundation.Size
	SetMaximumItemSize(value foundation.Size)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	Margins() foundation.EdgeInsets
	SetMargins(value foundation.EdgeInsets)
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
}

// A layout that displays a single section of items in a row and column grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout?language=objc
type CollectionViewGridLayout struct {
	CollectionViewLayout
}

func CollectionViewGridLayoutFrom(ptr unsafe.Pointer) CollectionViewGridLayout {
	return CollectionViewGridLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}

func (cc _CollectionViewGridLayoutClass) Alloc() CollectionViewGridLayout {
	rv := objc.Call[CollectionViewGridLayout](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewGridLayout_Alloc() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.Alloc()
}

func (cc _CollectionViewGridLayoutClass) New() CollectionViewGridLayout {
	rv := objc.Call[CollectionViewGridLayout](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewGridLayout() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.New()
}

func (c_ CollectionViewGridLayout) Init() CollectionViewGridLayout {
	rv := objc.Call[CollectionViewGridLayout](c_, objc.Sel("init"))
	return rv
}

// The array of background colors to use when drawing the grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1530955-backgroundcolors?language=objc
func (c_ CollectionViewGridLayout) BackgroundColors() []Color {
	rv := objc.Call[[]Color](c_, objc.Sel("backgroundColors"))
	return rv
}

// The array of background colors to use when drawing the grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1530955-backgroundcolors?language=objc
func (c_ CollectionViewGridLayout) SetBackgroundColors(value []IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundColors:"), value)
}

// The smallest allowable size for an item’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1534425-minimumitemsize?language=objc
func (c_ CollectionViewGridLayout) MinimumItemSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("minimumItemSize"))
	return rv
}

// The smallest allowable size for an item’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1534425-minimumitemsize?language=objc
func (c_ CollectionViewGridLayout) SetMinimumItemSize(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumItemSize:"), value)
}

// The maximum number of rows to display in the collection view’s visible area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1524389-maximumnumberofrows?language=objc
func (c_ CollectionViewGridLayout) MaximumNumberOfRows() uint {
	rv := objc.Call[uint](c_, objc.Sel("maximumNumberOfRows"))
	return rv
}

// The maximum number of rows to display in the collection view’s visible area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1524389-maximumnumberofrows?language=objc
func (c_ CollectionViewGridLayout) SetMaximumNumberOfRows(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setMaximumNumberOfRows:"), value)
}

// The largest allowable size for an item’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1530923-maximumitemsize?language=objc
func (c_ CollectionViewGridLayout) MaximumItemSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("maximumItemSize"))
	return rv
}

// The largest allowable size for an item’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1530923-maximumitemsize?language=objc
func (c_ CollectionViewGridLayout) SetMaximumItemSize(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setMaximumItemSize:"), value)
}

// The minimum spacing (in points) to use between items in the same row or column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1525116-minimuminteritemspacing?language=objc
func (c_ CollectionViewGridLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minimumInteritemSpacing"))
	return rv
}

// The minimum spacing (in points) to use between items in the same row or column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1525116-minimuminteritemspacing?language=objc
func (c_ CollectionViewGridLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumInteritemSpacing:"), value)
}

// The amount of empty space (in points) around the grid’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1527362-margins?language=objc
func (c_ CollectionViewGridLayout) Margins() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](c_, objc.Sel("margins"))
	return rv
}

// The amount of empty space (in points) around the grid’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1527362-margins?language=objc
func (c_ CollectionViewGridLayout) SetMargins(value foundation.EdgeInsets) {
	objc.Call[objc.Void](c_, objc.Sel("setMargins:"), value)
}

// The maximum number of columns to display in the collection view’s visible area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1533264-maximumnumberofcolumns?language=objc
func (c_ CollectionViewGridLayout) MaximumNumberOfColumns() uint {
	rv := objc.Call[uint](c_, objc.Sel("maximumNumberOfColumns"))
	return rv
}

// The maximum number of columns to display in the collection view’s visible area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1533264-maximumnumberofcolumns?language=objc
func (c_ CollectionViewGridLayout) SetMaximumNumberOfColumns(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setMaximumNumberOfColumns:"), value)
}

// The minimum spacing (in points) to use between rows or columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1535114-minimumlinespacing?language=objc
func (c_ CollectionViewGridLayout) MinimumLineSpacing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minimumLineSpacing"))
	return rv
}

// The minimum spacing (in points) to use between rows or columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewgridlayout/1535114-minimumlinespacing?language=objc
func (c_ CollectionViewGridLayout) SetMinimumLineSpacing(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumLineSpacing:"), value)
}