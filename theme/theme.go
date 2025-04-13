package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 30, G: 144, B: 255, A: 255} // Синий фон (DodgerBlue)
	case theme.ColorNameButton:
		return color.NRGBA{R: 70, G: 130, B: 180, A: 255} // Более темный синий для кнопок (SteelBlue)
	case theme.ColorNameForeground:
		return color.White // Белый цвет текста для контраста
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255} // Светло-серый для placeholder
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0, G: 191, B: 255, A: 255} // Ярко-синий для акцентов (DeepSkyBlue)
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
