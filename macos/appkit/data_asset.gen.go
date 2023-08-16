// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DataAsset] class.
var DataAssetClass = _DataAssetClass{objc.GetClass("NSDataAsset")}

type _DataAssetClass struct {
	objc.Class
}

// An interface definition for the [DataAsset] class.
type IDataAsset interface {
	objc.IObject
	Name() DataAssetName
	TypeIdentifier() string
	Data() []byte
}

// An object from a data set type stored in an asset catalog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdataasset?language=objc
type DataAsset struct {
	objc.Object
}

func DataAssetFrom(ptr unsafe.Pointer) DataAsset {
	return DataAsset{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DataAsset) InitWithName(name DataAssetName) DataAsset {
	rv := objc.Call[DataAsset](d_, objc.Sel("initWithName:"), name)
	return rv
}

// Initializes and returns an object with a reference to the named data asset in an asset catalog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdataasset/1403439-initwithname?language=objc
func DataAsset_InitWithName(name DataAssetName) DataAsset {
	return DataAssetClass.Alloc().InitWithName(name)
}

func (dc _DataAssetClass) Alloc() DataAsset {
	rv := objc.Call[DataAsset](dc, objc.Sel("alloc"))
	return rv
}

func DataAsset_Alloc() DataAsset {
	return DataAssetClass.Alloc()
}

func (dc _DataAssetClass) New() DataAsset {
	rv := objc.Call[DataAsset](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDataAsset() DataAsset {
	return DataAssetClass.New()
}

func (d_ DataAsset) Init() DataAsset {
	rv := objc.Call[DataAsset](d_, objc.Sel("init"))
	return rv
}

// The name of the data set in the asset catalog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdataasset/1403435-name?language=objc
func (d_ DataAsset) Name() DataAssetName {
	rv := objc.Call[DataAssetName](d_, objc.Sel("name"))
	return rv
}

// The uniform type identifier for the data asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdataasset/1403434-typeidentifier?language=objc
func (d_ DataAsset) TypeIdentifier() string {
	rv := objc.Call[string](d_, objc.Sel("typeIdentifier"))
	return rv
}

// The raw data values in the data asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdataasset/1403437-data?language=objc
func (d_ DataAsset) Data() []byte {
	rv := objc.Call[[]byte](d_, objc.Sel("data"))
	return rv
}