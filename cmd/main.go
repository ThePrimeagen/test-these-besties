package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"test.terminal.shop/pkg/pages"
)

var pageNav = lipgloss.NewStyle().
	Bold(true).
	PaddingBottom(2).
    BorderBottom(true).
	Align(lipgloss.Left)

var itemStyle = lipgloss.NewStyle().
    Align(lipgloss.Left).
    Width(15)

var activeItemStyle = itemStyle.Copy().
    Foreground(lipgloss.Color("#123456"))

var storeStyle = lipgloss.NewStyle().
	PaddingBottom(2)

var cartStyle = lipgloss.NewStyle().
	PaddingBottom(2)


func main() {
	if _, err := tea.NewProgram(pages.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
