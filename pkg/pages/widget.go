package pages

type WidgetPage struct{}

func (w *WidgetPage) Render(m *Model) string { return "WIDGET" }
