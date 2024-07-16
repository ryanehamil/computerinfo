package computerinfo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ryanehamil/computerinfo/src/utils"
)

// CustomTappableLabel is a label that supports tapping
type CustomTappableLabel struct {
	widget.Label
}

// NewCustomTappableLabel creates a new tappable label
func NewCustomTappableLabel(text string) *CustomTappableLabel {
	label := &CustomTappableLabel{
		Label: widget.Label{
			Text:      text,
			TextStyle: fyne.TextStyle{Bold: true},
		},
	}
	label.ExtendBaseWidget(label)
	return label
}

// Tapped is called when the label is tapped
func (l *CustomTappableLabel) Tapped(_ *fyne.PointEvent) {
	fyne.CurrentApp().Driver().AllWindows()[0].Clipboard().SetContent(l.Text)
}

// TappedSecondary is called when the label is right-clicked (not used here)
func (l *CustomTappableLabel) TappedSecondary(_ *fyne.PointEvent) {}

func obj_ComputerName() fyne.CanvasObject {
	computername := utils.GetComputerName()
	label := NewCustomTappableLabel(computername)
	label.Alignment = fyne.TextAlignCenter
	// label.TextSize = 100
	return label
}

func obj_IPAddress() []fyne.CanvasObject {
	ipAddresses := utils.GetIPAddress()
	labels := []fyne.CanvasObject{}

	for _, ip := range ipAddresses {
		container.NewHBox()
		label := NewCustomTappableLabel(ip)
		label.Alignment = fyne.TextAlignCenter
		// label.TextSize = 20
		labels = append(labels, label)
	}

	return labels
}

func allObjects() []fyne.CanvasObject {
	return append([]fyne.CanvasObject{
		obj_ComputerName(),
	}, obj_IPAddress()...)
}

func CanvasObject() fyne.CanvasObject {
	return container.NewVBox(allObjects()...)
}
