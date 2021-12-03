package component

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
)

var MI float64 = 6.0

func ViewSize(text string, font string, size float64) (width, height float64) {
	t := cocoa.NSTextView_Init(core.Rect(0, 0, 0, 0))
	t.SetString(text)
	t.SetFont(cocoa.Font(font, size))
	t.LayoutManager().EnsureLayoutForTextContainer(t.TextContainer())
	rect := t.LayoutManager().UsedRectForTextContainer(t.TextContainer())
	return rect.Size.Width + MI*2, rect.Size.Height
}
