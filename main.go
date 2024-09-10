package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Achno/gocheat/config"
	"github.com/Achno/gocheat/internal/ui"
	cheatstyles "github.com/Achno/gocheat/styles"
)

func main() {

	// Make sure the config file is created
	config.Init()

	// Make sure that the items are read from the config file
	ui.InitItems()

	// create the Screen where you view items
	model := ui.InitItemScreen()

	// Initialize all lipgloss styles based on the theme and accessed by 'Styles' variable
	cheatstyles.InitializeStyles(cheatstyles.InitTheme())

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
