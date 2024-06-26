// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MetadataObject] class.
var MetadataObjectClass = _MetadataObjectClass{objc.GetClass("AVMetadataObject")}

type _MetadataObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataObject] class.
type IMetadataObject interface {
	objc.IObject
	Bounds() coregraphics.Rect
	Time() coremedia.Time
	Duration() coremedia.Time
	Type() MetadataObjectType
}

// The abstract superclass for objects provided by a metadata capture output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobject?language=objc
type MetadataObject struct {
	objc.Object
}

func MetadataObjectFrom(ptr unsafe.Pointer) MetadataObject {
	return MetadataObject{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataObjectClass) Alloc() MetadataObject {
	rv := objc.Call[MetadataObject](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MetadataObjectClass) New() MetadataObject {
	rv := objc.Call[MetadataObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataObject() MetadataObject {
	return MetadataObjectClass.New()
}

func (m_ MetadataObject) Init() MetadataObject {
	rv := objc.Call[MetadataObject](m_, objc.Sel("init"))
	return rv
}

// The bounding rectangle associated with the metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobject/1386043-bounds?language=objc
func (m_ MetadataObject) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](m_, objc.Sel("bounds"))
	return rv
}

// The media time value associated with the metadata object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobject/1388593-time?language=objc
func (m_ MetadataObject) Time() coremedia.Time {
	rv := objc.Call[coremedia.Time](m_, objc.Sel("time"))
	return rv
}

// The duration of the media associated with this metadata object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobject/1386827-duration?language=objc
func (m_ MetadataObject) Duration() coremedia.Time {
	rv := objc.Call[coremedia.Time](m_, objc.Sel("duration"))
	return rv
}

// The type of metadata that this object provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobject/1387841-type?language=objc
func (m_ MetadataObject) Type() MetadataObjectType {
	rv := objc.Call[MetadataObjectType](m_, objc.Sel("type"))
	return rv
}
