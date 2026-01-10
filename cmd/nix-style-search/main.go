package main

import (
	"NixStyleSearch/internal/cmd"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(cmd.NewAppModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
