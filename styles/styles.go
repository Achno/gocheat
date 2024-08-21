package styles

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	// "github.com/eklairs/tlock/tlock-internal/context"
)

// Help
var Help help.Model

// Instance of styles
// Must be initialized on program's start
var Styles GCheatStyle

// Themes used all over tlock
type GCheatStyle struct {
	// Base
	Base lipgloss.Style

	// Title
	Title lipgloss.Style

	// Sub text
	SubText lipgloss.Style

	// Style for items over SubAltBg
	SubAltBg lipgloss.Style

	// Style for active list items
	ListItemActive lipgloss.Style

	// Style for inactive list item
	ListItemInactive lipgloss.Style

	// Style for input
	Input lipgloss.Style

	// Error
	Error lipgloss.Style

	// Style for placeholder
	Placeholder lipgloss.Style

	// Folder active list item
	FolderItemActive lipgloss.Style

	// Folder inactive list item
	FolderItemInactive lipgloss.Style

	// Tilte Bar
	AccentBgItem lipgloss.Style

	// Mock screen
	MockScreen lipgloss.Style

	// Background Over
	BackgroundOver lipgloss.Style

	// Sub text item
	SubTextItem lipgloss.Style

	// Overlay item
	OverlayItem lipgloss.Style

	// Time left for inactive cards
	TimeLeftInactive lipgloss.Style
}

// Initializes the styles
func InitializeStyles(theme Theme) {
	// Base
	base := lipgloss.NewStyle().Foreground(theme.Foreground)

	// Base for padded items
	paddedItem := with(base).Padding(1, 3)

	// Initialize styles
	Styles = GCheatStyle{
		Base:        with(base),
		Title:       with(base).Foreground(theme.Accent).Bold(true),
		SubText:     with(base).Foreground(theme.SubText),
		SubTextItem: with(base).Foreground(theme.SubText).Padding(0, 1),
		OverlayItem: with(base).Foreground(theme.SubText).Background(theme.BackgroundOver).Padding(0, 1),
		SubAltBg:    with(base).Background(theme.BackgroundOver).Padding(0, 1),
		Error:       with(base).Foreground(theme.Error).Bold(true),
		Input:       with(paddedItem).Width(65).Background(theme.BackgroundOver),
		Placeholder: with(base).Background(theme.BackgroundOver).Foreground(theme.SubText),
		// ListItemActive:     with(paddedItem).Background(theme.BackgroundOver),
		ListItemActive:     with(paddedItem).Border(lipgloss.RoundedBorder()),
		AccentBgItem:       with(base).Bold(true).Padding(0, 1).Background(theme.Accent).Foreground(theme.Background),
		BackgroundOver:     with(base).Background(theme.BackgroundOver),
		FolderItemInactive: with(paddedItem),
		MockScreen:         with(base).Background(theme.BackgroundOver).Align(lipgloss.Center, lipgloss.Center).Width(27).Height(9),
		FolderItemActive: with(paddedItem).
			Padding(1, 2).
			Background(theme.BackgroundOver).
			Border(lipgloss.OuterHalfBlockBorder(), false, false, false, true).
			BorderBackground(theme.BackgroundOver).
			BorderForeground(theme.Accent),
		ListItemInactive: with(paddedItem),
		TimeLeftInactive: with(base).Foreground(theme.BackgroundOver),
	}

	// Initialize help
	Help = help.New()

	// Comply help menu styles to themes
	Help.Styles.ShortKey = Styles.Title
	Help.Styles.ShortDesc = Styles.SubText
	Help.Styles.ShortSeparator = Styles.SubText
}

// Utility to copy a style
func with(style lipgloss.Style) lipgloss.Style {
	return style.Copy()
}

// Quick utilities

// Renders the text with `Title` style
func Title(text string) string {
	return Styles.Title.Render(text)
}

// Renders the text with `SubText` style
func Dimmed(text string) string {
	return Styles.SubText.Render(text)
}

// Renders the help menu
func HelpView(help help.KeyMap) string {
	return Help.View(help)
}
