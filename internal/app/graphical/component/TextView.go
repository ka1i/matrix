package component

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
)

func TextView(text string, font string, size float64, viewW, viewH float64) cocoa.NSTextView {
	t := cocoa.NSTextView_Init(core.Rect(MI, 0, viewW, viewH))
	t.SetString(text)
	t.SetFont(cocoa.Font(font, size))
	t.SetTextColor(cocoa.Color(1, 1, 1, 0.85))
	t.SetDrawsBackground(true)
	t.SetBackgroundColor(cocoa.Color(0.3, 0.3, 0.3, 0.33))
	t.SetEditable(false)
	t.SetSelectable(false)
	t.SetImportsGraphics(false)
	t.SetDrawsBackground(false)
	return t
}
