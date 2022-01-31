package computerinfo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/ryanehamil/computerinfo/src/utils"
)

// Canvas is the main object for the window
func obj_ComputerName() fyne.CanvasObject {

	computername := utils.GetComputerName()

	label := canvas.NewText(computername, theme.TextColor())
	label.TextSize = 100
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Alignment = fyne.TextAlignCenter

	return label
}
func obj_IPAddress() []fyne.CanvasObject {
	labels := []fyne.CanvasObject{}

	ipaddress := utils.GetIPAddress()

	for _, ip := range ipaddress {
		label := canvas.NewText(ip, theme.TextColor())
		label.TextSize = 20
		label.TextStyle = fyne.TextStyle{Bold: true}
		label.Alignment = fyne.TextAlignCenter
		labels = append(labels, label)
	}

	return labels
}

func allObjects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		obj_ComputerName(),
		container.NewVBox(
			obj_IPAddress()...,
		),
	}
}

func Container() fyne.CanvasObject {
	// Create a center layout
	// layout := fyneLayout.NewCenterLayout()

	// create a vertical layout
	verticalLayout := fyneLayout.NewVBoxLayout()

	// Create a new container
	cont := container.New(verticalLayout, allObjects()...)
	cont.Resize(fyne.NewSize(200, 200))

	return cont
}
