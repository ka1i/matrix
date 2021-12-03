package views

import (
	"fmt"
	"log"

	"github.com/ka1i/matrix/internal/app/graphical/component"
	"github.com/ka1i/matrix/internal/app/graphical/menubar"
	"github.com/ka1i/matrix/pkg/version"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
)

func WaterMark(fontName string, fontSize, offsetX, offsetY float64) {
	watermark := fmt.Sprintf("          Matrix Pro Inside Preview\nMardan(ka1i) Build Version:%s", version.Version.ToString())

	viewW, viewH := component.ViewSize(watermark, fontName, fontSize)

	windowSize := core.Rect(offsetX, offsetY, viewW, viewH)

	log.Printf("Create WaterMark :%v\n", windowSize)

	t := component.TextView(watermark, fontName, fontSize, viewW, viewH)

	c := cocoa.NSView_Init(core.Rect(0, 0, viewW, viewH))
	c.SetBackgroundColor(cocoa.Color(0, 0, 0, 0.223))
	c.SetWantsLayer(true)
	c.Layer().SetCornerRadius(16.0)
	c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

	w := cocoa.NSWindow_Init(windowSize,
		cocoa.NSBorderlessWindowMask,
		cocoa.NSBackingStoreBuffered,
		false,
	)
	w.SetContentView(c)
	w.SetTitlebarAppearsTransparent(true)
	w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
	w.SetOpaque(false)
	w.SetBackgroundColor(cocoa.NSColor_Clear())
	w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
	w.SetFrameDisplay(windowSize, true)
	w.MakeKeyAndOrderFront(nil)

	menubar.InitBar(&w)
}
