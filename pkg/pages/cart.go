package pages

type CartPage struct {
}

func NewCartPage(m Model) CartPage {
    return CartPage{
    }
}

func (c *CartPage) Title() string { return "Cart" }

func (c *CartPage) Render(m *Model) string {
	return ""
}
