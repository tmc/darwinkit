// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Dimension] class.
var DimensionClass = _DimensionClass{objc.GetClass("NSDimension")}

type _DimensionClass struct {
	objc.Class
}

// An interface definition for the [Dimension] class.
type IDimension interface {
	IUnit
	Converter() UnitConverter
}

// An abstract class representing a dimensional unit of measure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension?language=objc
type Dimension struct {
	Unit
}

func DimensionFrom(ptr unsafe.Pointer) Dimension {
	return Dimension{
		Unit: UnitFrom(ptr),
	}
}

func (dc _DimensionClass) BaseUnit() Dimension {
	rv := objc.Call[Dimension](dc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func Dimension_BaseUnit() Dimension {
	return DimensionClass.BaseUnit()
}

func (d_ Dimension) InitWithSymbolConverter(symbol string, converter IUnitConverter) Dimension {
	rv := objc.Call[Dimension](d_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func Dimension_InitWithSymbolConverter(symbol string, converter IUnitConverter) Dimension {
	return DimensionClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (dc _DimensionClass) Alloc() Dimension {
	rv := objc.Call[Dimension](dc, objc.Sel("alloc"))
	return rv
}

func Dimension_Alloc() Dimension {
	return DimensionClass.Alloc()
}

func (dc _DimensionClass) New() Dimension {
	rv := objc.Call[Dimension](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDimension() Dimension {
	return DimensionClass.New()
}

func (d_ Dimension) Init() Dimension {
	rv := objc.Call[Dimension](d_, objc.Sel("init"))
	return rv
}

func (d_ Dimension) InitWithSymbol(symbol string) Dimension {
	rv := objc.Call[Dimension](d_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func Dimension_InitWithSymbol(symbol string) Dimension {
	return DimensionClass.Alloc().InitWithSymbol(symbol)
}

// The unit converter that represents the unit in terms of the dimension’s base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823516-converter?language=objc
func (d_ Dimension) Converter() UnitConverter {
	rv := objc.Call[UnitConverter](d_, objc.Sel("converter"))
	return rv
}