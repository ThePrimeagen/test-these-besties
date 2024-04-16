package pages

import "github.com/charmbracelet/lipgloss"

type Theme interface {
	DescForeground() lipgloss.Style
	DescBackground() lipgloss.Style

	TitleForeground() lipgloss.Style
	TitleBackground() lipgloss.Style

    NormalForeground() lipgloss.Style
    NormalBackground() lipgloss.Style
}

func GetTheme(m Model) Theme {
    return nil
}
