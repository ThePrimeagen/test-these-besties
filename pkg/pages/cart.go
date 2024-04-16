package pages

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type CartPage struct {
}

func NewCartPage(m Model) CartPage {
    return CartPage{ }
}

func (c *CartPage) Title() string { return "Cart" }

func (c *CartPage) Render(m *Model) string {
    totalWidgetCount := m.theme.
        ActiveTitleForeground().
        Bold(true).
        PaddingBottom(1).
        BorderBottom(true).
        Render(fmt.Sprintf("Items: %d", m.cart.totalItems))

    widgetOrders := make([]string, 0)
    widgetOrders = append(widgetOrders, totalWidgetCount)

    for _, widgetInfo := range m.cart.widgets {
        widgetCount := m.theme.
            NormalForeground().
            Bold(true).
            AlignHorizontal(lipgloss.Right).
            Render(fmt.Sprintf("Count: %d", m.cart.totalItems))

        title := m.theme.TitleForeground().
            Render(widgetInfo.widget.Name)

        desc := m.theme.DescForeground().
            Render(widgetInfo.widget.Description)

        m.renderer.NewStyle()

        widgetOrders = append(widgetOrders,
            lipgloss.JoinHorizontal(
                0,
                lipgloss.JoinVertical(
                    0,
                    title,
                    desc,
                ),
                widgetCount,
            ))
    }

    return lipgloss.JoinVertical(0, widgetOrders...)
}
