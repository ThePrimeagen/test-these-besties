package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"test.terminal.shop/pkg/shop"
)

type WidgetPage struct {
	widget *shop.Widget
}

func (w *WidgetPage) Title() string { return "Widget" }

func (w *WidgetPage) Render(m *Model) string {
	titleStyle := m.renderer.NewStyle().
		Bold(true).Foreground(lipgloss.Color("#b294bb")).Underline(true).AlignHorizontal(lipgloss.Center).
		Width(m.Width)

	artHeight := len(strings.Split(w.widget.Art, "\n"))
	descriptionStyle := m.renderer.NewStyle().
		Margin(0, 2).Padding(0, 2).Height(artHeight)

	row := lipgloss.JoinHorizontal(lipgloss.Left, w.widget.Art, descriptionStyle.Render(w.widget.Description))

	return fmt.Sprintf(`%s
%s`, titleStyle.Render(w.widget.Name), row)
}
