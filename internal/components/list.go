package components

import (
	"strings"

	cheatstyles "github.com/Achno/gocheat/styles"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// === List item implimentation ==
func listItemImpl(width int, title, suffix string, spacerStyle, style lipgloss.Style) string {
	// Required space
	space_width := width - lipgloss.Width(title) - lipgloss.Width(suffix)

	// Join
	ui := lipgloss.JoinHorizontal(lipgloss.Center, title, spacerStyle.Render(strings.Repeat(" ", space_width)), suffix)

	// Return
	return style.Render(ui)
}

// List item active block
func ListItemActive(width int, title, suffix string) string {
	return listItemImpl(
		width,
		cheatstyles.Styles.Title.Render(title),
		cheatstyles.Styles.Base.Render(cheatstyles.Styles.Title.Render(suffix)),
		cheatstyles.Styles.Base, cheatstyles.Styles.ListItemActive,
	)
}

// List item inactive block
func ListItemInactive(width int, title, suffix string) string {
	return listItemImpl(
		width,
		cheatstyles.Styles.SubText.Render(title),
		cheatstyles.Styles.SubText.Render(suffix),
		cheatstyles.Styles.SubText, cheatstyles.Styles.ListItemInactive,
	)
}

func Paginator(listview list.Model) string {
	// Total pages
	totalPages := listview.Paginator.TotalPages

	// Paginator items
	paginatorItems := make([]string, totalPages)

	// render only 5 dots (pages) at maximum
	if totalPages > 5 {
		totalPages = 5
	}

	// Add paginator dots
	for index := 0; index < totalPages; index++ {
		renderer := cheatstyles.Styles.SubText.Copy().Bold(true).Render

		if index == listview.Paginator.Page {
			renderer = cheatstyles.Styles.Title.Render
		}

		paginatorItems = append(paginatorItems, renderer("•"))
	}

	// Add to ui
	return lipgloss.JoinHorizontal(lipgloss.Center, paginatorItems...)
}

// Builds a listview model devoid of every feature
func ListViewSimple(items []list.Item, delegate list.ItemDelegate, width, height int) list.Model {
	listview := list.New(items, delegate, width, height)

	listview.SetShowHelp(false)
	listview.SetShowTitle(true)
	listview.SetShowFilter(true)
	listview.SetShowStatusBar(false)
	listview.SetShowPagination(false)
	listview.DisableQuitKeybindings()
	listview.SetFilteringEnabled(true)
	listview.ShowPagination()

	// hide the title of the list , but allow it to exist to show notifications with listview.NewStatusMessage(...)
	listview.Title = ""
	listview.Styles.Title = cheatstyles.Styles.Base

	// change Filter and cursor foregorund colors
	listview.FilterInput.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")).Bold(true)
	listview.FilterInput.Cursor.Style = cheatstyles.Styles.Title

	return listview
}
