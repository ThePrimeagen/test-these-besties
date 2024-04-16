package pages

import "github.com/charmbracelet/lipgloss"

type Theme interface {
	DescForeground() lipgloss.Style
	TitleForeground() lipgloss.Style
	NormalForeground() lipgloss.Style
	ActiveDescForeground() lipgloss.Style
	ActiveTitleForeground() lipgloss.Style
	ActiveNormalForeground() lipgloss.Style
}

type BasicTheme struct {
	desc         lipgloss.Style
	title        lipgloss.Style
	normal       lipgloss.Style
	activeDesc   lipgloss.Style
	activeTitle  lipgloss.Style
	activeNormal lipgloss.Style
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

func (b *BasicTheme) ActiveDescForeground() lipgloss.Style {
	return b.activeDesc
}

func (b *BasicTheme) ActiveTitleForeground() lipgloss.Style {
	return b.activeTitle
}

func (b *BasicTheme) ActiveNormalForeground() lipgloss.Style {
	return b.activeNormal
}
func GetTheme(m Model) Theme {
	return &BasicTheme{
		desc:         m.renderer.NewStyle(),
		title:        m.renderer.NewStyle(),
		normal:       m.renderer.NewStyle(),
		activeDesc:   m.renderer.NewStyle(),
		activeTitle:  m.renderer.NewStyle(),
		activeNormal: m.renderer.NewStyle(),
	}
}
