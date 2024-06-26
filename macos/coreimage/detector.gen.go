// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Detector] class.
var DetectorClass = _DetectorClass{objc.GetClass("CIDetector")}

type _DetectorClass struct {
	objc.Class
}

// An interface definition for the [Detector] class.
type IDetector interface {
	objc.IObject
	FeaturesInImage(image IImage) []Feature
	FeaturesInImageOptions(image IImage, options map[string]objc.IObject) []Feature
}

// An image processor that identifies notable features (such as faces and barcodes) in a still image or video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetector?language=objc
type Detector struct {
	objc.Object
}

func DetectorFrom(ptr unsafe.Pointer) Detector {
	return Detector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DetectorClass) Alloc() Detector {
	rv := objc.Call[Detector](dc, objc.Sel("alloc"))
	return rv
}

func (dc _DetectorClass) New() Detector {
	rv := objc.Call[Detector](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetector() Detector {
	return DetectorClass.New()
}

func (d_ Detector) Init() Detector {
	rv := objc.Call[Detector](d_, objc.Sel("init"))
	return rv
}

// Creates and returns a configured detector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetector/1437884-detectoroftype?language=objc
func (dc _DetectorClass) DetectorOfTypeContextOptions(type_ string, context IContext, options map[string]objc.IObject) Detector {
	rv := objc.Call[Detector](dc, objc.Sel("detectorOfType:context:options:"), type_, context, options)
	return rv
}

// Creates and returns a configured detector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetector/1437884-detectoroftype?language=objc
func Detector_DetectorOfTypeContextOptions(type_ string, context IContext, options map[string]objc.IObject) Detector {
	return DetectorClass.DetectorOfTypeContextOptions(type_, context, options)
}

// Searches for features in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetector/1438049-featuresinimage?language=objc
func (d_ Detector) FeaturesInImage(image IImage) []Feature {
	rv := objc.Call[[]Feature](d_, objc.Sel("featuresInImage:"), image)
	return rv
}

// Searches for features in an image based on the specified image orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetector/1438189-featuresinimage?language=objc
func (d_ Detector) FeaturesInImageOptions(image IImage, options map[string]objc.IObject) []Feature {
	rv := objc.Call[[]Feature](d_, objc.Sel("featuresInImage:options:"), image, options)
	return rv
}
