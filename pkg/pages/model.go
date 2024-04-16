package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"test.terminal.shop/pkg/shop"
)

type Model struct {
	currentPage int
	pages       []Page

	renderer *lipgloss.Renderer
}

func NewModel() *Model {
	widgets := shop.GetWidgets()

	return &Model{
		currentPage: 0,
		pages: []Page{
			&CartPage{},
			&WidgetPage{
				widget: &widgets[0],
			},
		},
	}
}

type Page interface {
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
	return fmt.Sprintf("KEKW: %d\n%s", m.currentPage, page.Render(&m))
}
