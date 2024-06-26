// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Morphology] class.
var MorphologyClass = _MorphologyClass{objc.GetClass("NSMorphology")}

type _MorphologyClass struct {
	objc.Class
}

// An interface definition for the [Morphology] class.
type IMorphology interface {
	objc.IObject
	IsUnspecified() bool
	GrammaticalGender() GrammaticalGender
	SetGrammaticalGender(value GrammaticalGender)
	PartOfSpeech() GrammaticalPartOfSpeech
	SetPartOfSpeech(value GrammaticalPartOfSpeech)
	Number() GrammaticalNumber
	SetNumber(value GrammaticalNumber)
}

// A description of the grammatical properties of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology?language=objc
type Morphology struct {
	objc.Object
}

func MorphologyFrom(ptr unsafe.Pointer) Morphology {
	return Morphology{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MorphologyClass) Alloc() Morphology {
	rv := objc.Call[Morphology](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MorphologyClass) New() Morphology {
	rv := objc.Call[Morphology](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMorphology() Morphology {
	return MorphologyClass.New()
}

func (m_ Morphology) Init() Morphology {
	rv := objc.Call[Morphology](m_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether this instance specifies no particular grammar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746953-unspecified?language=objc
func (m_ Morphology) IsUnspecified() bool {
	rv := objc.Call[bool](m_, objc.Sel("isUnspecified"))
	return rv
}

// The grammatical gender used for inflecting strings with this morphology. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746949-grammaticalgender?language=objc
func (m_ Morphology) GrammaticalGender() GrammaticalGender {
	rv := objc.Call[GrammaticalGender](m_, objc.Sel("grammaticalGender"))
	return rv
}

// The grammatical gender used for inflecting strings with this morphology. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746949-grammaticalgender?language=objc
func (m_ Morphology) SetGrammaticalGender(value GrammaticalGender) {
	objc.Call[objc.Void](m_, objc.Sel("setGrammaticalGender:"), value)
}

// The addressing preferences of the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746954-usermorphology?language=objc
func (mc _MorphologyClass) UserMorphology() Morphology {
	rv := objc.Call[Morphology](mc, objc.Sel("userMorphology"))
	return rv
}

// The addressing preferences of the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746954-usermorphology?language=objc
func Morphology_UserMorphology() Morphology {
	return MorphologyClass.UserMorphology()
}

// The grammatical part of speech used for inflecting strings with this morphology. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746951-partofspeech?language=objc
func (m_ Morphology) PartOfSpeech() GrammaticalPartOfSpeech {
	rv := objc.Call[GrammaticalPartOfSpeech](m_, objc.Sel("partOfSpeech"))
	return rv
}

// The grammatical part of speech used for inflecting strings with this morphology. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746951-partofspeech?language=objc
func (m_ Morphology) SetPartOfSpeech(value GrammaticalPartOfSpeech) {
	objc.Call[objc.Void](m_, objc.Sel("setPartOfSpeech:"), value)
}

// The grammatical number used for inflecting strings with this morphology. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746950-number?language=objc
func (m_ Morphology) Number() GrammaticalNumber {
	rv := objc.Call[GrammaticalNumber](m_, objc.Sel("number"))
	return rv
}

// The grammatical number used for inflecting strings with this morphology. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphology/3746950-number?language=objc
func (m_ Morphology) SetNumber(value GrammaticalNumber) {
	objc.Call[objc.Void](m_, objc.Sel("setNumber:"), value)
}
