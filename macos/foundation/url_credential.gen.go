// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [URLCredential] class.
var URLCredentialClass = _URLCredentialClass{objc.GetClass("NSURLCredential")}

type _URLCredentialClass struct {
	objc.Class
}

// An interface definition for the [URLCredential] class.
type IURLCredential interface {
	objc.IObject
	Certificates() []objc.Object
	Persistence() URLCredentialPersistence
	Password() string
	User() string
	HasPassword() bool
}

// An authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential?language=objc
type URLCredential struct {
	objc.Object
}

func URLCredentialFrom(ptr unsafe.Pointer) URLCredential {
	return URLCredential{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLCredential) InitWithUserPasswordPersistence(user string, password string, persistence URLCredentialPersistence) URLCredential {
	rv := objc.Call[URLCredential](u_, objc.Sel("initWithUser:password:persistence:"), user, password, persistence)
	return rv
}

// Creates a URL credential instance initialized with a given user name and password, using a given persistence setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1417977-initwithuser?language=objc
func NewURLCredentialWithUserPasswordPersistence(user string, password string, persistence URLCredentialPersistence) URLCredential {
	instance := URLCredentialClass.Alloc().InitWithUserPasswordPersistence(user, password, persistence)
	instance.Autorelease()
	return instance
}

func (uc _URLCredentialClass) Alloc() URLCredential {
	rv := objc.Call[URLCredential](uc, objc.Sel("alloc"))
	return rv
}

func (uc _URLCredentialClass) New() URLCredential {
	rv := objc.Call[URLCredential](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLCredential() URLCredential {
	return URLCredentialClass.New()
}

func (u_ URLCredential) Init() URLCredential {
	rv := objc.Call[URLCredential](u_, objc.Sel("init"))
	return rv
}

// Creates a URL credential instance for internet password authentication with a given user name and password, using a given persistence setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1428174-credentialwithuser?language=objc
func (uc _URLCredentialClass) CredentialWithUserPasswordPersistence(user string, password string, persistence URLCredentialPersistence) URLCredential {
	rv := objc.Call[URLCredential](uc, objc.Sel("credentialWithUser:password:persistence:"), user, password, persistence)
	return rv
}

// Creates a URL credential instance for internet password authentication with a given user name and password, using a given persistence setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1428174-credentialwithuser?language=objc
func URLCredential_CredentialWithUserPasswordPersistence(user string, password string, persistence URLCredentialPersistence) URLCredential {
	return URLCredentialClass.CredentialWithUserPasswordPersistence(user, password, persistence)
}

// The intermediate certificates of the credential, if it is a client certificate credential. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1412369-certificates?language=objc
func (u_ URLCredential) Certificates() []objc.Object {
	rv := objc.Call[[]objc.Object](u_, objc.Sel("certificates"))
	return rv
}

// The credential’s persistence setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1416977-persistence?language=objc
func (u_ URLCredential) Persistence() URLCredentialPersistence {
	rv := objc.Call[URLCredentialPersistence](u_, objc.Sel("persistence"))
	return rv
}

// The credential’s password. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1417913-password?language=objc
func (u_ URLCredential) Password() string {
	rv := objc.Call[string](u_, objc.Sel("password"))
	return rv
}

// The credential’s user name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1408654-user?language=objc
func (u_ URLCredential) User() string {
	rv := objc.Call[string](u_, objc.Sel("user"))
	return rv
}

// A Boolean value that indicates whether the credential has a password. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredential/1418072-haspassword?language=objc
func (u_ URLCredential) HasPassword() bool {
	rv := objc.Call[bool](u_, objc.Sel("hasPassword"))
	return rv
}
