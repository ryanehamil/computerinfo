package computerinfo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/ryanehamil/computerinfo/src/utils"
)

// obj_ComputerName returns a canvas object containing the computer name text
func obj_ComputerName() fyne.CanvasObject {
	// Get the computer name from the utils package
	computername := utils.GetComputerName()

	// Create a new text label with the computer name
	label := canvas.NewText(computername, theme.TextColor())
	label.TextSize = 100
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Alignment = fyne.TextAlignCenter

	// Return the label as a canvas object
	return label
}

// obj_IPAddress returns a slice of canvas objects containing the IP addresses of the computer
func obj_IPAddress() []fyne.CanvasObject {
	labels := []fyne.CanvasObject{}

	// Get the IP addresses of the computer
	ipaddress := utils.GetIPAddress()

	// Loop over the IP addresses and create a label for each one
	for _, ip := range ipaddress {
		label := canvas.NewText(ip, theme.TextColor())
		label.TextSize = 20
		label.TextStyle = fyne.TextStyle{Bold: true}
		label.Alignment = fyne.TextAlignCenter
		labels = append(labels, label)
	}

	// Return the slice of labels as canvas objects
	return labels
}

// allObjects returns a slice of canvas objects containing both the computer name and IP addresses
func allObjects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		// Get the computer name canvas object
		obj_ComputerName(),
		// Create a VBox container containing the IP addresses
		container.NewVBox(
			obj_IPAddress()...,
		),
	}
}

// Container creates a container object for the window
func Container() fyne.CanvasObject {
	// Create a vertical layout
	verticalLayout := fyneLayout.NewVBoxLayout()

	// Create a new container using the vertical layout and the canvas objects
	cont := container.New(verticalLayout, allObjects()...)
	cont.Resize(fyne.NewSize(200, 200))

	// Return the container
	return cont
}
