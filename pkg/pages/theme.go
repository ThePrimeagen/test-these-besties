package pages

import "github.com/charmbracelet/lipgloss"

type Theme interface {
	DescForeground() lipgloss.Style
	TitleForeground() lipgloss.Style
    NormalForeground() lipgloss.Style
}

type BasicTheme struct {
    desc lipgloss.Style
    title lipgloss.Style
    normal lipgloss.Style
}

func (b *BasicTheme) DescForeground() lipgloss.Style {
    return b.desc
}

func (b *BasicTheme) TitleForeground() lipgloss.Style {
    return b.title
}

func (b *BasicTheme) NormalForeground() lipgloss.Style {
    return b.normal
}

func GetTheme(m Model) Theme {
    return &BasicTheme{
        desc: m.renderer.NewStyle(),
        title: m.renderer.NewStyle(),
        normal: m.renderer.NewStyle(),
    }
}
