package pages

import "github.com/charmbracelet/lipgloss"

type Theme interface {
	DescForeground() lipgloss.Color
	DescBackground() lipgloss.Color

	TitleForeground() lipgloss.Color
	TitleBackground() lipgloss.Color

    NormalForeground() lipgloss.Color
    NormalBackground() lipgloss.Color
}

