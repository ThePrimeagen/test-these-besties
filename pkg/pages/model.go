package pages

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
    currentPage int
    pages []Page
}

func NewModel() *Model {
    return &Model {
        currentPage: 0,
        pages: []Page{
            &CartPage{},
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
		}
	}
	return m, nil
}

func (m Model) View() string {
    panic("This shoul*d never happen")
}


