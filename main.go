package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ryanehamil/computerinfo/src/computerinfo"
)

func main() {
	a := app.New()

	// create fyne resource for icon.png
	// a.SetIcon(resourceIconPng)

	w := a.NewWindow("Computer Info")

	w.Resize(fyne.NewSize(800, 200))

	w.SetContent(computerinfo.CanvasObject())

	w.ShowAndRun()
}
