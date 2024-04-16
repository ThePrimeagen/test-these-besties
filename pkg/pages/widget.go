package pages

import (
	"fmt"

	"test.terminal.shop/pkg/shop"
)

type WidgetPage struct {
	widget *shop.Widget
}

func (w *WidgetPage) Render(m *Model) string {
	return fmt.Sprintf("WIDGET: %s", w.widget.Name)
}
