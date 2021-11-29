package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ryanehamil/computerinfo/src/computerinfo"
)

func main() {
	a := app.New()

	// create fyne resource for icon.png
	icon, _ := fyne.LoadResourceFromPath("./Icon.png")
	a.SetIcon(icon)

	w := a.NewWindow("Computer Info")

	w.Resize(fyne.NewSize(800, 200))

	w.SetContent(computerinfo.Container())

	w.ShowAndRun()
}
