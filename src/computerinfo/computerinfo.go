package computerinfo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"github.com/ryanehamil/computerinfo/src/utils"
)

// obj_ComputerName returns a canvas object containing the computer name text
func obj_ComputerName() fyne.CanvasObject {
	computername := utils.GetComputerName()

	object := utils.NewClickableText(computername)
	object.Text.TextSize = 100

	return object
}

// obj_IPAddress returns a slice of canvas objects containing the IP addresses of the computer
func obj_IPAddress() []fyne.CanvasObject {
	// ipaddress := utils.GetIPAddress()
	ipaddress := utils.GetIPAddressObject()

	InterfaceObjects := []fyne.CanvasObject{}

	for _, ip := range ipaddress {
		interfaceName := ip["name"]
		ip := ip["ip"]
		object := container.NewHBox()

		ClickableName := utils.NewClickableText(interfaceName)
		ClickableName.Text.TextSize = 20

		object.Add(ClickableName)

		ClickableIP := utils.NewClickableText(ip)
		ClickableIP.Text.TextSize = 20

		object.Add(ClickableIP)

		wrapper := container.NewCenter(object)

		InterfaceObjects = append(InterfaceObjects, wrapper)
	}

	return InterfaceObjects
}

// allObjects returns a slice of canvas objects containing both the computer name and IP addresses
func allObjects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		obj_ComputerName(),
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
	// cont.Resize(fyne.NewSize(200, 200))

	// Return the container
	return cont
}
