package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Achno/gocheat/config"
	"github.com/Achno/gocheat/internal/components"
	tlockstyles "github.com/Achno/gocheat/styles"
)

var items []list.Item

func Init() {

	items = ConvertSelectItemWrappers(config.GoCheatOptions.Items)

}

// var items = []list.Item{
// 	SelectedItem{Title: "Maximize Window : meta+up", Tag: "Kwin"},
// 	SelectedItem{Title: "Minimize Window : meta+m", Tag: "Kwin"},
// 	SelectedItem{Title: "Rofi : fn+end", Tag: "Rofi"},
// 	SelectedItem{Title: "Take a screenshot  : f2", Tag: "Flameshot"},
// 	SelectedItem{Title: "Open the menu : f1", Tag: "wlogout"},
// 	SelectedItem{Title: "cube : meta + w", Tag: "kwin"},
// 	SelectedItem{Title: "resize Window : alt+k", Tag: "Flameshot"},
// 	SelectedItem{Title: "Lock windows in place : ctrl+alt", Tag: "Kwin"},
// }

// Controls the filtering mode
var FilterbyTag = false

var selectUserAscii = `

█▀▀ █ █ █▀▀ ▄▀█ ▀█▀ █▀ █ █ █▀▀ █▀▀ ▀█▀
█▄▄ █▀█ ██▄ █▀█  █  ▄█ █▀█ ██▄ ██▄  █
`

// impliments list.Item interface : FilterValue()
type SelectedItem struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
}

// The value the fuzzy filter , filters by
func (item SelectedItem) FilterValue() string {
	if FilterbyTag {
		return item.Tag
	} else {
		return item.Title
	}
}

type SelectItemDelegate struct{}

func (delegate SelectItemDelegate) Height() int { return 3 }

// Spacing
func (delegate SelectItemDelegate) Spacing() int { return 0 }

// Update
func (d SelectItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

// Render
func (d SelectItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(SelectedItem)

	if !ok {
		return
	}

	// Decide the renderer based on focused index
	renderer := components.ListItemInactive
	if index == m.Index() {
		renderer = components.ListItemActive
	}

	// Render the title of the item and its tag
	fmt.Fprint(w, renderer(65, string(item.Title), string(item.Tag)))
}

// Explanation: Keybindings that the list listens on.
//
// Impliments: Keymap interface
type selectItemKeyMap struct {
	Up         key.Binding
	Down       key.Binding
	Filter     key.Binding
	Back       key.Binding
	FilterMode key.Binding
}

func (k selectItemKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Filter, k.Back, k.FilterMode}
}

func (k selectItemKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up},
		{k.Down},
		{k.Filter},
		{k.Back},
		{k.FilterMode},
	}
}

// Initialize keybin
var selectItemKeys = selectItemKeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Filter: key.NewBinding(
		key.WithKeys("filter", "/"),
		key.WithHelp("/", "filter"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "exit filtering"),
	),
	FilterMode: key.NewBinding(
		key.WithKeys("ctrl+f"),
		key.WithHelp("ctrl+f", "filter tag"),
	),
}

// Model for the select screen
//
// Impliments the tea.Model interface : Init() Update() View()
type SelectItemScreen struct {

	// the List ui model
	listview list.Model
}

func InitSelectItemScreen() SelectItemScreen {
	return SelectItemScreen{
		listview: components.ListViewSimple(items, SelectItemDelegate{}, 65, min(12, len(items)*3)),
	}
}

func (screen SelectItemScreen) Init() tea.Cmd {
	return nil
}

func (screen SelectItemScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// List of cmds to send
	cmds := make([]tea.Cmd, 0)

	// Handle key presses
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return screen, tea.Quit

		case "ctrl+f":
			FilterbyTag = !FilterbyTag
			// For a status msg to be shown the title of the list needs to be visible
			statusCmd := screen.listview.NewStatusMessage(tlockstyles.Styles.StatusMsg.Render(fmt.Sprintf("Filter by tag: %t", FilterbyTag)))
			return screen, statusCmd
		}
	}
	// Update listview
	screen.listview, cmd = screen.listview.Update(msg)
	cmds = append(cmds, cmd)

	// Return
	return screen, tea.Batch(cmds...)
}

// View
func (screen SelectItemScreen) View() string {
	// Set height
	screen.listview.SetHeight(min(12, len(items)*3))

	// List of items to render
	items := []string{
		tlockstyles.Title(selectUserAscii), "",
		// tlockstyles.Dimmed("Select a user to login as"), "",
		screen.listview.View(), "",
	}

	// Add paginator
	if screen.listview.Paginator.TotalPages > 1 {
		items = append(items, components.Paginator(screen.listview), "")
	}

	// Add help
	items = append(items, tlockstyles.HelpView(selectItemKeys))

	// Return
	joinedItems := lipgloss.JoinVertical(
		lipgloss.Center,
		items...,
	)

	// Place list in the center
	return lipgloss.Place(screen.listview.Width(), screen.listview.Height(), lipgloss.Center, lipgloss.Center, joinedItems)
}

// ConvertSelectItemWrappers converts []SelectItemWrapper to []list.Item
func ConvertSelectItemWrappers(wrappers []config.SelectItemWrapper) []list.Item {
	var items []list.Item
	for _, wrapper := range wrappers {
		item := SelectedItem{
			Title: wrapper.Title,
			Tag:   wrapper.Tag,
		}
		items = append(items, item)
	}
	return items
}
