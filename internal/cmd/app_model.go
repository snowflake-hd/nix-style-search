package cmd

import (
	"nix-style-search/internal/cmd/index"
	"nix-style-search/internal/cmd/search"

	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	screen   tea.Model
	selected string
}

func NewAppModel() *AppModel {
	return &AppModel{screen: &index.Screen{}}
}

func (m *AppModel) Init() tea.Cmd {
	return m.screen.Init()
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case index.SelectedMsg:
		m.selected = msg.Index
		m.screen = search.NewSearchScreen(m.selected)
		return m, m.screen.Init()
	}

	next, cmdr := m.screen.Update(msg)
	m.screen = next
	return m, cmdr
}

func (m *AppModel) View() string {
	return m.screen.View()
}
