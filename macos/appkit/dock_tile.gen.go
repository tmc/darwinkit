// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [DockTile] class.
var DockTileClass = _DockTileClass{objc.GetClass("NSDockTile")}

type _DockTileClass struct {
	objc.Class
}

// An interface definition for the [DockTile] class.
type IDockTile interface {
	objc.IObject
	Display()
	ContentView() View
	SetContentView(value IView)
	ShowsApplicationBadge() bool
	SetShowsApplicationBadge(value bool)
	Owner() objc.Object
	BadgeLabel() string
	SetBadgeLabel(value string)
	Size() foundation.Size
}

// The visual representation of your app’s miniaturized windows and app icon as they appear in the Dock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile?language=objc
type DockTile struct {
	objc.Object
}

func DockTileFrom(ptr unsafe.Pointer) DockTile {
	return DockTile{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DockTileClass) Alloc() DockTile {
	rv := objc.Call[DockTile](dc, objc.Sel("alloc"))
	return rv
}

func (dc _DockTileClass) New() DockTile {
	rv := objc.Call[DockTile](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDockTile() DockTile {
	return DockTileClass.New()
}

func (d_ DockTile) Init() DockTile {
	rv := objc.Call[DockTile](d_, objc.Sel("init"))
	return rv
}

// Redraws the dock tile’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1527292-display?language=objc
func (d_ DockTile) Display() {
	objc.Call[objc.Void](d_, objc.Sel("display"))
}

// The view to use for drawing the dock tile contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1525995-contentview?language=objc
func (d_ DockTile) ContentView() View {
	rv := objc.Call[View](d_, objc.Sel("contentView"))
	return rv
}

// The view to use for drawing the dock tile contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1525995-contentview?language=objc
func (d_ DockTile) SetContentView(value IView) {
	objc.Call[objc.Void](d_, objc.Sel("setContentView:"), value)
}

// A Boolean showing whether the tile is badged with the application’s icon [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1528057-showsapplicationbadge?language=objc
func (d_ DockTile) ShowsApplicationBadge() bool {
	rv := objc.Call[bool](d_, objc.Sel("showsApplicationBadge"))
	return rv
}

// A Boolean showing whether the tile is badged with the application’s icon [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1528057-showsapplicationbadge?language=objc
func (d_ DockTile) SetShowsApplicationBadge(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setShowsApplicationBadge:"), value)
}

// The object represented by the dock tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1533723-owner?language=objc
func (d_ DockTile) Owner() objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("owner"))
	return rv
}

// The string to be displayed in the tile’s badging area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1524433-badgelabel?language=objc
func (d_ DockTile) BadgeLabel() string {
	rv := objc.Call[string](d_, objc.Sel("badgeLabel"))
	return rv
}

// The string to be displayed in the tile’s badging area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1524433-badgelabel?language=objc
func (d_ DockTile) SetBadgeLabel(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setBadgeLabel:"), value)
}

// The size of the tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/1534239-size?language=objc
func (d_ DockTile) Size() foundation.Size {
	rv := objc.Call[foundation.Size](d_, objc.Sel("size"))
	return rv
}
