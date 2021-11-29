package computerinfo

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"github.com/ryanehamil/computerinfo/src/utils"
)

// Canvas is the main object for the window
func canvasName() fyne.CanvasObject {

	computername := utils.GetComputerName()

	label := canvas.NewText(computername, color.White)
	label.TextSize = 100
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Alignment = fyne.TextAlignCenter

	return label
}

func Container() fyne.CanvasObject {
	// Create a layout
	layout := fyneLayout.NewCenterLayout()

	// Create a new container
	cont := container.New(layout, canvasName())
	cont.Resize(fyne.NewSize(200, 200))

	return cont
}
