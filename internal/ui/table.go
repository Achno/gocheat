package ui

import (
	"os"
	"strconv"

	"github.com/Achno/gocheat/config"
	cheatstyles "github.com/Achno/gocheat/styles"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

var tagsAscii = `
▀█▀ ▄▀█ █▀▀ █▀
 █  █▀█ █▄█ ▄█
`

type TableScreen struct {
	rows     [][]string
	table    table.Table
	viewport viewport.Model
}

func InitTableScreen() TableScreen {
	_, height, _ := term.GetSize(int(os.Stdout.Fd()))

	viewport := viewport.New(40, height)
	viewport.SetContent(buildTableScreen())

	return TableScreen{
		viewport: viewport,
	}
}

func (screen TableScreen) Init() tea.Cmd {
	return screen.viewport.Init()
}

func (screen TableScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgType := msg.(type) {
	case tea.KeyMsg:
		switch msgType.String() {
		case "esc":
			ItemScreen := InitItemScreen()
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

func (screen TableScreen) View() string {
	// get the width and height of the terminal
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))

	// place the help screen viewport in the middle of the screen
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, screen.viewport.View())

}

func buildTableScreen() string {
	table := initTable()

	return lipgloss.JoinVertical(
		lipgloss.Center,
		cheatstyles.Styles.Title.Render(tagsAscii), "",
		cheatstyles.Styles.SubText.Render("View all your tags (esc to go back)"), "",
		table.Render(),
	)
}

func getRows(items []list.Item) [][]string {
	tagCount := make(map[string]int)
	var rows [][]string

	if len(items) <= 0 {
		return [][]string{
			{"No tags", "null"},
		}
	}

	// Count the occurrences of each tag
	for _, item := range items {
		itm, _ := item.(Item)
		tagCount[itm.Tag]++
	}

	// Add each tag and its count to the rows
	for tag, count := range tagCount {
		rows = append(rows, []string{tag, strconv.Itoa(count)})
	}

	return rows

}

func initTable() table.Table {
	table := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(cheatstyles.Styles.Title).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return cheatstyles.Styles.Title
			default:
				return cheatstyles.Styles.SubText
			}
		}).
		Headers("TAG", "NUMBER OF ENTRIES").
		Rows(getRows(ConvertItemWrappers(config.GoCheatOptions.Items))...)

	return *table
}
