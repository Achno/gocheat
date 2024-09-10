package ui

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/Achno/gocheat/config"
	"github.com/Achno/gocheat/internal/components"
	cheatstyles "github.com/Achno/gocheat/styles"
)

// gloval var which holds the item of the list
var items []list.Item

// Init items based on the config.json config file
func InitItems() {

	items = ConvertItemWrappers(config.GoCheatOptions.Items)
}

// Controls the filtering mode
var FilterbyTag = false

var selectUserAscii = `

█▀▀ █ █ █▀▀ ▄▀█ ▀█▀ █▀ █ █ █▀▀ █▀▀ ▀█▀
█▄▄ █▀█ ██▄ █▀█  █  ▄█ █▀█ ██▄ ██▄  █
`

// impliments list.Item interface : FilterValue()
type Item struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
}

// The value the fuzzy filter , filters by
func (item Item) FilterValue() string {
	if FilterbyTag {
		return item.Tag
	} else {
		return item.Title
	}
}

type ItemDelegate struct{}

func (delegate ItemDelegate) Height() int { return 3 }

// Spacing
func (delegate ItemDelegate) Spacing() int { return 0 }

// Update
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

// Render
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(Item)

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
type itemKeyMap struct {
	Up         key.Binding
	Down       key.Binding
	Filter     key.Binding
	Back       key.Binding
	FilterMode key.Binding
}

func (k itemKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Filter, k.Back, k.FilterMode}
}

func (k itemKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up},
		{k.Down},
		{k.Filter},
		{k.Back},
		{k.FilterMode},
	}
}

// Initialize keybinds
var selectItemKeys = itemKeyMap{
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
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "help"),
	),
}

// Model for the select screen
//
// Impliments the tea.Model interface : Init() Update() View()
type ItemScreen struct {

	// the List ui model
	listview list.Model
}

func InitItemScreen() ItemScreen {
	return ItemScreen{
		listview: components.ListViewSimple(items, ItemDelegate{}, 65, min(12, len(items)*3)),
	}
}

func (screen ItemScreen) Init() tea.Cmd {
	return nil
}

func (screen ItemScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			statusCmd := screen.listview.NewStatusMessage(cheatstyles.Styles.StatusMsg.Render(fmt.Sprintf("Filter by tag: %t", FilterbyTag)))
			return screen, statusCmd

		case "ctrl+h":
			helpModel := InitializeHelpScreen()
			return helpModel, cmd

		case "ctrl+j":
			InputModel := InitInputFormScreen()
			return InputModel, cmd

		case "ctrl+k":
			TableModel := InitTableScreen()
			return TableModel, cmd

		case "ctrl+x":
			return HandleRemovingItem(screen)
		}
	}
	// Update listview
	screen.listview, cmd = screen.listview.Update(msg)
	cmds = append(cmds, cmd)

	// Return
	return screen, tea.Batch(cmds...)
}

// View
func (screen ItemScreen) View() string {
	// Set height
	screen.listview.SetHeight(min(12, len(items)*3))

	// List of items to render
	items := []string{
		cheatstyles.Title(selectUserAscii), "",
		screen.listview.View(), "",
	}

	// Add paginator
	if screen.listview.Paginator.TotalPages > 1 {
		items = append(items, components.Paginator(screen.listview), "")
	}

	// Add help
	items = append(items, cheatstyles.HelpView(selectItemKeys))

	joinedItems := lipgloss.JoinVertical(
		lipgloss.Center,
		items...,
	)

	// get terminal dimensions
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	// Place list in the center
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, joinedItems)
}

// ConverItemWrappers converts []ItemWrapper to []list.Item
func ConvertItemWrappers(wrappers []config.ItemWrapper) []list.Item {
	var items []list.Item
	for _, wrapper := range wrappers {
		item := Item{
			Title: wrapper.Title,
			Tag:   wrapper.Tag,
		}
		items = append(items, item)
	}
	return items
}

// Convert []list.Item to []ItemWrapper
func ConvertListItemsToItemWrappers(items []list.Item) []config.ItemWrapper {
	var wrappers []config.ItemWrapper
	for _, item := range items {
		// Assert that item is of type Item
		selectedItem, ok := item.(Item)
		if !ok {
			// Handle the case where item is not of type Item
			continue
		}

		// Create ItemWrapper from Item
		wrapper := config.ItemWrapper{
			Title: selectedItem.Title,
			Tag:   selectedItem.Tag,
		}
		wrappers = append(wrappers, wrapper)
	}
	return wrappers
}

// Removes an ItemWrapper from the config.json and screen
func removeItemFromConfig(slice []config.ItemWrapper) error {

	config.GoCheatOptions.Items = slice

	err := config.UpdateConfig()

	if err != nil {
		return err
	}

	return nil
}

// Removes the selected item from the screen and from config.json
func HandleRemovingItem(screen ItemScreen) (tea.Model, tea.Cmd) {
	// check filter state
	filterState := screen.listview.FilterState()

	if filterState.String() == "filter applied" {

		// We cant get the index of the item with listview.Index() so remove manually
		itm := screen.listview.SelectedItem()
		i := itm.(Item)

		for index, v := range items {
			selectedItem, _ := v.(Item)
			if selectedItem.Title == i.Title {
				items = append(items[:index], items[index+1:]...)
			}

		}

		removeItemFromConfig(ConvertListItemsToItemWrappers(items))
		s := InitItemScreen()
		statusCmd := s.listview.NewStatusMessage(cheatstyles.Styles.StatusMsg.Render("Deleted item"))

		return s, statusCmd
	}

	//Else just delete the item from the screen & config updading only the list and not the screen
	index := screen.listview.Index()
	screen.listview.RemoveItem(index)
	removeItemFromConfig(ConvertListItemsToItemWrappers(items))
	statusCmd := screen.listview.NewStatusMessage(cheatstyles.Styles.StatusMsg.Render("Deleted item"))
	return screen, statusCmd
}
