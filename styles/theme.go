package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	// Name
	Name string

	// Background
	Background lipgloss.Color

	// Background focused
	BackgroundOver lipgloss.Color

	// Sub text
	SubText lipgloss.Color

	// Accent
	Accent lipgloss.Color

	// Foreground
	Foreground lipgloss.Color

	// Error
	Error lipgloss.Color
}

func InitTheme() Theme {
	return Theme{
		Name:           "Catppuccin",
		Background:     "#181825",
		BackgroundOver: "#7f849c",
		SubText:        "#6c7086",
		Accent:         "#b4befe",
		Foreground:     "#cdd6f4",
		Error:          "#f38ba8",
	}
}
