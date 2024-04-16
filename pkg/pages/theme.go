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
func GetTheme(renderer *lipgloss.Renderer) Theme {
	return &BasicTheme{
		desc:         renderer.NewStyle(),
		title:        renderer.NewStyle(),
		normal:       renderer.NewStyle(),
		activeDesc:   renderer.NewStyle(),
		activeTitle:  renderer.NewStyle(),
		activeNormal: renderer.NewStyle(),
	}
}
