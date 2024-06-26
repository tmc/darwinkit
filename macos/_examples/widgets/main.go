package main

import (
	"fmt"
	"time"

	"github.com/progrium/darwinkit/dispatch"
	"github.com/progrium/darwinkit/helper/action"
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	w := appkit.NewWindowWithSize(600, 400)
	objc.Retain(&w)

	w.SetTitle("Test widgets")

	filePathField := appkit.NewTextFieldWithFrame(rectOf(10, 330, 200, 20))
	filePathField.SetEditable(false)
	w.ContentView().AddSubview(filePathField)

	saveButton := appkit.NewButtonWithTitle("Save...")
	saveButton.SetFrame(rectOf(250, 330, 80, 20))
	action.Set(saveButton, func(sender objc.Object) {
		savePanel := appkit.SavePanelClass.New()
		// if savePanel.RunModal() == appkit.ModalResponseOK {
		// 	filePathField.SetStringValue(savePanel.URL().Path())
		// }
		savePanel.BeginWithCompletionHandler(func(response appkit.ModalResponse) {
			if response == appkit.ModalResponseOK {
				filePathField.SetStringValue(savePanel.URL().Path())
			}
		})
	})
	w.ContentView().AddSubview(saveButton)

	presentationTF := appkit.NewTextFieldWithFrame(rectOf(10, 290, 100, 25))
	w.ContentView().AddSubview(presentationTF)

	stepper := appkit.NewStepperWithFrame(rectOf(130, 290, 16, 25))
	stepper.SetDoubleValue(100)
	w.ContentView().AddSubview(stepper)

	colorWell := appkit.NewColorWellWithFrame(rectOf(160, 290, 30, 25))
	w.ContentView().AddSubview(colorWell)

	comboBox := appkit.NewComboBoxWithFrame(rectOf(210, 290, 100, 25))
	comboBox.AddItemsWithObjectValues([]objc.IObject{
		foundation.NewStringWithString("Test1"),
		foundation.NewStringWithString("Test2"),
	})
	comboBox.SelectItemAtIndex(0)
	w.ContentView().AddSubview(comboBox)

	slider := appkit.NewSliderWithFrame(rectOf(330, 290, 100, 25))
	action.Set(slider, func(sender objc.Object) {
		presentationTF.SetDoubleValue(slider.DoubleValue())
	})
	slider.SetMaxValue(10)
	w.ContentView().AddSubview(slider)

	datePicker := appkit.NewDatePickerWithFrame(rectOf(450, 290, 140, 25))
	w.ContentView().AddSubview(datePicker)

	// buttons
	cb := appkit.NewCheckBox("check box")
	cb.SetFrame(rectOf(10, 250, 80, 25))
	w.ContentView().AddSubview(cb)

	rb := appkit.NewRadioButton("radio button")
	rb.SetFrame(rectOf(150, 250, 120, 25))
	w.ContentView().AddSubview(rb)

	sw := appkit.NewSwitchWithFrame(rectOf(260, 250, 120, 25))
	w.ContentView().AddSubview(sw)

	li := appkit.NewLevelIndicatorWithFrame(rectOf(370, 250, 120, 25))
	li.SetCriticalValue(4)
	li.SetDoubleValue(3)
	w.ContentView().AddSubview(li)

	btn := appkit.NewButtonWithTitle("change color")
	btn.SetFrame(rectOf(10, 160, 120, 25))
	w.ContentView().AddSubview(btn)

	quitBtn := appkit.NewButtonWithTitle("Quit")
	quitBtn.SetFrame(rectOf(10, 130, 80, 25))
	action.Set(quitBtn, func(sender objc.Object) {
		app.Terminate(nil)
	})
	w.ContentView().AddSubview(quitBtn)

	// text field
	tf := appkit.NewTextField()
	w.ContentView().AddSubview(tf)
	tf.SetFrame(rectOf(10, 100, 150, 25))

	// label
	label := appkit.NewLabel("")
	label.SetFrame(rectOf(170, 100, 150, 25))
	w.ContentView().AddSubview(label)
	tfd := &appkit.TextFieldDelegate{}
	tfd.SetControlTextDidChange(func(obj foundation.Notification) {
		dispatch.MainQueue().DispatchAsync(func() {
			label.SetStringValue(tf.StringValue())
		})
	})
	tf.SetDelegate(tfd)
	action.Set(btn, func(sender objc.Object) {
		label.SetTextColor(appkit.ColorClass.RedColor())
	})

	// password
	stf := appkit.NewSecureTextField()
	stf.SetFrame(rectOf(340, 100, 150, 25))
	tfd2 := &appkit.TextFieldDelegate{}
	tfd2.SetControlTextDidChange(func(obj foundation.Notification) {
		dispatch.MainQueue().DispatchAsync(func() {
			label.SetStringValue(stf.StringValue())
		})
	})
	stf.SetDelegate(tfd2)
	w.ContentView().AddSubview(stf)

	// progress indicator
	indicator := appkit.NewProgressIndicatorWithFrame(rectOf(10, 70, 350, 25))
	indicator.SetMinValue(0)
	indicator.SetMaxValue(1)
	indicator.SetIndeterminate(false)
	indicator.SetDisplayedWhenStopped(false)
	w.ContentView().AddSubview(indicator)
	go func() {
		dispatch.MainQueue().DispatchAsync(func() {
			indicator.StartAnimation(indicator)
		})
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			dispatch.MainQueue().DispatchAsync(func() {
				indicator.SetDoubleValue(0.1 * float64(i))
			})
		}
		dispatch.MainQueue().DispatchAsync(func() {
			indicator.StopAnimation(indicator)
		})
	}()

	// text view & scroll view
	sv := appkit.TextView_ScrollableTextView()
	sv.SetFrame(rectOf(10, 200, 200, 30))
	appkit.TextViewFrom(sv.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv)

	sv2 := appkit.TextView_ScrollableTextView()
	sv2.SetFrame(rectOf(250, 200, 200, 30))
	appkit.TextViewFrom(sv2.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv2)

	wd := &appkit.WindowDelegate{}
	wd.SetWindowDidMove(func(notification foundation.Notification) {
		frame := w.Frame()
		origin := frame.Origin
		fmt.Println("window move to ", origin.X, origin.Y)
	})
	w.SetDelegate(wd)
	w.Center()
	w.MakeKeyAndOrderFront(nil)

	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
}

func rectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}
