package main

import (
	"NixStyleSearch/internal/cmd"
	"NixStyleSearch/internal/flags"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	shouldStart := flags.DefineFlags()

	if !shouldStart {
		return
	}

	p := tea.NewProgram(cmd.NewAppModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
