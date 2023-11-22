// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol that handles content key requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate?language=objc
type PContentKeySessionDelegate interface {
	// optional
	ContentKeySessionContentKeyRequestDidFailWithError(session ContentKeySession, keyRequest ContentKeyRequest, err foundation.Error)
	HasContentKeySessionContentKeyRequestDidFailWithError() bool

	// optional
	ContentKeySessionDidProvideContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest)
	HasContentKeySessionDidProvideContentKeyRequest() bool

	// optional
	ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session ContentKeySession, persistableContentKey []byte, keyIdentifier objc.Object)
	HasContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier() bool

	// optional
	ContentKeySessionDidProvideRenewingContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest)
	HasContentKeySessionDidProvideRenewingContentKeyRequest() bool

	// optional
	ContentKeySessionDidProvidePersistableContentKeyRequest(session ContentKeySession, keyRequest PersistableContentKeyRequest)
	HasContentKeySessionDidProvidePersistableContentKeyRequest() bool

	// optional
	ContentKeySessionDidGenerateExpiredSessionReport(session ContentKeySession)
	HasContentKeySessionDidGenerateExpiredSessionReport() bool

	// optional
	ContentKeySessionContentProtectionSessionIdentifierDidChange(session ContentKeySession)
	HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool

	// optional
	ContentKeySessionContentKeyRequestDidSucceed(session ContentKeySession, keyRequest ContentKeyRequest)
	HasContentKeySessionContentKeyRequestDidSucceed() bool

	// optional
	ContentKeySessionShouldRetryContentKeyRequestReason(session ContentKeySession, keyRequest ContentKeyRequest, retryReason ContentKeyRequestRetryReason) bool
	HasContentKeySessionShouldRetryContentKeyRequestReason() bool
}

// A delegate implementation builder for the [PContentKeySessionDelegate] protocol.
type ContentKeySessionDelegate struct {
	_ContentKeySessionContentKeyRequestDidFailWithError                     func(session ContentKeySession, keyRequest ContentKeyRequest, err foundation.Error)
	_ContentKeySessionDidProvideContentKeyRequest                           func(session ContentKeySession, keyRequest ContentKeyRequest)
	_ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier func(session ContentKeySession, persistableContentKey []byte, keyIdentifier objc.Object)
	_ContentKeySessionDidProvideRenewingContentKeyRequest                   func(session ContentKeySession, keyRequest ContentKeyRequest)
	_ContentKeySessionDidProvidePersistableContentKeyRequest                func(session ContentKeySession, keyRequest PersistableContentKeyRequest)
	_ContentKeySessionDidGenerateExpiredSessionReport                       func(session ContentKeySession)
	_ContentKeySessionContentProtectionSessionIdentifierDidChange           func(session ContentKeySession)
	_ContentKeySessionContentKeyRequestDidSucceed                           func(session ContentKeySession, keyRequest ContentKeyRequest)
	_ContentKeySessionShouldRetryContentKeyRequestReason                    func(session ContentKeySession, keyRequest ContentKeyRequest, retryReason ContentKeyRequestRetryReason) bool
}

func (di *ContentKeySessionDelegate) HasContentKeySessionContentKeyRequestDidFailWithError() bool {
	return di._ContentKeySessionContentKeyRequestDidFailWithError != nil
}

// Tells the receiver that the content key request failed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799201-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionContentKeyRequestDidFailWithError(f func(session ContentKeySession, keyRequest ContentKeyRequest, err foundation.Error)) {
	di._ContentKeySessionContentKeyRequestDidFailWithError = f
}

// Tells the receiver that the content key request failed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799201-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionContentKeyRequestDidFailWithError(session ContentKeySession, keyRequest ContentKeyRequest, err foundation.Error) {
	di._ContentKeySessionContentKeyRequestDidFailWithError(session, keyRequest, err)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionDidProvideContentKeyRequest() bool {
	return di._ContentKeySessionDidProvideContentKeyRequest != nil
}

// Provides the receiver with a new content key request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799204-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionDidProvideContentKeyRequest(f func(session ContentKeySession, keyRequest ContentKeyRequest)) {
	di._ContentKeySessionDidProvideContentKeyRequest = f
}

// Provides the receiver with a new content key request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799204-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionDidProvideContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest) {
	di._ContentKeySessionDidProvideContentKeyRequest(session, keyRequest)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier() bool {
	return di._ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier != nil
}

// Provides the receiver with an updated persistable content key for a specific key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2881821-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(f func(session ContentKeySession, persistableContentKey []byte, keyIdentifier objc.Object)) {
	di._ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier = f
}

// Provides the receiver with an updated persistable content key for a specific key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2881821-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session ContentKeySession, persistableContentKey []byte, keyIdentifier objc.Object) {
	di._ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session, persistableContentKey, keyIdentifier)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionDidProvideRenewingContentKeyRequest() bool {
	return di._ContentKeySessionDidProvideRenewingContentKeyRequest != nil
}

// Provides the receiver with a new content key request object for the renewal of an existing content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799168-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionDidProvideRenewingContentKeyRequest(f func(session ContentKeySession, keyRequest ContentKeyRequest)) {
	di._ContentKeySessionDidProvideRenewingContentKeyRequest = f
}

// Provides the receiver with a new content key request object for the renewal of an existing content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799168-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionDidProvideRenewingContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest) {
	di._ContentKeySessionDidProvideRenewingContentKeyRequest(session, keyRequest)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionDidProvidePersistableContentKeyRequest() bool {
	return di._ContentKeySessionDidProvidePersistableContentKeyRequest != nil
}

// Provides the receiver with a new content key request object to process a persistable content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799200-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionDidProvidePersistableContentKeyRequest(f func(session ContentKeySession, keyRequest PersistableContentKeyRequest)) {
	di._ContentKeySessionDidProvidePersistableContentKeyRequest = f
}

// Provides the receiver with a new content key request object to process a persistable content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799200-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionDidProvidePersistableContentKeyRequest(session ContentKeySession, keyRequest PersistableContentKeyRequest) {
	di._ContentKeySessionDidProvidePersistableContentKeyRequest(session, keyRequest)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionDidGenerateExpiredSessionReport() bool {
	return di._ContentKeySessionDidGenerateExpiredSessionReport != nil
}

// Notifies the sender that an expired session report has been generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966513-contentkeysessiondidgenerateexpi?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionDidGenerateExpiredSessionReport(f func(session ContentKeySession)) {
	di._ContentKeySessionDidGenerateExpiredSessionReport = f
}

// Notifies the sender that an expired session report has been generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966513-contentkeysessiondidgenerateexpi?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionDidGenerateExpiredSessionReport(session ContentKeySession) {
	di._ContentKeySessionDidGenerateExpiredSessionReport(session)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool {
	return di._ContentKeySessionContentProtectionSessionIdentifierDidChange != nil
}

// Tells the receiver the content protection session identifier changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799160-contentkeysessioncontentprotecti?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionContentProtectionSessionIdentifierDidChange(f func(session ContentKeySession)) {
	di._ContentKeySessionContentProtectionSessionIdentifierDidChange = f
}

// Tells the receiver the content protection session identifier changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799160-contentkeysessioncontentprotecti?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionContentProtectionSessionIdentifierDidChange(session ContentKeySession) {
	di._ContentKeySessionContentProtectionSessionIdentifierDidChange(session)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionContentKeyRequestDidSucceed() bool {
	return di._ContentKeySessionContentKeyRequestDidSucceed != nil
}

// Tells the content key session that the response to a content key requeset was successfully processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966512-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionContentKeyRequestDidSucceed(f func(session ContentKeySession, keyRequest ContentKeyRequest)) {
	di._ContentKeySessionContentKeyRequestDidSucceed = f
}

// Tells the content key session that the response to a content key requeset was successfully processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966512-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionContentKeyRequestDidSucceed(session ContentKeySession, keyRequest ContentKeyRequest) {
	di._ContentKeySessionContentKeyRequestDidSucceed(session, keyRequest)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionShouldRetryContentKeyRequestReason() bool {
	return di._ContentKeySessionShouldRetryContentKeyRequestReason != nil
}

// Provides the receiver with a content key request object to retry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799210-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionShouldRetryContentKeyRequestReason(f func(session ContentKeySession, keyRequest ContentKeyRequest, retryReason ContentKeyRequestRetryReason) bool) {
	di._ContentKeySessionShouldRetryContentKeyRequestReason = f
}

// Provides the receiver with a content key request object to retry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799210-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionShouldRetryContentKeyRequestReason(session ContentKeySession, keyRequest ContentKeyRequest, retryReason ContentKeyRequestRetryReason) bool {
	return di._ContentKeySessionShouldRetryContentKeyRequestReason(session, keyRequest, retryReason)
}

// ensure impl type implements protocol interface
var _ PContentKeySessionDelegate = (*ContentKeySessionDelegateObject)(nil)

// A concrete type for the [PContentKeySessionDelegate] protocol.
type ContentKeySessionDelegateObject struct {
	objc.Object
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionContentKeyRequestDidFailWithError() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:contentKeyRequest:didFailWithError:"))
}

// Tells the receiver that the content key request failed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799201-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionContentKeyRequestDidFailWithError(session ContentKeySession, keyRequest ContentKeyRequest, err foundation.Error) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:contentKeyRequest:didFailWithError:"), objc.Ptr(session), objc.Ptr(keyRequest), objc.Ptr(err))
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionDidProvideContentKeyRequest() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:didProvideContentKeyRequest:"))
}

// Provides the receiver with a new content key request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799204-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionDidProvideContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:didProvideContentKeyRequest:"), objc.Ptr(session), objc.Ptr(keyRequest))
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:didUpdatePersistableContentKey:forContentKeyIdentifier:"))
}

// Provides the receiver with an updated persistable content key for a specific key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2881821-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session ContentKeySession, persistableContentKey []byte, keyIdentifier objc.Object) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:didUpdatePersistableContentKey:forContentKeyIdentifier:"), objc.Ptr(session), persistableContentKey, keyIdentifier)
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionDidProvideRenewingContentKeyRequest() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:didProvideRenewingContentKeyRequest:"))
}

// Provides the receiver with a new content key request object for the renewal of an existing content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799168-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionDidProvideRenewingContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:didProvideRenewingContentKeyRequest:"), objc.Ptr(session), objc.Ptr(keyRequest))
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionDidProvidePersistableContentKeyRequest() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:didProvidePersistableContentKeyRequest:"))
}

// Provides the receiver with a new content key request object to process a persistable content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799200-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionDidProvidePersistableContentKeyRequest(session ContentKeySession, keyRequest PersistableContentKeyRequest) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:didProvidePersistableContentKeyRequest:"), objc.Ptr(session), objc.Ptr(keyRequest))
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionDidGenerateExpiredSessionReport() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySessionDidGenerateExpiredSessionReport:"))
}

// Notifies the sender that an expired session report has been generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966513-contentkeysessiondidgenerateexpi?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionDidGenerateExpiredSessionReport(session ContentKeySession) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySessionDidGenerateExpiredSessionReport:"), objc.Ptr(session))
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySessionContentProtectionSessionIdentifierDidChange:"))
}

// Tells the receiver the content protection session identifier changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799160-contentkeysessioncontentprotecti?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionContentProtectionSessionIdentifierDidChange(session ContentKeySession) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySessionContentProtectionSessionIdentifierDidChange:"), objc.Ptr(session))
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionContentKeyRequestDidSucceed() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:contentKeyRequestDidSucceed:"))
}

// Tells the content key session that the response to a content key requeset was successfully processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966512-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionContentKeyRequestDidSucceed(session ContentKeySession, keyRequest ContentKeyRequest) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:contentKeyRequestDidSucceed:"), objc.Ptr(session), objc.Ptr(keyRequest))
}

func (c_ ContentKeySessionDelegateObject) HasContentKeySessionShouldRetryContentKeyRequestReason() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:shouldRetryContentKeyRequest:reason:"))
}

// Provides the receiver with a content key request object to retry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799210-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateObject) ContentKeySessionShouldRetryContentKeyRequestReason(session ContentKeySession, keyRequest ContentKeyRequest, retryReason ContentKeyRequestRetryReason) bool {
	rv := objc.Call[bool](c_, objc.Sel("contentKeySession:shouldRetryContentKeyRequest:reason:"), objc.Ptr(session), objc.Ptr(keyRequest), retryReason)
	return rv
}