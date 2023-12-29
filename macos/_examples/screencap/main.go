package main

/*
// Adjust the CFLAGS and LDFLAGS if needed based on the frameworks and libraries CMSampleBufferCallForEachSample depends on.
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreMedia -framework CoreFoundation
#include <CoreMedia/CoreMedia.h>

// Wrapper function for CMSampleBufferCallForEachSample.
// We're using void* for the callback and refcon parameters to keep it generic. You might need to adjust these types based on how you're going to use them in Go.
OSStatus Go_CMSampleBufferCallForEachSample(CMSampleBufferRef sbuf, void* callback, void* refcon) {
    // Cast the void* back to the appropriate function pointer type.
    OSStatus (*callbackFunc)(CMSampleBufferRef, CMItemCount, void*) = (OSStatus (*)(CMSampleBufferRef, CMItemCount, void*))callback;

    // Call CMSampleBufferCallForEachSample with the provided parameters.
    return CMSampleBufferCallForEachSample(sbuf, callbackFunc, refcon);
}

// Trampoline function declaration.
static OSStatus trampoline(CMSampleBufferRef sbuf, CMItemCount index, void *refcon);

// Function pointer to store the Go callback.
static OSStatus (*goCallback)(CMSampleBufferRef, CMItemCount, void *) = NULL;

// Function to set the Go callback.
void setGoCallback(OSStatus (*cb)(CMSampleBufferRef, CMItemCount, void *)) {
    goCallback = cb;
}

// Trampoline function implementation.
static OSStatus trampoline(CMSampleBufferRef sbuf, CMItemCount index, void *refcon) {
    if (goCallback != NULL) {
        return goCallback(sbuf, index, refcon);
    }
    return -1; // or another appropriate error code
}

// Function pointer type for the callback
typedef OSStatus (*CallbackFunc)(CMSampleBufferRef, CMItemCount, void *);

// Intermediary function
OSStatus callGoFunction(CMSampleBufferRef sbuf, CMItemCount index, void* refcon) {
    // The Go function must be exported using //export
    extern OSStatus mySampleBufferCallback(CMSampleBufferRef, CMItemCount, void*);

    // Call the Go function
    return mySampleBufferCallback(sbuf, index, refcon);
}
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/avfoundation"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/screencapturekit"
	"github.com/progrium/macdriver/objc"
)

func main() {
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)

	captureHandler := &screenCaptureHandler{
		//StreamOutputObject: &screencapturekit.StreamOutputObject{},
		streamDelegate: &screencapturekit.StreamDelegate{},
	}
	captureHandler.streamDelegate.SetStreamDidStopWithError(func(s screencapturekit.Stream, err foundation.Error) {
		fmt.Println("stream stopped", err)
	})
	captureHandler.refreshCapturableWindows()

	sc := screencapturekit.NewStreamConfiguration()
	//cf := screencapturekit.NewContentFilter()
	sc.SetWidth(uint(captureHandler.selectedDisplay.Width()))
	sc.SetHeight(uint(captureHandler.selectedDisplay.Height()))

	cf := captureHandler.GetContentFilter()
	s := screencapturekit.NewStreamWithFilterConfigurationDelegate(cf, sc, captureHandler.streamDelegate)

	var dispatchQueue dispatch.Queue
	dispatchQueue = dispatch.CreateQueue("com.example.queue", dispatch.QueueTypeSerial)
	//dispatchQueue = dispatch.MainQueue()

	StreamOutputClass := objc.NewClass[StreamOutput](
		objc.Sel("stream:didOutputSampleBuffer:ofType:"),
	)
	objc.RegisterClass(StreamOutputClass)
	o := StreamOutputClass.New()

	outputSettings := map[string]objc.IObject{}
	writer = avfoundation.NewAssetWriterInput().InitWithMediaTypeOutputSettings(
		avfoundation.MediaTypeVideo,
		outputSettings,
	)

	err := foundation.Error{}
	ok := s.AddStreamOutputTypeSampleHandlerQueueError(o, screencapturekit.StreamOutputTypeScreen, dispatchQueue, err)
	if !ok {
		fmt.Println("s.AddStreamOutputTypeSampleHandlerQueueError", err)
	}

	fmt.Println("s.StartCaptureWithCompletionHandler")
	s.StartCaptureWithCompletionHandler(func(err foundation.Error) {
		fmt.Println("s.StartCaptureWithCompletionHandler start success", err.IsNil())
	})
	dispatch.Main()
}

type CaptureType int

const (
	CaptureTypeDisplay CaptureType = iota
	CaptureTypeWindow  CaptureType = iota
)

//export mySampleBufferCallback
func mySampleBufferCallback(sbuf C.CMSampleBufferRef, index C.CMItemCount, refcon unsafe.Pointer) C.OSStatus {
	// Your callback logic here
	return 0
}

func callForEachSample(sbuf C.CMSampleBufferRef, refcon unsafe.Pointer) C.OSStatus {
	return C.Go_CMSampleBufferCallForEachSample(sbuf, C.callGoFunction, refcon)
}

func (h *screenCaptureHandler) refreshCapturableWindows() {
	fmt.Println("listing sharable content")
	ch := make(chan struct{})
	screencapturekit.ShareableContentClass.GetShareableContentWithCompletionHandler(func(sc screencapturekit.ShareableContent, err foundation.Error) {
		defer close(ch)
		if !err.IsNil() {
			fmt.Println("error listing sharable content:", err.LocalizedDescription())
			os.Exit(1)
		}
		// h.availabileDisplays = sc.Displays()
		for _, d := range sc.Displays() {
			h.availabileDisplays = append(h.availabileDisplays, d)
			objc.Retain(&d)
		}
		if h.selectedDisplay == nil && len(h.availabileDisplays) > 0 {
			h.selectedDisplay = &h.availabileDisplays[0]
			objc.Retain(h.selectedDisplay)
		}

		ws := sc.Windows()
		h.availableWindows = append(h.availableWindows, ws[0])
		fmt.Println(len(ws), "windows")
		for _, w := range sc.Windows() {
			//fmt.Println(i, w.IsOnScreen(), w.IsNil(), w.IsProxy(), w.Title(), w.OwningApplication().ApplicationName())
			// 	h.availableWindows = append(h.availableWindows, w)
			// 	fmt.Println("window", i, w.Description())
			if w.IsOnScreen() && w.OwningApplication().ApplicationName() != "Finder" {
				h.capturedWindows = append(h.capturedWindows, w)
				fmt.Println("captured window", w.Title())
				w.Retain()
			}

		}
		if h.selectedWindow == nil && len(h.availableWindows) > 0 {
			h.selectedWindow = &h.availableWindows[0]
		}
		objc.Retain(h.selectedWindow)
		// for _, app := range sc.Applications() {
		// 	fmt.Println("app", app.ApplicationName())
		// }
		fmt.Println("done listing sharable content")
	})
	fmt.Println("waiting for listing sharable content")
	<-ch
	fmt.Println("done waiting for listing sharable content")
}

type screenCaptureHandler struct {
	objc.Object
	//*screencapturekit.StreamOutputObject

	streamDelegate     *screencapturekit.StreamDelegate
	availabileDisplays []screencapturekit.Display
	availableWindows   []screencapturekit.Window

	selectedDisplay *screencapturekit.Display
	selectedWindow  *screencapturekit.Window

	capturedWindows []screencapturekit.IWindow

	captureType CaptureType
}

func (sh *screenCaptureHandler) GetContentFilter() screencapturekit.ContentFilter {
	filter := screencapturekit.NewContentFilter()

	fmt.Println("GetContentFilter", sh.captureType)
	switch sh.captureType {
	case CaptureTypeDisplay:
		display := sh.selectedDisplay
		for _, w := range sh.availableWindows {
			_ = w
			// if len(sh.capturedWindows) < 3 {
			sh.capturedWindows = append(sh.capturedWindows, w)
			// objc.Retain(&w)
			// }
		}
		//filter = screencapturekit.NewContentFilterWithDisplayIncludingWindows(display, sh.capturedWindows)
		_ = display
		filter = screencapturekit.NewContentFilterWithDesktopIndependentWindow(sh.capturedWindows[0])
	case CaptureTypeWindow:
	}
	return filter
}

type StreamOutput struct {
	objc.Object
	//screencpaturekit.SCStreamOutput
}

var _ screencapturekit.PStreamOutput = (*StreamOutput)(nil)

// HasStreamDidOutputSampleBufferOfType

func (s StreamOutput) HasStreamDidOutputSampleBufferOfType() bool {
	return s.RespondsToSelector(objc.Sel("stream:didOutputSampleBuffer:ofType:"))
}

var writer avfoundation.AssetWriterInput

func (s StreamOutput) StreamDidOutputSampleBufferOfType(stream screencapturekit.Stream, sampleBuffer coremedia.SampleBufferRef, type_ screencapturekit.StreamOutputType) {
	fmt.Println("StreamDidOutputSampleBufferOfType", stream, sampleBuffer, type_)
	if type_ == screencapturekit.StreamOutputTypeScreen {
		fmt.Println("StreamDidOutputSampleBufferOfType screencapturekit.StreamOutputTypeScreen")
		// save data into a file:

	}
}

// var _ screencapturekit.PStreamOutput = (*screenCaptureHandler)(nil)
// var _ screencapturekit.PStreamDelegate = (*screenCaptureHandler)(nil)

// // StreamOutput methods

// func (sh *screenCaptureHandler) HasStreamDidOutputSampleBufferOfType() bool {
// 	panic(errors.New("*streamHandler.HasStreamDidOutputSampleBufferOfType not implemented"))
// }

// func (sh *screenCaptureHandler) StreamDidOutputSampleBufferOfType(s screencapturekit.Stream, buf coremedia.SampleBufferRef, out screencapturekit.StreamOutputType) {
// 	panic(errors.New("*streamHandler.StreamDidOutputSampleBufferOfType not implemented"))
// }

// // StreamDelegate methods

// func (sh *screenCaptureHandler) HasStreamDidStopWithError() bool {
// 	panic(errors.New("*streamHandler.HasStreamDidStopWithError not implemented"))
// 	return true
// }

// func (sh *screenCaptureHandler) StreamDidStopWithError(s screencapturekit.Stream, err foundation.Error) {
// 	fmt.Println("stream stopped", err)
// 	panic(errors.New("*streamHandler.StreamDidStopWithError not implemented"))
// }
