package main

//TODO clean stuff up,
//TODO create json file for configs for themes and keybindings
import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Achno/gocheat/internal/ui"
	tlockstyles "github.com/Achno/gocheat/styles"
)

func main() {

	// create the Screen where you view items
	model := ui.InitSelectItemScreen()

	// Initialize all lipgloss styles based on the theme and accessed by 'Styles' variable
	tlockstyles.InitializeStyles(tlockstyles.InitTheme())

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
