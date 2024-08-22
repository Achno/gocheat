package main

//TODO clean stuff up,
//TODO modify json file for themes
import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Achno/gocheat/config"
	"github.com/Achno/gocheat/internal/ui"
	tlockstyles "github.com/Achno/gocheat/styles"
)

func main() {

	// Make sure the config file is created
	config.Init()

	// Make sure that the items are read from the config file
	ui.Init()

	// create the Screen where you view items
	model := ui.InitSelectItemScreen()

	// Initialize all lipgloss styles based on the theme and accessed by 'Styles' variable
	tlockstyles.InitializeStyles(tlockstyles.InitTheme())

	//! fmt.Printf("CONFIGGGGGGGGG: %s", config.GoCheatOptions.Items)

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
