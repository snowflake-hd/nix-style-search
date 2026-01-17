package main

import (
	"nix-style-search/internal/cmd"
	"nix-style-search/internal/flags"

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
