package utils

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func setClipboard(text string) {
	fyne.CurrentApp().Driver().AllWindows()[0].Clipboard().SetContent(text)
}

type TextButton struct {
	*canvas.Text
}

func NewTextButton(text string) *TextButton {
	return &TextButton{
		Text: canvas.NewText(text, theme.Color("foreground")),
	}
}

func (tb *TextButton) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(tb.Text)
}

func (tb *TextButton) Tapped(_ *fyne.PointEvent) {
	setClipboard(tb.Text.Text)

	// Set the text color to green when tapped, and reset it after 1 second
	tb.Color = theme.Color("success")

	go func() {
		<-time.After(time.Second)
		tb.Color = theme.Color("foreground")
		tb.Refresh()
	}()

	tb.Refresh()
}
