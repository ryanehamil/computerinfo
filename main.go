package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ryanehamil/computerinfo/src/computerinfo"
	"github.com/ryanehamil/computerinfo/src/utils"
)

func main() {
	App := app.New()
	App.Settings().SetTheme(&utils.Theme{})

	Window := App.NewWindow("Computer Info")

	Window.Resize(fyne.NewSize(800, 200))

	Window.SetContent(computerinfo.Container())

	Window.ShowAndRun()
}
