package utils

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Theme struct{}

var _ fyne.Theme = (*Theme)(nil)

// The color package has to be imported from "image/color".

func (m Theme) GetColor(name string) color.Color {
	variant := fyne.CurrentApp().Settings().ThemeVariant()
	switch variant {
	case theme.VariantDark:
		return color.RGBA{0, 0, 0, 255}
	case theme.VariantLight:
		return color.RGBA{255, 255, 255, 255}
	}
	return color.RGBA{0, 0, 0, 255}
}

func (m Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			// return color #eeeeee
			return color.RGBA{238, 238, 238, 255}
		}
		// return color #222222
		return color.RGBA{34, 34, 34, 255}
	}
	if name == theme.ColorNameForeground {
		if variant == theme.VariantLight {
			// return color #222222
			return color.RGBA{34, 34, 34, 255}
		}
		// return color #eeeeee
		return color.RGBA{238, 238, 238, 255}
	}
	if name == theme.ColorNameButton {
		if variant == theme.VariantLight {
			// return color #eeeeee
			return color.RGBA{238, 238, 238, 255}
		}
		// return color #222222
		return color.RGBA{34, 34, 34, 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m Theme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
