package pages

import "github.com/charmbracelet/lipgloss"

type CartStyles struct {
}

type CartPage struct {
    styles CartStyles
}

func NewCartPage(m Model) CartPage {
    return CartPage{
        styles: CartStyles{
            activeTitle: m.renderer.NewStyle()
        }
    }
}

func (c *CartPage) Title() string { return "Cart" }

func (c *CartPage) Render(m *Model) string {
	return ""
}
