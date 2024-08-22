package components

import (
	"strings"

	tlockstyles "github.com/Achno/gocheat/styles"

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
		tlockstyles.Styles.Title.Render(title),
		tlockstyles.Styles.Base.Render(tlockstyles.Styles.Title.Render(suffix)),
		tlockstyles.Styles.Base, tlockstyles.Styles.ListItemActive,
	)
}

// List item inactive block
func ListItemInactive(width int, title, suffix string) string {
	return listItemImpl(
		width,
		tlockstyles.Styles.SubText.Render(title),
		tlockstyles.Styles.SubText.Render(suffix),
		tlockstyles.Styles.SubText, tlockstyles.Styles.ListItemInactive,
	)
}

func Paginator(listview list.Model) string {
	// Total pages
	totalPages := listview.Paginator.TotalPages

	// Paginator items
	paginatorItems := make([]string, totalPages)

	// Add paginator dots
	for index := 0; index < totalPages; index++ {
		renderer := tlockstyles.Styles.SubText.Copy().Bold(true).Render

		if index == listview.Paginator.Page {
			renderer = tlockstyles.Styles.Title.Render
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
	listview.Styles.Title = tlockstyles.Styles.Base

	// change Filter and cursor foregorund colors
	listview.FilterInput.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")).Bold(true)
	listview.FilterInput.Cursor.Style = tlockstyles.Styles.Title

	return listview
}
