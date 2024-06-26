// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [URLCache] class.
var URLCacheClass = _URLCacheClass{objc.GetClass("NSURLCache")}

type _URLCacheClass struct {
	objc.Class
}

// An interface definition for the [URLCache] class.
type IURLCache interface {
	objc.IObject
	RemoveCachedResponsesSinceDate(date IDate)
	StoreCachedResponseForDataTask(cachedResponse ICachedURLResponse, dataTask IURLSessionDataTask)
	GetCachedResponseForDataTaskCompletionHandler(dataTask IURLSessionDataTask, completionHandler func(cachedResponse CachedURLResponse))
	StoreCachedResponseForRequest(cachedResponse ICachedURLResponse, request IURLRequest)
	CachedResponseForRequest(request IURLRequest) CachedURLResponse
	RemoveAllCachedResponses()
	RemoveCachedResponseForRequest(request IURLRequest)
	RemoveCachedResponseForDataTask(dataTask IURLSessionDataTask)
	MemoryCapacity() uint
	SetMemoryCapacity(value uint)
	DiskCapacity() uint
	SetDiskCapacity(value uint)
	CurrentDiskUsage() uint
	CurrentMemoryUsage() uint
}

// An object that maps URL requests to cached response objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache?language=objc
type URLCache struct {
	objc.Object
}

func URLCacheFrom(ptr unsafe.Pointer) URLCache {
	return URLCache{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLCache) InitWithMemoryCapacityDiskCapacityDirectoryURL(memoryCapacity uint, diskCapacity uint, directoryURL IURL) URLCache {
	rv := objc.Call[URLCache](u_, objc.Sel("initWithMemoryCapacity:diskCapacity:directoryURL:"), memoryCapacity, diskCapacity, directoryURL)
	return rv
}

// Creates a URL cache object with the specified memory and disk capacities, in the specified directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/3240612-initwithmemorycapacity?language=objc
func NewURLCacheWithMemoryCapacityDiskCapacityDirectoryURL(memoryCapacity uint, diskCapacity uint, directoryURL IURL) URLCache {
	instance := URLCacheClass.Alloc().InitWithMemoryCapacityDiskCapacityDirectoryURL(memoryCapacity, diskCapacity, directoryURL)
	instance.Autorelease()
	return instance
}

func (uc _URLCacheClass) Alloc() URLCache {
	rv := objc.Call[URLCache](uc, objc.Sel("alloc"))
	return rv
}

func (uc _URLCacheClass) New() URLCache {
	rv := objc.Call[URLCache](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLCache() URLCache {
	return URLCacheClass.New()
}

func (u_ URLCache) Init() URLCache {
	rv := objc.Call[URLCache](u_, objc.Sel("init"))
	return rv
}

// Clears the given cache of any cached responses since the provided date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1415231-removecachedresponsessincedate?language=objc
func (u_ URLCache) RemoveCachedResponsesSinceDate(date IDate) {
	objc.Call[objc.Void](u_, objc.Sel("removeCachedResponsesSinceDate:"), date)
}

// Stores a cached URL response for a specified data task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1414434-storecachedresponse?language=objc
func (u_ URLCache) StoreCachedResponseForDataTask(cachedResponse ICachedURLResponse, dataTask IURLSessionDataTask) {
	objc.Call[objc.Void](u_, objc.Sel("storeCachedResponse:forDataTask:"), cachedResponse, dataTask)
}

// Gets the cached URL response for a data task, passing it to the provided completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1409184-getcachedresponsefordatatask?language=objc
func (u_ URLCache) GetCachedResponseForDataTaskCompletionHandler(dataTask IURLSessionDataTask, completionHandler func(cachedResponse CachedURLResponse)) {
	objc.Call[objc.Void](u_, objc.Sel("getCachedResponseForDataTask:completionHandler:"), dataTask, completionHandler)
}

// Stores a cached URL response for a specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1410340-storecachedresponse?language=objc
func (u_ URLCache) StoreCachedResponseForRequest(cachedResponse ICachedURLResponse, request IURLRequest) {
	objc.Call[objc.Void](u_, objc.Sel("storeCachedResponse:forRequest:"), cachedResponse, request)
}

// Returns the cached URL response in the cache for the specified URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1411817-cachedresponseforrequest?language=objc
func (u_ URLCache) CachedResponseForRequest(request IURLRequest) CachedURLResponse {
	rv := objc.Call[CachedURLResponse](u_, objc.Sel("cachedResponseForRequest:"), request)
	return rv
}

// Clears the receiver’s cache, removing all stored cached URL responses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1417802-removeallcachedresponses?language=objc
func (u_ URLCache) RemoveAllCachedResponses() {
	objc.Call[objc.Void](u_, objc.Sel("removeAllCachedResponses"))
}

// Removes the cached URL response for a specified URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1415377-removecachedresponseforrequest?language=objc
func (u_ URLCache) RemoveCachedResponseForRequest(request IURLRequest) {
	objc.Call[objc.Void](u_, objc.Sel("removeCachedResponseForRequest:"), request)
}

// Removes the cached URL response for a specified data task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1412258-removecachedresponsefordatatask?language=objc
func (u_ URLCache) RemoveCachedResponseForDataTask(dataTask IURLSessionDataTask) {
	objc.Call[objc.Void](u_, objc.Sel("removeCachedResponseForDataTask:"), dataTask)
}

// The capacity of the in-memory cache, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1409781-memorycapacity?language=objc
func (u_ URLCache) MemoryCapacity() uint {
	rv := objc.Call[uint](u_, objc.Sel("memoryCapacity"))
	return rv
}

// The capacity of the in-memory cache, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1409781-memorycapacity?language=objc
func (u_ URLCache) SetMemoryCapacity(value uint) {
	objc.Call[objc.Void](u_, objc.Sel("setMemoryCapacity:"), value)
}

// The capacity of the on-disk cache, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1413505-diskcapacity?language=objc
func (u_ URLCache) DiskCapacity() uint {
	rv := objc.Call[uint](u_, objc.Sel("diskCapacity"))
	return rv
}

// The capacity of the on-disk cache, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1413505-diskcapacity?language=objc
func (u_ URLCache) SetDiskCapacity(value uint) {
	objc.Call[objc.Void](u_, objc.Sel("setDiskCapacity:"), value)
}

// The current size of the on-disk cache, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1407771-currentdiskusage?language=objc
func (u_ URLCache) CurrentDiskUsage() uint {
	rv := objc.Call[uint](u_, objc.Sel("currentDiskUsage"))
	return rv
}

// The shared URL cache instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1413377-sharedurlcache?language=objc
func (uc _URLCacheClass) SharedURLCache() URLCache {
	rv := objc.Call[URLCache](uc, objc.Sel("sharedURLCache"))
	return rv
}

// The shared URL cache instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1413377-sharedurlcache?language=objc
func URLCache_SharedURLCache() URLCache {
	return URLCacheClass.SharedURLCache()
}

// The shared URL cache instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1413377-sharedurlcache?language=objc
func (uc _URLCacheClass) SetSharedURLCache(value IURLCache) {
	objc.Call[objc.Void](uc, objc.Sel("setSharedURLCache:"), value)
}

// The shared URL cache instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1413377-sharedurlcache?language=objc
func URLCache_SetSharedURLCache(value IURLCache) {
	URLCacheClass.SetSharedURLCache(value)
}

// The current size of the in-memory cache, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcache/1408199-currentmemoryusage?language=objc
func (u_ URLCache) CurrentMemoryUsage() uint {
	rv := objc.Call[uint](u_, objc.Sel("currentMemoryUsage"))
	return rv
}
