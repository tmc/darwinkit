// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SaveRequest] class.
var SaveRequestClass = _SaveRequestClass{objc.GetClass("CNSaveRequest")}

type _SaveRequestClass struct {
	objc.Class
}

// An interface definition for the [SaveRequest] class.
type ISaveRequest interface {
	objc.IObject
	DeleteGroup(group IMutableGroup)
	UpdateGroup(group IMutableGroup)
	AddMemberToGroup(contact IContact, group IGroup)
	DeleteContact(contact IMutableContact)
	RemoveMemberFromGroup(contact IContact, group IGroup)
	UpdateContact(contact IMutableContact)
	AddGroupToContainerWithIdentifier(group IMutableGroup, identifier string)
	AddContactToContainerWithIdentifier(contact IMutableContact, identifier string)
	AddSubgroupToGroup(subgroup IGroup, group IGroup)
	RemoveSubgroupFromGroup(subgroup IGroup, group IGroup)
	TransactionAuthor() string
	SetTransactionAuthor(value string)
}

// An object that collects the changes you want to save to the user's contacts database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest?language=objc
type SaveRequest struct {
	objc.Object
}

func SaveRequestFrom(ptr unsafe.Pointer) SaveRequest {
	return SaveRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SaveRequestClass) Alloc() SaveRequest {
	rv := objc.Call[SaveRequest](sc, objc.Sel("alloc"))
	return rv
}

func SaveRequest_Alloc() SaveRequest {
	return SaveRequestClass.Alloc()
}

func (sc _SaveRequestClass) New() SaveRequest {
	rv := objc.Call[SaveRequest](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSaveRequest() SaveRequest {
	return SaveRequestClass.New()
}

func (s_ SaveRequest) Init() SaveRequest {
	rv := objc.Call[SaveRequest](s_, objc.Sel("init"))
	return rv
}

// Deletes a group from the contact store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1402859-deletegroup?language=objc
func (s_ SaveRequest) DeleteGroup(group IMutableGroup) {
	objc.Call[objc.Void](s_, objc.Sel("deleteGroup:"), objc.Ptr(group))
}

// Updates an existing group in the contact store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1403387-updategroup?language=objc
func (s_ SaveRequest) UpdateGroup(group IMutableGroup) {
	objc.Call[objc.Void](s_, objc.Sel("updateGroup:"), objc.Ptr(group))
}

// Adds a contact as a member of a group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1403180-addmember?language=objc
func (s_ SaveRequest) AddMemberToGroup(contact IContact, group IGroup) {
	objc.Call[objc.Void](s_, objc.Sel("addMember:toGroup:"), objc.Ptr(contact), objc.Ptr(group))
}

// Deletes a contact from the contact store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1402970-deletecontact?language=objc
func (s_ SaveRequest) DeleteContact(contact IMutableContact) {
	objc.Call[objc.Void](s_, objc.Sel("deleteContact:"), objc.Ptr(contact))
}

// Removes a contact as a member of a group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1403373-removemember?language=objc
func (s_ SaveRequest) RemoveMemberFromGroup(contact IContact, group IGroup) {
	objc.Call[objc.Void](s_, objc.Sel("removeMember:fromGroup:"), objc.Ptr(contact), objc.Ptr(group))
}

// Updates an existing contact in the contact store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1403074-updatecontact?language=objc
func (s_ SaveRequest) UpdateContact(contact IMutableContact) {
	objc.Call[objc.Void](s_, objc.Sel("updateContact:"), objc.Ptr(contact))
}

// Adds a group to the contact store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1402821-addgroup?language=objc
func (s_ SaveRequest) AddGroupToContainerWithIdentifier(group IMutableGroup, identifier string) {
	objc.Call[objc.Void](s_, objc.Sel("addGroup:toContainerWithIdentifier:"), objc.Ptr(group), identifier)
}

// Adds the specified contact to the contact store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1403036-addcontact?language=objc
func (s_ SaveRequest) AddContactToContainerWithIdentifier(contact IMutableContact, identifier string) {
	objc.Call[objc.Void](s_, objc.Sel("addContact:toContainerWithIdentifier:"), objc.Ptr(contact), identifier)
}

// Add the specified group to a parent group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1403342-addsubgroup?language=objc
func (s_ SaveRequest) AddSubgroupToGroup(subgroup IGroup, group IGroup) {
	objc.Call[objc.Void](s_, objc.Sel("addSubgroup:toGroup:"), objc.Ptr(subgroup), objc.Ptr(group))
}

// Remove a subgroup from the specified parent group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/1403018-removesubgroup?language=objc
func (s_ SaveRequest) RemoveSubgroupFromGroup(subgroup IGroup, group IGroup) {
	objc.Call[objc.Void](s_, objc.Sel("removeSubgroup:fromGroup:"), objc.Ptr(subgroup), objc.Ptr(group))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/3824780-transactionauthor?language=objc
func (s_ SaveRequest) TransactionAuthor() string {
	rv := objc.Call[string](s_, objc.Sel("transactionAuthor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/3824780-transactionauthor?language=objc
func (s_ SaveRequest) SetTransactionAuthor(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setTransactionAuthor:"), value)
}