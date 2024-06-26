// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/darwinkit/objc"
)

// The interface a file manager's delegate uses to intervene during operations or if an error occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate?language=objc
type PFileManagerDelegate interface {
	// optional
	FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool
	HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath() bool

	// optional
	FileManagerShouldMoveItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool
	HasFileManagerShouldMoveItemAtPathToPath() bool

	// optional
	FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool
	HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL() bool

	// optional
	FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool
	HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL() bool

	// optional
	FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool
	HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath() bool

	// optional
	FileManagerShouldMoveItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool
	HasFileManagerShouldMoveItemAtURLToURL() bool

	// optional
	FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool
	HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL() bool

	// optional
	FileManagerShouldLinkItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool
	HasFileManagerShouldLinkItemAtPathToPath() bool

	// optional
	FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager FileManager, error Error, path string) bool
	HasFileManagerShouldProceedAfterErrorRemovingItemAtPath() bool

	// optional
	FileManagerShouldLinkItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool
	HasFileManagerShouldLinkItemAtURLToURL() bool

	// optional
	FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager FileManager, error Error, URL URL) bool
	HasFileManagerShouldProceedAfterErrorRemovingItemAtURL() bool

	// optional
	FileManagerShouldRemoveItemAtPath(fileManager FileManager, path string) bool
	HasFileManagerShouldRemoveItemAtPath() bool

	// optional
	FileManagerShouldCopyItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool
	HasFileManagerShouldCopyItemAtPathToPath() bool

	// optional
	FileManagerShouldCopyItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool
	HasFileManagerShouldCopyItemAtURLToURL() bool

	// optional
	FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool
	HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath() bool

	// optional
	FileManagerShouldRemoveItemAtURL(fileManager FileManager, URL URL) bool
	HasFileManagerShouldRemoveItemAtURL() bool
}

// A delegate implementation builder for the [PFileManagerDelegate] protocol.
type FileManagerDelegate struct {
	_FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath func(fileManager FileManager, error Error, srcPath string, dstPath string) bool
	_FileManagerShouldMoveItemAtPathToPath                     func(fileManager FileManager, srcPath string, dstPath string) bool
	_FileManagerShouldProceedAfterErrorMovingItemAtURLToURL    func(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool
	_FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL   func(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool
	_FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath func(fileManager FileManager, error Error, srcPath string, dstPath string) bool
	_FileManagerShouldMoveItemAtURLToURL                       func(fileManager FileManager, srcURL URL, dstURL URL) bool
	_FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL   func(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool
	_FileManagerShouldLinkItemAtPathToPath                     func(fileManager FileManager, srcPath string, dstPath string) bool
	_FileManagerShouldProceedAfterErrorRemovingItemAtPath      func(fileManager FileManager, error Error, path string) bool
	_FileManagerShouldLinkItemAtURLToURL                       func(fileManager FileManager, srcURL URL, dstURL URL) bool
	_FileManagerShouldProceedAfterErrorRemovingItemAtURL       func(fileManager FileManager, error Error, URL URL) bool
	_FileManagerShouldRemoveItemAtPath                         func(fileManager FileManager, path string) bool
	_FileManagerShouldCopyItemAtPathToPath                     func(fileManager FileManager, srcPath string, dstPath string) bool
	_FileManagerShouldCopyItemAtURLToURL                       func(fileManager FileManager, srcURL URL, dstURL URL) bool
	_FileManagerShouldProceedAfterErrorMovingItemAtPathToPath  func(fileManager FileManager, error Error, srcPath string, dstPath string) bool
	_FileManagerShouldRemoveItemAtURL                          func(fileManager FileManager, URL URL) bool
}

func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath() bool {
	return di._FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath != nil
}

// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1410189-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(f func(fileManager FileManager, error Error, srcPath string, dstPath string) bool) {
	di._FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath = f
}

// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1410189-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool {
	return di._FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager, error, srcPath, dstPath)
}
func (di *FileManagerDelegate) HasFileManagerShouldMoveItemAtPathToPath() bool {
	return di._FileManagerShouldMoveItemAtPathToPath != nil
}

// Asks the delegate if the file manager should move the specified item to the new path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1407734-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldMoveItemAtPathToPath(f func(fileManager FileManager, srcPath string, dstPath string) bool) {
	di._FileManagerShouldMoveItemAtPathToPath = f
}

// Asks the delegate if the file manager should move the specified item to the new path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1407734-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldMoveItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool {
	return di._FileManagerShouldMoveItemAtPathToPath(fileManager, srcPath, dstPath)
}
func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL() bool {
	return di._FileManagerShouldProceedAfterErrorMovingItemAtURLToURL != nil
}

// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411289-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorMovingItemAtURLToURL(f func(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool) {
	di._FileManagerShouldProceedAfterErrorMovingItemAtURLToURL = f
}

// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411289-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool {
	return di._FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager, error, srcURL, dstURL)
}
func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL() bool {
	return di._FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL != nil
}

// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1410788-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(f func(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool) {
	di._FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL = f
}

// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1410788-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool {
	return di._FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager, error, srcURL, dstURL)
}
func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath() bool {
	return di._FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath != nil
}

// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1415627-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(f func(fileManager FileManager, error Error, srcPath string, dstPath string) bool) {
	di._FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath = f
}

// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1415627-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool {
	return di._FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager, error, srcPath, dstPath)
}
func (di *FileManagerDelegate) HasFileManagerShouldMoveItemAtURLToURL() bool {
	return di._FileManagerShouldMoveItemAtURLToURL != nil
}

// Asks the delegate if the file manager should move the specified item to the new URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411878-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldMoveItemAtURLToURL(f func(fileManager FileManager, srcURL URL, dstURL URL) bool) {
	di._FileManagerShouldMoveItemAtURLToURL = f
}

// Asks the delegate if the file manager should move the specified item to the new URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411878-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldMoveItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool {
	return di._FileManagerShouldMoveItemAtURLToURL(fileManager, srcURL, dstURL)
}
func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL() bool {
	return di._FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL != nil
}

// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1408003-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(f func(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool) {
	di._FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL = f
}

// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1408003-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool {
	return di._FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager, error, srcURL, dstURL)
}
func (di *FileManagerDelegate) HasFileManagerShouldLinkItemAtPathToPath() bool {
	return di._FileManagerShouldLinkItemAtPathToPath != nil
}

// Asks the delegate if a hard link should be created between the items at the two paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1414699-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldLinkItemAtPathToPath(f func(fileManager FileManager, srcPath string, dstPath string) bool) {
	di._FileManagerShouldLinkItemAtPathToPath = f
}

// Asks the delegate if a hard link should be created between the items at the two paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1414699-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldLinkItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool {
	return di._FileManagerShouldLinkItemAtPathToPath(fileManager, srcPath, dstPath)
}
func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorRemovingItemAtPath() bool {
	return di._FileManagerShouldProceedAfterErrorRemovingItemAtPath != nil
}

// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1409791-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorRemovingItemAtPath(f func(fileManager FileManager, error Error, path string) bool) {
	di._FileManagerShouldProceedAfterErrorRemovingItemAtPath = f
}

// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1409791-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager FileManager, error Error, path string) bool {
	return di._FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager, error, path)
}
func (di *FileManagerDelegate) HasFileManagerShouldLinkItemAtURLToURL() bool {
	return di._FileManagerShouldLinkItemAtURLToURL != nil
}

// Asks the delegate if a hard link should be created between the items at the two URLs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1417589-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldLinkItemAtURLToURL(f func(fileManager FileManager, srcURL URL, dstURL URL) bool) {
	di._FileManagerShouldLinkItemAtURLToURL = f
}

// Asks the delegate if a hard link should be created between the items at the two URLs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1417589-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldLinkItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool {
	return di._FileManagerShouldLinkItemAtURLToURL(fileManager, srcURL, dstURL)
}
func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorRemovingItemAtURL() bool {
	return di._FileManagerShouldProceedAfterErrorRemovingItemAtURL != nil
}

// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1408660-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorRemovingItemAtURL(f func(fileManager FileManager, error Error, URL URL) bool) {
	di._FileManagerShouldProceedAfterErrorRemovingItemAtURL = f
}

// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1408660-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager FileManager, error Error, URL URL) bool {
	return di._FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager, error, URL)
}
func (di *FileManagerDelegate) HasFileManagerShouldRemoveItemAtPath() bool {
	return di._FileManagerShouldRemoveItemAtPath != nil
}

// Asks the delegate whether the item at the specified path should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1412994-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldRemoveItemAtPath(f func(fileManager FileManager, path string) bool) {
	di._FileManagerShouldRemoveItemAtPath = f
}

// Asks the delegate whether the item at the specified path should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1412994-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldRemoveItemAtPath(fileManager FileManager, path string) bool {
	return di._FileManagerShouldRemoveItemAtPath(fileManager, path)
}
func (di *FileManagerDelegate) HasFileManagerShouldCopyItemAtPathToPath() bool {
	return di._FileManagerShouldCopyItemAtPathToPath != nil
}

// Asks the delegate if the file manager should copy the specified item to the new path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1414922-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldCopyItemAtPathToPath(f func(fileManager FileManager, srcPath string, dstPath string) bool) {
	di._FileManagerShouldCopyItemAtPathToPath = f
}

// Asks the delegate if the file manager should copy the specified item to the new path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1414922-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldCopyItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool {
	return di._FileManagerShouldCopyItemAtPathToPath(fileManager, srcPath, dstPath)
}
func (di *FileManagerDelegate) HasFileManagerShouldCopyItemAtURLToURL() bool {
	return di._FileManagerShouldCopyItemAtURLToURL != nil
}

// Asks the delegate if the file manager should copy the specified item to the new URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1417936-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldCopyItemAtURLToURL(f func(fileManager FileManager, srcURL URL, dstURL URL) bool) {
	di._FileManagerShouldCopyItemAtURLToURL = f
}

// Asks the delegate if the file manager should copy the specified item to the new URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1417936-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldCopyItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool {
	return di._FileManagerShouldCopyItemAtURLToURL(fileManager, srcURL, dstURL)
}
func (di *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath() bool {
	return di._FileManagerShouldProceedAfterErrorMovingItemAtPathToPath != nil
}

// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1412865-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorMovingItemAtPathToPath(f func(fileManager FileManager, error Error, srcPath string, dstPath string) bool) {
	di._FileManagerShouldProceedAfterErrorMovingItemAtPathToPath = f
}

// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1412865-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool {
	return di._FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager, error, srcPath, dstPath)
}
func (di *FileManagerDelegate) HasFileManagerShouldRemoveItemAtURL() bool {
	return di._FileManagerShouldRemoveItemAtURL != nil
}

// Asks the delegate whether the item at the specified URL should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411918-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldRemoveItemAtURL(f func(fileManager FileManager, URL URL) bool) {
	di._FileManagerShouldRemoveItemAtURL = f
}

// Asks the delegate whether the item at the specified URL should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411918-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldRemoveItemAtURL(fileManager FileManager, URL URL) bool {
	return di._FileManagerShouldRemoveItemAtURL(fileManager, URL)
}

// ensure impl type implements protocol interface
var _ PFileManagerDelegate = (*FileManagerDelegateObject)(nil)

// A concrete type for the [PFileManagerDelegate] protocol.
type FileManagerDelegateObject struct {
	objc.Object
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtPath:toPath:"))
}

// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1410189-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtPath:toPath:"), fileManager, error, srcPath, dstPath)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldMoveItemAtPathToPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldMoveItemAtPath:toPath:"))
}

// Asks the delegate if the file manager should move the specified item to the new path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1407734-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldMoveItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldMoveItemAtPath:toPath:"), fileManager, srcPath, dstPath)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:movingItemAtURL:toURL:"))
}

// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411289-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:movingItemAtURL:toURL:"), fileManager, error, srcURL, dstURL)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtURL:toURL:"))
}

// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1410788-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtURL:toURL:"), fileManager, error, srcURL, dstURL)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtPath:toPath:"))
}

// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1415627-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtPath:toPath:"), fileManager, error, srcPath, dstPath)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldMoveItemAtURLToURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldMoveItemAtURL:toURL:"))
}

// Asks the delegate if the file manager should move the specified item to the new URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411878-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldMoveItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldMoveItemAtURL:toURL:"), fileManager, srcURL, dstURL)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtURL:toURL:"))
}

// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1408003-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager FileManager, error Error, srcURL URL, dstURL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtURL:toURL:"), fileManager, error, srcURL, dstURL)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldLinkItemAtPathToPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldLinkItemAtPath:toPath:"))
}

// Asks the delegate if a hard link should be created between the items at the two paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1414699-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldLinkItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldLinkItemAtPath:toPath:"), fileManager, srcPath, dstPath)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorRemovingItemAtPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:removingItemAtPath:"))
}

// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1409791-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager FileManager, error Error, path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:removingItemAtPath:"), fileManager, error, path)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldLinkItemAtURLToURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldLinkItemAtURL:toURL:"))
}

// Asks the delegate if a hard link should be created between the items at the two URLs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1417589-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldLinkItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldLinkItemAtURL:toURL:"), fileManager, srcURL, dstURL)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorRemovingItemAtURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:removingItemAtURL:"))
}

// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1408660-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager FileManager, error Error, URL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:removingItemAtURL:"), fileManager, error, URL)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldRemoveItemAtPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldRemoveItemAtPath:"))
}

// Asks the delegate whether the item at the specified path should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1412994-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldRemoveItemAtPath(fileManager FileManager, path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldRemoveItemAtPath:"), fileManager, path)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldCopyItemAtPathToPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldCopyItemAtPath:toPath:"))
}

// Asks the delegate if the file manager should copy the specified item to the new path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1414922-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldCopyItemAtPathToPath(fileManager FileManager, srcPath string, dstPath string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldCopyItemAtPath:toPath:"), fileManager, srcPath, dstPath)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldCopyItemAtURLToURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldCopyItemAtURL:toURL:"))
}

// Asks the delegate if the file manager should copy the specified item to the new URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1417936-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldCopyItemAtURLToURL(fileManager FileManager, srcURL URL, dstURL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldCopyItemAtURL:toURL:"), fileManager, srcURL, dstURL)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldProceedAfterError:movingItemAtPath:toPath:"))
}

// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1412865-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager FileManager, error Error, srcPath string, dstPath string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldProceedAfterError:movingItemAtPath:toPath:"), fileManager, error, srcPath, dstPath)
	return rv
}

func (f_ FileManagerDelegateObject) HasFileManagerShouldRemoveItemAtURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldRemoveItemAtURL:"))
}

// Asks the delegate whether the item at the specified URL should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411918-filemanager?language=objc
func (f_ FileManagerDelegateObject) FileManagerShouldRemoveItemAtURL(fileManager FileManager, URL URL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldRemoveItemAtURL:"), fileManager, URL)
	return rv
}
