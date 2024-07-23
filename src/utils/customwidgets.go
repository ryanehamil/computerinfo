package utils

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func setClipboard(text string) {
	fyne.CurrentApp().Driver().AllWindows()[0].Clipboard().SetContent(text)
}

type ClickableText struct {
	widget.BaseWidget
	tapBG *canvas.Rectangle
	Text  *canvas.Text
}

func NewClickableText(text string) *ClickableText {
	object := &ClickableText{
		Text:  canvas.NewText(text, Theme{}.Color(fyne.ThemeColorName("foreground"), fyne.CurrentApp().Settings().ThemeVariant())),
		tapBG: canvas.NewRectangle(Theme{}.Color(fyne.ThemeColorName("background"), fyne.CurrentApp().Settings().ThemeVariant())),
	}

	object.Text.TextStyle.Bold = true

	object.ExtendBaseWidget(object)

	return object
}

func (ct *ClickableText) CreateRenderer() fyne.WidgetRenderer {

	object := container.NewCenter()

	ct.tapBG.SetMinSize(ct.Text.MinSize().Add(fyne.NewSize(10, 2)))
	ct.tapBG.CornerRadius = 10

	object.Add(ct.tapBG)
	object.Add(ct.Text)

	return widget.NewSimpleRenderer(object)
}

func (ct *ClickableText) Tapped(_ *fyne.PointEvent) {
	setClipboard(ct.Text.Text)

	// Set the text color to green when tapped, and reset it after 1 second
	ct.tapBG.FillColor = color.RGBA{0, 120, 20, 200}

	go func() {
		<-time.After(time.Millisecond * 400)
		ct.tapBG.FillColor = theme.Color("background")
		ct.Refresh()
	}()

	ct.Refresh()
}
