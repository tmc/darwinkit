// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Sound] class.
var SoundClass = _SoundClass{objc.GetClass("NSSound")}

type _SoundClass struct {
	objc.Class
}

// An interface definition for the [Sound] class.
type ISound interface {
	objc.IObject
	Stop() bool
	Pause() bool
	Resume() bool
	WriteToPasteboard(pasteboard IPasteboard)
	SetName(string_ SoundName) bool
	Play() bool
	Volume() float32
	SetVolume(value float32)
	Duration() foundation.TimeInterval
	IsPlaying() bool
	Delegate() SoundDelegateObject
	SetDelegate(value PSoundDelegate)
	SetDelegateObject(valueObject objc.IObject)
	Name() SoundName
	Loops() bool
	SetLoops(value bool)
	PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier
	SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier)
	CurrentTime() foundation.TimeInterval
	SetCurrentTime(value foundation.TimeInterval)
}

// A simple interface for loading and playing audio files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound?language=objc
type Sound struct {
	objc.Object
}

func SoundFrom(ptr unsafe.Pointer) Sound {
	return Sound{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ Sound) InitWithData(data []byte) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes the receiver with a given audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477292-initwithdata?language=objc
func NewSoundWithData(data []byte) Sound {
	instance := SoundClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (s_ Sound) InitWithContentsOfFileByReference(path string, byRef bool) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithContentsOfFile:byReference:"), path, byRef)
	return rv
}

// Initializes the receiver with the audio data located at a given filepath. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477274-initwithcontentsoffile?language=objc
func NewSoundWithContentsOfFileByReference(path string, byRef bool) Sound {
	instance := SoundClass.Alloc().InitWithContentsOfFileByReference(path, byRef)
	instance.Autorelease()
	return instance
}

func (s_ Sound) InitWithPasteboard(pasteboard IPasteboard) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithPasteboard:"), pasteboard)
	return rv
}

// Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by NSSound. NSSound expects the data to have a proper magic number, sound header, and data for the formats it supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477294-initwithpasteboard?language=objc
func NewSoundWithPasteboard(pasteboard IPasteboard) Sound {
	instance := SoundClass.Alloc().InitWithPasteboard(pasteboard)
	instance.Autorelease()
	return instance
}

func (s_ Sound) InitWithContentsOfURLByReference(url foundation.IURL, byRef bool) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithContentsOfURL:byReference:"), url, byRef)
	return rv
}

// Initializes the receiver with the audio data located at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477288-initwithcontentsofurl?language=objc
func NewSoundWithContentsOfURLByReference(url foundation.IURL, byRef bool) Sound {
	instance := SoundClass.Alloc().InitWithContentsOfURLByReference(url, byRef)
	instance.Autorelease()
	return instance
}

func (sc _SoundClass) Alloc() Sound {
	rv := objc.Call[Sound](sc, objc.Sel("alloc"))
	return rv
}

func (sc _SoundClass) New() Sound {
	rv := objc.Call[Sound](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSound() Sound {
	return SoundClass.New()
}

func (s_ Sound) Init() Sound {
	rv := objc.Call[Sound](s_, objc.Sel("init"))
	return rv
}

// Indicates whether the receiver can create an instance of itself from the data in a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477276-caninitwithpasteboard?language=objc
func (sc _SoundClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Call[bool](sc, objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}

// Indicates whether the receiver can create an instance of itself from the data in a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477276-caninitwithpasteboard?language=objc
func Sound_CanInitWithPasteboard(pasteboard IPasteboard) bool {
	return SoundClass.CanInitWithPasteboard(pasteboard)
}

// Concludes audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477282-stop?language=objc
func (s_ Sound) Stop() bool {
	rv := objc.Call[bool](s_, objc.Sel("stop"))
	return rv
}

// Pauses audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477307-pause?language=objc
func (s_ Sound) Pause() bool {
	rv := objc.Call[bool](s_, objc.Sel("pause"))
	return rv
}

// Resumes audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477336-resume?language=objc
func (s_ Sound) Resume() bool {
	rv := objc.Call[bool](s_, objc.Sel("resume"))
	return rv
}

// Writes the receiver’s data to a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477338-writetopasteboard?language=objc
func (s_ Sound) WriteToPasteboard(pasteboard IPasteboard) {
	objc.Call[objc.Void](s_, objc.Sel("writeToPasteboard:"), pasteboard)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477286-setname?language=objc
func (s_ Sound) SetName(string_ SoundName) bool {
	rv := objc.Call[bool](s_, objc.Sel("setName:"), string_)
	return rv
}

// Returns the NSSound instance associated with a given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477318-soundnamed?language=objc
func (sc _SoundClass) SoundNamed(name SoundName) Sound {
	rv := objc.Call[Sound](sc, objc.Sel("soundNamed:"), name)
	return rv
}

// Returns the NSSound instance associated with a given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477318-soundnamed?language=objc
func Sound_SoundNamed(name SoundName) Sound {
	return SoundClass.SoundNamed(name)
}

// Initiates audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477322-play?language=objc
func (s_ Sound) Play() bool {
	rv := objc.Call[bool](s_, objc.Sel("play"))
	return rv
}

// The volume of the sound. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477315-volume?language=objc
func (s_ Sound) Volume() float32 {
	rv := objc.Call[float32](s_, objc.Sel("volume"))
	return rv
}

// The volume of the sound. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477315-volume?language=objc
func (s_ Sound) SetVolume(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setVolume:"), value)
}

// The duration of the sound, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477313-duration?language=objc
func (s_ Sound) Duration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("duration"))
	return rv
}

// A Boolean that indicates whether the sound is playing its audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477302-playing?language=objc
func (s_ Sound) IsPlaying() bool {
	rv := objc.Call[bool](s_, objc.Sel("isPlaying"))
	return rv
}

// The sound’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (s_ Sound) Delegate() SoundDelegateObject {
	rv := objc.Call[SoundDelegateObject](s_, objc.Sel("delegate"))
	return rv
}

// The sound’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (s_ Sound) SetDelegate(value PSoundDelegate) {
	po0 := objc.WrapAsProtocol("NSSoundDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The sound’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (s_ Sound) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), valueObject)
}

// The name assigned to the sound. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477296-name?language=objc
func (s_ Sound) Name() SoundName {
	rv := objc.Call[SoundName](s_, objc.Sel("name"))
	return rv
}

// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc
func (s_ Sound) Loops() bool {
	rv := objc.Call[bool](s_, objc.Sel("loops"))
	return rv
}

// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc
func (s_ Sound) SetLoops(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setLoops:"), value)
}

// Identifies the sound’s output device [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477284-playbackdeviceidentifier?language=objc
func (s_ Sound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	rv := objc.Call[SoundPlaybackDeviceIdentifier](s_, objc.Sel("playbackDeviceIdentifier"))
	return rv
}

// Identifies the sound’s output device [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477284-playbackdeviceidentifier?language=objc
func (s_ Sound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier) {
	objc.Call[objc.Void](s_, objc.Sel("setPlaybackDeviceIdentifier:"), value)
}

// The sound’s playback progress, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477320-currenttime?language=objc
func (s_ Sound) CurrentTime() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("currentTime"))
	return rv
}

// The sound’s playback progress, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477320-currenttime?language=objc
func (s_ Sound) SetCurrentTime(value foundation.TimeInterval) {
	objc.Call[objc.Void](s_, objc.Sel("setCurrentTime:"), value)
}

// Provides the file types the NSSound class understands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477290-soundunfilteredtypes?language=objc
func (sc _SoundClass) SoundUnfilteredTypes() []string {
	rv := objc.Call[[]string](sc, objc.Sel("soundUnfilteredTypes"))
	return rv
}

// Provides the file types the NSSound class understands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477290-soundunfilteredtypes?language=objc
func Sound_SoundUnfilteredTypes() []string {
	return SoundClass.SoundUnfilteredTypes()
}
