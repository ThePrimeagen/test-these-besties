package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"test.terminal.shop/pkg/shop"
)

// type currentPage int
//
// const (
// 	Cart currentPage = iota
// 	Widget
// )

type Model struct {
	currentPage int
	pages       []Page

	renderer *lipgloss.Renderer
	theme    Theme
}

func NewModel() *Model {
	widgets := shop.GetWidgets()

	renderer := lipgloss.DefaultRenderer()

	return &Model{
		renderer:    renderer,
		currentPage: 0,
		theme:       GetTheme(renderer),
		pages: []Page{
			&CartPage{},
			&WidgetPage{
				widget: &widgets[0],
			},
		},
	}
}

type Page interface {
	Title() string
	Render(m *Model) string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "tab":
			m.currentPage = (m.currentPage + 1) % len(m.pages)
		}
	}
	return m, nil
}

func (m Model) View() string {
	page := m.pages[m.currentPage]

	titleStyle := m.renderer.NewStyle().Border(lipgloss.BlockBorder())

	return fmt.Sprintf("%s\n%s", titleStyle.Render(page.Title()), page.Render(&m))
}
