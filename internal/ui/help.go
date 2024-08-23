package ui

import (
	"os"
	"strings"

	tlockstyles "github.com/Achno/gocheat/styles"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var helpAsciiArt = `
█ █ █▀▀ █   █▀█
█▀█ ██▄ █▄▄ █▀▀`

// Help key specification which holds a key and its description
type HelpKeyBindingSpec struct {
	// Key
	Key string
	// Description
	Desc string
}

// Holds the arrays for the keys and descriptions
type helpKeyBindings struct {
	// List
	List []HelpKeyBindingSpec

	// Input form
	InputForm []HelpKeyBindingSpec
}

// Builds the help menu for the given set of key bindings and the tile
func BuildHelpItem(title string, keys []HelpKeyBindingSpec) string {
	items := make([]string, 0)

	// Add title ex. List
	items = append(items, tlockstyles.Styles.Title.Render(title), "")

	// Add keys ex. Filter list '/'
	for _, key := range keys {
		ui := lipgloss.JoinHorizontal(
			lipgloss.Center,
			tlockstyles.Dimmed(key.Desc),
			strings.Repeat(" ", 65-len(key.Desc)-len(key.Key)),
			tlockstyles.Title(key.Key),
		)

		items = append(items, ui, "")
	}

	// Return
	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// Builds the help menu Ascii logo text and help sections with BuildHelpItem()
func BuildHelpMenu() string {

	var helpKeys = helpKeyBindings{
		List: []HelpKeyBindingSpec{
			{
				Key:  "↑/k",
				Desc: "Move up",
			},
			{
				Key:  "↓/j",
				Desc: "Move down",
			},
			{
				Key:  "→/l",
				Desc: "Move right (pagination)",
			},
			{
				Key:  "←/h",
				Desc: "Move left (pagination)",
			},
			{
				Key:  "/",
				Desc: "Filter the list",
			},
			{
				Key:  "ctrl+f",
				Desc: "Toggle filter list by tag",
			},
			{
				Key:  "ctrl+j",
				Desc: "Add an item to the list",
			},
			{
				Key:  "ctrl+x",
				Desc: "Delete focused item from the list",
			},
			{
				Key:  "esc",
				Desc: "Exit filter state",
			},
			{
				Key:  "ctrl+c",
				Desc: "Exit the app",
			},
		},
		InputForm: []HelpKeyBindingSpec{
			{
				Key:  "↓",
				Desc: "Move down",
			},
			{
				Key:  "↑",
				Desc: "Move up",
			},
			{
				Key:  "Enter",
				Desc: "Create an entry to the list",
			},
			{
				Key:  "esc",
				Desc: "Go back to the screen with the list",
			},
			{
				Key:  "ctrl+c",
				Desc: "Exit the app",
			},
		},
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		tlockstyles.Styles.Title.Render(helpAsciiArt), "",
		tlockstyles.Styles.SubText.Render("Keybindings to move around (esc to go back)"), "",
		BuildHelpItem("List", helpKeys.List),
		BuildHelpItem("Input Form", helpKeys.InputForm),
	)
}

// Help screen (model)
// Impliments tea.Model interface with : Init, Update , View
type HelpScreen struct {
	viewport viewport.Model
}

// Function to init a new instance of the Help screen
func InitializeHelpScreen() HelpScreen {
	_, height, _ := term.GetSize(int(os.Stdout.Fd()))

	viewport := viewport.New(65, height)
	viewport.SetContent(BuildHelpMenu())

	return HelpScreen{
		viewport: viewport,
	}
}

// Init
func (screen HelpScreen) Init() tea.Cmd {
	return screen.viewport.Init()
}

// Update
func (screen HelpScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgType := msg.(type) {
	case tea.KeyMsg:
		switch msgType.String() {
		case "esc":
			ItemScreen := InitSelectItemScreen()
			return ItemScreen, nil

		case "ctrl+c":
			return screen, tea.Quit
		}
	case tea.WindowSizeMsg:
		screen.viewport.Height = msgType.Height
	}

	// Update viewport
	screen.viewport, _ = screen.viewport.Update(msg)

	return screen, nil
}

// View
func (screen HelpScreen) View() string {
	// get the width and height of the terminal
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))

	// place the help screen viewport in the middle of the screen
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, screen.viewport.View())

}
