// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LevelIndicatorCell] class.
var LevelIndicatorCellClass = _LevelIndicatorCellClass{objc.GetClass("NSLevelIndicatorCell")}

type _LevelIndicatorCellClass struct {
	objc.Class
}

// An interface definition for the [LevelIndicatorCell] class.
type ILevelIndicatorCell interface {
	IActionCell
	RectOfTickMarkAtIndex(index int) foundation.Rect
	TickMarkValueAtIndex(index int) float64
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	CriticalValue() float64
	SetCriticalValue(value float64)
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	WarningValue() float64
	SetWarningValue(value float64)
	LevelIndicatorStyle() LevelIndicatorStyle
	SetLevelIndicatorStyle(value LevelIndicatorStyle)
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
}

// NSLevelIndicatorCell is a subclass of NSActionCell that provides several level indicator display styles including: capacity, ranking and relevancy. The capacity style provides both continuous and discrete modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell?language=objc
type LevelIndicatorCell struct {
	ActionCell
}

func LevelIndicatorCellFrom(ptr unsafe.Pointer) LevelIndicatorCell {
	return LevelIndicatorCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (l_ LevelIndicatorCell) InitWithLevelIndicatorStyle(levelIndicatorStyle LevelIndicatorStyle) LevelIndicatorCell {
	rv := objc.Call[LevelIndicatorCell](l_, objc.Sel("initWithLevelIndicatorStyle:"), levelIndicatorStyle)
	return rv
}

// Initializes the receiver with the style specified by levelIndicatorStyle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1527498-initwithlevelindicatorstyle?language=objc
func LevelIndicatorCell_InitWithLevelIndicatorStyle(levelIndicatorStyle LevelIndicatorStyle) LevelIndicatorCell {
	return LevelIndicatorCellClass.Alloc().InitWithLevelIndicatorStyle(levelIndicatorStyle)
}

func (lc _LevelIndicatorCellClass) Alloc() LevelIndicatorCell {
	rv := objc.Call[LevelIndicatorCell](lc, objc.Sel("alloc"))
	return rv
}

func LevelIndicatorCell_Alloc() LevelIndicatorCell {
	return LevelIndicatorCellClass.Alloc()
}

func (lc _LevelIndicatorCellClass) New() LevelIndicatorCell {
	rv := objc.Call[LevelIndicatorCell](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLevelIndicatorCell() LevelIndicatorCell {
	return LevelIndicatorCellClass.New()
}

func (l_ LevelIndicatorCell) Init() LevelIndicatorCell {
	rv := objc.Call[LevelIndicatorCell](l_, objc.Sel("init"))
	return rv
}

func (l_ LevelIndicatorCell) InitImageCell(image IImage) LevelIndicatorCell {
	rv := objc.Call[LevelIndicatorCell](l_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func LevelIndicatorCell_InitImageCell(image IImage) LevelIndicatorCell {
	return LevelIndicatorCellClass.Alloc().InitImageCell(image)
}

func (l_ LevelIndicatorCell) InitTextCell(string_ string) LevelIndicatorCell {
	rv := objc.Call[LevelIndicatorCell](l_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func LevelIndicatorCell_InitTextCell(string_ string) LevelIndicatorCell {
	return LevelIndicatorCellClass.Alloc().InitTextCell(string_)
}

// Returns the bounding rectangle of the tick mark identified by index (the minimum-value tick mark is at index 0). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1534530-rectoftickmarkatindex?language=objc
func (l_ LevelIndicatorCell) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.Call[foundation.Rect](l_, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}

// Returns the receiver’s value represented by the tick mark at index (the minimum-value tick mark has an index of 0). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1535863-tickmarkvalueatindex?language=objc
func (l_ LevelIndicatorCell) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Call[float64](l_, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}

// The placement of tick marks on the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1532399-tickmarkposition?language=objc
func (l_ LevelIndicatorCell) TickMarkPosition() TickMarkPosition {
	rv := objc.Call[TickMarkPosition](l_, objc.Sel("tickMarkPosition"))
	return rv
}

// The placement of tick marks on the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1532399-tickmarkposition?language=objc
func (l_ LevelIndicatorCell) SetTickMarkPosition(value TickMarkPosition) {
	objc.Call[objc.Void](l_, objc.Sel("setTickMarkPosition:"), value)
}

// The critical value of the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1525337-criticalvalue?language=objc
func (l_ LevelIndicatorCell) CriticalValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("criticalValue"))
	return rv
}

// The critical value of the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1525337-criticalvalue?language=objc
func (l_ LevelIndicatorCell) SetCriticalValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setCriticalValue:"), value)
}

// The number of major tick marks displayed by the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1528987-numberofmajortickmarks?language=objc
func (l_ LevelIndicatorCell) NumberOfMajorTickMarks() int {
	rv := objc.Call[int](l_, objc.Sel("numberOfMajorTickMarks"))
	return rv
}

// The number of major tick marks displayed by the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1528987-numberofmajortickmarks?language=objc
func (l_ LevelIndicatorCell) SetNumberOfMajorTickMarks(value int) {
	objc.Call[objc.Void](l_, objc.Sel("setNumberOfMajorTickMarks:"), value)
}

// The warning value of the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1528974-warningvalue?language=objc
func (l_ LevelIndicatorCell) WarningValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("warningValue"))
	return rv
}

// The warning value of the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1528974-warningvalue?language=objc
func (l_ LevelIndicatorCell) SetWarningValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setWarningValue:"), value)
}

// The style of the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1531954-levelindicatorstyle?language=objc
func (l_ LevelIndicatorCell) LevelIndicatorStyle() LevelIndicatorStyle {
	rv := objc.Call[LevelIndicatorStyle](l_, objc.Sel("levelIndicatorStyle"))
	return rv
}

// The style of the level indicator control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1531954-levelindicatorstyle?language=objc
func (l_ LevelIndicatorCell) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	objc.Call[objc.Void](l_, objc.Sel("setLevelIndicatorStyle:"), value)
}

// The minimum value of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1534472-minvalue?language=objc
func (l_ LevelIndicatorCell) MinValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("minValue"))
	return rv
}

// The minimum value of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1534472-minvalue?language=objc
func (l_ LevelIndicatorCell) SetMinValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setMinValue:"), value)
}

// The maximum value of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1528309-maxvalue?language=objc
func (l_ LevelIndicatorCell) MaxValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("maxValue"))
	return rv
}

// The maximum value of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1528309-maxvalue?language=objc
func (l_ LevelIndicatorCell) SetMaxValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setMaxValue:"), value)
}

// The number of tick marks displayed by the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1534680-numberoftickmarks?language=objc
func (l_ LevelIndicatorCell) NumberOfTickMarks() int {
	rv := objc.Call[int](l_, objc.Sel("numberOfTickMarks"))
	return rv
}

// The number of tick marks displayed by the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/1534680-numberoftickmarks?language=objc
func (l_ LevelIndicatorCell) SetNumberOfTickMarks(value int) {
	objc.Call[objc.Void](l_, objc.Sel("setNumberOfTickMarks:"), value)
}