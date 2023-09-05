// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "Metal/Metal.h"
// MTLRegion RegionMake2D(uint x, uint y, uint width, uint height);
// void RemoveDeviceObserver(void * observer);
// MTLRegion RegionMake1D(uint x, uint width);
// MTLCoordinate2D Coordinate2DMake(float x, float y);
// MTLRegion RegionMake3D(uint x, uint y, uint z, uint width, uint height, uint depth);
// MTLSamplePosition SamplePositionMake(float x, float y);
// MTLClearColor ClearColorMake(double red, double green, double blue, double alpha);
// MTLTextureSwizzleChannels TextureSwizzleChannelsMake(MTLTextureSwizzle r, MTLTextureSwizzle g, MTLTextureSwizzle b, MTLTextureSwizzle a);
// MTLOrigin OriginMake(uint x, uint y, uint z);
// void * CreateSystemDefaultDevice();
import "C"
import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// Creates a 3D representation of a 2D region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/1515675-mtlregionmake2d?language=objc
func RegionMake2D(x uint, y uint, width uint, height uint) Region {
	rv := C.RegionMake2D(
		// *typing.PrimitiveType
		C.uint(x),
		// *typing.PrimitiveType
		C.uint(y),
		// *typing.PrimitiveType
		C.uint(width),
		// *typing.PrimitiveType
		C.uint(height),
	)
	// *typing.StructType
	return *(*Region)(unsafe.Pointer(&rv))
}

// Removes a registered observer of device notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/2869724-mtlremovedeviceobserver?language=objc
func RemoveDeviceObserver(observer objc.Object) {
	C.RemoveDeviceObserver(
		// *typing.IDType
		observer.Ptr(),
	)
}

// Creates a 3D representation of a 1D region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/1516011-mtlregionmake1d?language=objc
func RegionMake1D(x uint, width uint) Region {
	rv := C.RegionMake1D(
		// *typing.PrimitiveType
		C.uint(x),
		// *typing.PrimitiveType
		C.uint(width),
	)
	// *typing.StructType
	return *(*Region)(unsafe.Pointer(&rv))
}

// Returns a new 2D point with the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/3088856-mtlcoordinate2dmake?language=objc
func Coordinate2DMake(x float64, y float64) Coordinate2D {
	rv := C.Coordinate2DMake(
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
	// *typing.AliasType
	return *(*Coordinate2D)(unsafe.Pointer(&rv))
}

// Creates a 3D region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/1515765-mtlregionmake3d?language=objc
func RegionMake3D(x uint, y uint, z uint, width uint, height uint, depth uint) Region {
	rv := C.RegionMake3D(
		// *typing.PrimitiveType
		C.uint(x),
		// *typing.PrimitiveType
		C.uint(y),
		// *typing.PrimitiveType
		C.uint(z),
		// *typing.PrimitiveType
		C.uint(width),
		// *typing.PrimitiveType
		C.uint(height),
		// *typing.PrimitiveType
		C.uint(depth),
	)
	// *typing.StructType
	return *(*Region)(unsafe.Pointer(&rv))
}

// Returns a new sample position on a subpixel grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/2866347-mtlsamplepositionmake?language=objc
func SamplePositionMake(x float64, y float64) SamplePosition {
	rv := C.SamplePositionMake(
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
	// *typing.StructType
	return *(*SamplePosition)(unsafe.Pointer(&rv))
}

// Returns a color value used to clear a color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/1437971-mtlclearcolormake?language=objc
func ClearColorMake(red float64, green float64, blue float64, alpha float64) ClearColor {
	rv := C.ClearColorMake(
		// *typing.PrimitiveType
		C.double(red),
		// *typing.PrimitiveType
		C.double(green),
		// *typing.PrimitiveType
		C.double(blue),
		// *typing.PrimitiveType
		C.double(alpha),
	)
	// *typing.StructType
	return *(*ClearColor)(unsafe.Pointer(&rv))
}

// Creates a new swizzle pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/3114319-mtltextureswizzlechannelsmake?language=objc
func TextureSwizzleChannelsMake(r TextureSwizzle, g TextureSwizzle, b TextureSwizzle, a TextureSwizzle) TextureSwizzleChannels {
	rv := C.TextureSwizzleChannelsMake(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.MTLTextureSwizzle)(r),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.MTLTextureSwizzle)(g),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.MTLTextureSwizzle)(b),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.MTLTextureSwizzle)(a),
	)
	// *typing.StructType
	return *(*TextureSwizzleChannels)(unsafe.Pointer(&rv))
}

// Returns a new origin with the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/1516140-mtloriginmake?language=objc
func OriginMake(x uint, y uint, z uint) Origin {
	rv := C.OriginMake(
		// *typing.PrimitiveType
		C.uint(x),
		// *typing.PrimitiveType
		C.uint(y),
		// *typing.PrimitiveType
		C.uint(z),
	)
	// *typing.StructType
	return *(*Origin)(unsafe.Pointer(&rv))
}

// Returns the device instance Metal selects as the default. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/1433401-mtlcreatesystemdefaultdevice?language=objc
func CreateSystemDefaultDevice() DeviceWrapper {
	rv := C.CreateSystemDefaultDevice()
	// *typing.ProtocolType
	return DeviceWrapper{objc.ObjectFrom(rv)}
}
