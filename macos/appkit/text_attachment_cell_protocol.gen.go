// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A set of methods that declares the interface for objects that draw text attachment icons and handle mouse events on their icons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell?language=objc
type PTextAttachmentCell interface {
	// optional
	HighlightWithFrameInView(flag bool, cellFrame foundation.Rect, controlView View)
	HasHighlightWithFrameInView() bool

	// optional
	CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer TextContainer, lineFrag foundation.Rect, position foundation.Point, charIndex uint) foundation.Rect
	HasCellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool

	// optional
	DrawWithFrameInViewCharacterIndex(cellFrame foundation.Rect, controlView View, charIndex uint)
	HasDrawWithFrameInViewCharacterIndex() bool

	// optional
	CellSize() foundation.Size
	HasCellSize() bool

	// optional
	CellBaselineOffset() foundation.Point
	HasCellBaselineOffset() bool

	// optional
	DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame foundation.Rect, controlView View, charIndex uint, layoutManager LayoutManager)
	HasDrawWithFrameInViewCharacterIndexLayoutManager() bool

	// optional
	TrackMouseInRectOfViewUntilMouseUp(theEvent Event, cellFrame foundation.Rect, controlView View, flag bool) bool
	HasTrackMouseInRectOfViewUntilMouseUp() bool

	// optional
	WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent Event, cellFrame foundation.Rect, controlView View, charIndex uint) bool
	HasWantsToTrackMouseForEventInRectOfViewAtCharacterIndex() bool

	// optional
	WantsToTrackMouse() bool
	HasWantsToTrackMouse() bool

	// optional
	TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent Event, cellFrame foundation.Rect, controlView View, charIndex uint, flag bool) bool
	HasTrackMouseInRectOfViewAtCharacterIndexUntilMouseUp() bool

	// optional
	DrawWithFrameInView(cellFrame foundation.Rect, controlView View)
	HasDrawWithFrameInView() bool

	// optional
	SetAttachment(value TextAttachment)
	HasSetAttachment() bool

	// optional
	Attachment() TextAttachment
	HasAttachment() bool
}

// ensure impl type implements protocol interface
var _ PTextAttachmentCell = (*TextAttachmentCellObject)(nil)

// A concrete type for the [PTextAttachmentCell] protocol.
type TextAttachmentCellObject struct {
	objc.Object
}

func (t_ TextAttachmentCellObject) HasHighlightWithFrameInView() bool {
	return t_.RespondsToSelector(objc.Sel("highlight:withFrame:inView:"))
}

// Draws the receiver’s image with optional highlighting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508384-highlight?language=objc
func (t_ TextAttachmentCellObject) HighlightWithFrameInView(flag bool, cellFrame foundation.Rect, controlView View) {
	objc.Call[objc.Void](t_, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}

func (t_ TextAttachmentCellObject) HasCellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool {
	return t_.RespondsToSelector(objc.Sel("cellFrameForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"))
}

// Returns the frame of the cell to draw at the specified position in a text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508402-cellframefortextcontainer?language=objc
func (t_ TextAttachmentCellObject) CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer TextContainer, lineFrag foundation.Rect, position foundation.Point, charIndex uint) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("cellFrameForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), textContainer, lineFrag, position, charIndex)
	return rv
}

func (t_ TextAttachmentCellObject) HasDrawWithFrameInViewCharacterIndex() bool {
	return t_.RespondsToSelector(objc.Sel("drawWithFrame:inView:characterIndex:"))
}

// Draws the cell's image at the specified index point in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508412-drawwithframe?language=objc
func (t_ TextAttachmentCellObject) DrawWithFrameInViewCharacterIndex(cellFrame foundation.Rect, controlView View, charIndex uint) {
	objc.Call[objc.Void](t_, objc.Sel("drawWithFrame:inView:characterIndex:"), cellFrame, controlView, charIndex)
}

func (t_ TextAttachmentCellObject) HasCellSize() bool {
	return t_.RespondsToSelector(objc.Sel("cellSize"))
}

// Returns the size of the attachment’s icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508405-cellsize?language=objc
func (t_ TextAttachmentCellObject) CellSize() foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("cellSize"))
	return rv
}

func (t_ TextAttachmentCellObject) HasCellBaselineOffset() bool {
	return t_.RespondsToSelector(objc.Sel("cellBaselineOffset"))
}

// Returns the text position where you draw the attachment cell’s image, relative to the current point established in the glyph layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508420-cellbaselineoffset?language=objc
func (t_ TextAttachmentCellObject) CellBaselineOffset() foundation.Point {
	rv := objc.Call[foundation.Point](t_, objc.Sel("cellBaselineOffset"))
	return rv
}

func (t_ TextAttachmentCellObject) HasDrawWithFrameInViewCharacterIndexLayoutManager() bool {
	return t_.RespondsToSelector(objc.Sel("drawWithFrame:inView:characterIndex:layoutManager:"))
}

// Draws the cell's image using the specified layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508385-drawwithframe?language=objc
func (t_ TextAttachmentCellObject) DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame foundation.Rect, controlView View, charIndex uint, layoutManager LayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("drawWithFrame:inView:characterIndex:layoutManager:"), cellFrame, controlView, charIndex, layoutManager)
}

func (t_ TextAttachmentCellObject) HasTrackMouseInRectOfViewUntilMouseUp() bool {
	return t_.RespondsToSelector(objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"))
}

// Handles a mouse-down event on the cell's image, and optionally waits until a mouse-up event [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508418-trackmouse?language=objc
func (t_ TextAttachmentCellObject) TrackMouseInRectOfViewUntilMouseUp(theEvent Event, cellFrame foundation.Rect, controlView View, flag bool) bool {
	rv := objc.Call[bool](t_, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), theEvent, cellFrame, controlView, flag)
	return rv
}

func (t_ TextAttachmentCellObject) HasWantsToTrackMouseForEventInRectOfViewAtCharacterIndex() bool {
	return t_.RespondsToSelector(objc.Sel("wantsToTrackMouseForEvent:inRect:ofView:atCharacterIndex:"))
}

// Allows an attachment to specify the events for which it tracks the mouse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508399-wantstotrackmouseforevent?language=objc
func (t_ TextAttachmentCellObject) WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent Event, cellFrame foundation.Rect, controlView View, charIndex uint) bool {
	rv := objc.Call[bool](t_, objc.Sel("wantsToTrackMouseForEvent:inRect:ofView:atCharacterIndex:"), theEvent, cellFrame, controlView, charIndex)
	return rv
}

func (t_ TextAttachmentCellObject) HasWantsToTrackMouse() bool {
	return t_.RespondsToSelector(objc.Sel("wantsToTrackMouse"))
}

// Returns a Boolean value that indicates whether the attachment handles mouse events occurring over its image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508415-wantstotrackmouse?language=objc
func (t_ TextAttachmentCellObject) WantsToTrackMouse() bool {
	rv := objc.Call[bool](t_, objc.Sel("wantsToTrackMouse"))
	return rv
}

func (t_ TextAttachmentCellObject) HasTrackMouseInRectOfViewAtCharacterIndexUntilMouseUp() bool {
	return t_.RespondsToSelector(objc.Sel("trackMouse:inRect:ofView:atCharacterIndex:untilMouseUp:"))
}

// Handles a mouse-down event on the image at the specified character position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508380-trackmouse?language=objc
func (t_ TextAttachmentCellObject) TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent Event, cellFrame foundation.Rect, controlView View, charIndex uint, flag bool) bool {
	rv := objc.Call[bool](t_, objc.Sel("trackMouse:inRect:ofView:atCharacterIndex:untilMouseUp:"), theEvent, cellFrame, controlView, charIndex, flag)
	return rv
}

func (t_ TextAttachmentCellObject) HasDrawWithFrameInView() bool {
	return t_.RespondsToSelector(objc.Sel("drawWithFrame:inView:"))
}

// Draws the cell's image in the specified rectangle of the currently focused view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508392-drawwithframe?language=objc
func (t_ TextAttachmentCellObject) DrawWithFrameInView(cellFrame foundation.Rect, controlView View) {
	objc.Call[objc.Void](t_, objc.Sel("drawWithFrame:inView:"), cellFrame, controlView)
}

func (t_ TextAttachmentCellObject) HasSetAttachment() bool {
	return t_.RespondsToSelector(objc.Sel("setAttachment:"))
}

// Returns the text attachment object that owns the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508396-attachment?language=objc
func (t_ TextAttachmentCellObject) SetAttachment(value TextAttachment) {
	objc.Call[objc.Void](t_, objc.Sel("setAttachment:"), value)
}

func (t_ TextAttachmentCellObject) HasAttachment() bool {
	return t_.RespondsToSelector(objc.Sel("attachment"))
}

// Returns the text attachment object that owns the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508396-attachment?language=objc
func (t_ TextAttachmentCellObject) Attachment() TextAttachment {
	rv := objc.Call[TextAttachment](t_, objc.Sel("attachment"))
	return rv
}
