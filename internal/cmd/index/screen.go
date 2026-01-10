package index

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Screen struct {
	Indices []string
	Cursor  int
	Err     string
}

func (i Screen) Init() tea.Cmd {
	return nil
}

func (i Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(i.Indices) == 0 && i.Err == "" {
		i = i.loadIndices()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			i.moveCursor(-1)
		case "down":
			i.moveCursor(1)
		case "enter":
			if len(i.Indices) == 0 {
				return i, nil
			}
			selected := i.Indices[i.Cursor]
			return i, func() tea.Msg { return SelectedMsg{Index: selected} }
		case "esc", "ctrl+c":
			return i, tea.Quit
		}
	}

	return i, nil
}
