package pages

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"test.terminal.shop/pkg/shop"
)

type WidgetPage struct {
	widget *shop.Widget
}

func (w *WidgetPage) Title() string { return "Widget" }

func (w *WidgetPage) Render(m *Model) string {
	titleStyle := m.renderer.NewStyle().Bold(true).Foreground(lipgloss.Color("#b294bb"))
	return fmt.Sprintf(`<headers go here...>
%s
WIDGET: %s`, titleStyle.Render(w.widget.Name), w.widget.Name)
}
