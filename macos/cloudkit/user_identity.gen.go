// Code generated by DarwinKit. DO NOT EDIT.

package cloudkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [UserIdentity] class.
var UserIdentityClass = _UserIdentityClass{objc.GetClass("CKUserIdentity")}

type _UserIdentityClass struct {
	objc.Class
}

// An interface definition for the [UserIdentity] class.
type IUserIdentity interface {
	objc.IObject
	HasiCloudAccount() bool
	LookupInfo() UserIdentityLookupInfo
	NameComponents() foundation.PersonNameComponents
	ContactIdentifiers() []string
	UserRecordID() RecordID
}

// The identity of a user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity?language=objc
type UserIdentity struct {
	objc.Object
}

func UserIdentityFrom(ptr unsafe.Pointer) UserIdentity {
	return UserIdentity{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UserIdentityClass) Alloc() UserIdentity {
	rv := objc.Call[UserIdentity](uc, objc.Sel("alloc"))
	return rv
}

func (uc _UserIdentityClass) New() UserIdentity {
	rv := objc.Call[UserIdentity](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserIdentity() UserIdentity {
	return UserIdentityClass.New()
}

func (u_ UserIdentity) Init() UserIdentity {
	rv := objc.Call[UserIdentity](u_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the user has an iCloud account. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/1640513-hasicloudaccount?language=objc
func (u_ UserIdentity) HasiCloudAccount() bool {
	rv := objc.Call[bool](u_, objc.Sel("hasiCloudAccount"))
	return rv
}

// The lookup info for retrieving the user identity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/1640371-lookupinfo?language=objc
func (u_ UserIdentity) LookupInfo() UserIdentityLookupInfo {
	rv := objc.Call[UserIdentityLookupInfo](u_, objc.Sel("lookupInfo"))
	return rv
}

// The user’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/1640458-namecomponents?language=objc
func (u_ UserIdentity) NameComponents() foundation.PersonNameComponents {
	rv := objc.Call[foundation.PersonNameComponents](u_, objc.Sel("nameComponents"))
	return rv
}

// Identifiers that match contacts in the local Contacts database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/2866227-contactidentifiers?language=objc
func (u_ UserIdentity) ContactIdentifiers() []string {
	rv := objc.Call[[]string](u_, objc.Sel("contactIdentifiers"))
	return rv
}

// The user record ID for the corresponding user record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/1640504-userrecordid?language=objc
func (u_ UserIdentity) UserRecordID() RecordID {
	rv := objc.Call[RecordID](u_, objc.Sel("userRecordID"))
	return rv
}
