package graphical

import (
	"runtime"

	"github.com/ka1i/matrix/internal/app/graphical/views"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

func init() {
	runtime.LockOSThread()
}

func UserInterface() {
	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		views.WaterMark("Monaco", 16, 15, 15)
	})
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
