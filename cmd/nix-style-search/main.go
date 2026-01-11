package main

import (
	"NixStyleSearch/internal/cmd"
	"NixStyleSearch/internal/flags"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	shouldStart := flags.HandleFlags()

	if !shouldStart {
		return
	}

	p := tea.NewProgram(cmd.NewAppModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
