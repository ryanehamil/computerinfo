package computerinfo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"github.com/ryanehamil/computerinfo/src/utils"
)

// obj_ComputerName returns a canvas object containing the computer name text
func obj_ComputerName() fyne.CanvasObject {
	// Get the computer name from the utils package
	computername := utils.GetComputerName()

	// Create a new text label with the computer name
	object := utils.NewTextButton(computername)
	object.TextSize = 100
	object.TextStyle = fyne.TextStyle{Bold: true}
	object.Alignment = fyne.TextAlignCenter

	// Return the label as a canvas object
	return object
}

// obj_IPAddress returns a slice of canvas objects containing the IP addresses of the computer
func obj_IPAddress() []fyne.CanvasObject {
	objects := []fyne.CanvasObject{}

	// Get the IP addresses of the computer
	ipaddress := utils.GetIPAddress()

	// Get only the first IP address
	ipaddress = ipaddress[:1]

	// Loop over the IP addresses and create a label for each one
	for _, ip := range ipaddress {
		// Create a new text label with the IP address
		object := utils.NewTextButton(ip)
		// object := canvas.NewText(ip, theme.Color("foreground"))
		object.TextSize = 20
		object.TextStyle = fyne.TextStyle{Bold: true}
		object.Alignment = fyne.TextAlignCenter
		// object.Color = theme.Color("foreground")

		// Append the label to the slice of labels
		objects = append(objects, object)

	}

	// Return the slice of labels as canvas objects
	return objects
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
