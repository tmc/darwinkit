// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Form] class.
var FormClass = _FormClass{objc.GetClass("NSForm")}

type _FormClass struct {
	objc.Class
}

// An interface definition for the [Form] class.
type IForm interface {
	IMatrix
}

// An NSForm object is a vertical matrix of NSFormCell objects to implement the fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsform?language=objc
type Form struct {
	Matrix
}

func FormFrom(ptr unsafe.Pointer) Form {
	return Form{
		Matrix: MatrixFrom(ptr),
	}
}

func (fc _FormClass) Alloc() Form {
	rv := objc.Call[Form](fc, objc.Sel("alloc"))
	return rv
}

func (fc _FormClass) New() Form {
	rv := objc.Call[Form](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewForm() Form {
	return FormClass.New()
}

func (f_ Form) Init() Form {
	rv := objc.Call[Form](f_, objc.Sel("init"))
	return rv
}

func (f_ Form) InitWithFrameModePrototypeNumberOfRowsNumberOfColumns(frameRect foundation.Rect, mode MatrixMode, cell ICell, rowsHigh int, colsWide int) Form {
	rv := objc.Call[Form](f_, objc.Sel("initWithFrame:mode:prototype:numberOfRows:numberOfColumns:"), frameRect, mode, cell, rowsHigh, colsWide)
	return rv
}

// Initializes and returns a newly allocated matrix of the specified size using the given cell as a prototype. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436386-initwithframe?language=objc
func NewFormWithFrameModePrototypeNumberOfRowsNumberOfColumns(frameRect foundation.Rect, mode MatrixMode, cell ICell, rowsHigh int, colsWide int) Form {
	instance := FormClass.Alloc().InitWithFrameModePrototypeNumberOfRowsNumberOfColumns(frameRect, mode, cell, rowsHigh, colsWide)
	instance.Autorelease()
	return instance
}

func (f_ Form) InitWithFrame(frameRect foundation.Rect) Form {
	rv := objc.Call[Form](f_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a newly allocated matrix with the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436428-initwithframe?language=objc
func NewFormWithFrame(frameRect foundation.Rect) Form {
	instance := FormClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

func (f_ Form) InitWithFrameModeCellClassNumberOfRowsNumberOfColumns(frameRect foundation.Rect, mode MatrixMode, factoryId objc.IClass, rowsHigh int, colsWide int) Form {
	rv := objc.Call[Form](f_, objc.Sel("initWithFrame:mode:cellClass:numberOfRows:numberOfColumns:"), frameRect, mode, factoryId, rowsHigh, colsWide)
	return rv
}

// Initializes and returns a newly allocated matrix of the specified size using cells of the given class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436400-initwithframe?language=objc
func NewFormWithFrameModeCellClassNumberOfRowsNumberOfColumns(frameRect foundation.Rect, mode MatrixMode, factoryId objc.IClass, rowsHigh int, colsWide int) Form {
	instance := FormClass.Alloc().InitWithFrameModeCellClassNumberOfRowsNumberOfColumns(frameRect, mode, factoryId, rowsHigh, colsWide)
	instance.Autorelease()
	return instance
}
