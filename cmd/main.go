package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	index int
}

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

var pages = []string{"store", "cart"}

func newModel() model {
	return model{index: 0}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
        case "tab":
            m.index = (m.index + 1) % len(pages)
		}
	}
	return m, nil
}

func (m model) View() string {
    page := pages[m.index]
    switch page {
    case "store":
        return lipgloss.JoinVertical(
            0,
            lipgloss.JoinHorizontal(
                0,
                activeItemStyle.Render("* Store"),
                itemStyle.Render("  Cart")),

            storeStyle.Render("I AM STORE"),
        )
    case "cart":
        return lipgloss.JoinVertical(
            0,
            lipgloss.JoinHorizontal(
                0,
                itemStyle.Render("  Store"),
                activeItemStyle.Render("* Cart")),

            storeStyle.Render("I AM Cart"),
        )
    }

    panic("This shoul*d never happen")
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
