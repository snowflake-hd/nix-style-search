package search

import (
	"NixStyleSearch/internal"
	"NixStyleSearch/internal/search"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	pageSize      = internal.PAGE_SIZE
	debounceDelay = 150 * time.Millisecond
)

type searchTriggerMsg struct{ seq int }

type Screen struct {
	Index       string
	Ti          textinput.Model
	Pkgs        []string
	Sources     []search.PackageSource
	Cursor      int
	End         int
	Err         string
	ShowDetails bool

	query       string
	endReached  bool
	start       int
	next        int
	debounceSeq int
}

func NewSearchScreen(index string) *Screen {
	ti := textinput.New()

	ti.Placeholder = "Type your search..."
	ti.CharLimit = 256
	ti.Focus()
	ti.Width = 50

	return &Screen{
		Index: index,
		Ti:    ti,
	}
}

func (s *Screen) Init() tea.Cmd {
	return nil
}

func (s *Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	downPressedAtEnd := false
	upPressedAtStart := false
	var cmds []tea.Cmd

	if m, ok := msg.(searchTriggerMsg); ok {
		if m.seq == s.debounceSeq && s.query != "" && len(s.Sources) == 0 {
			s.fetchForward(false)
		}
		if s.Cursor > len(s.Pkgs)-1 {
			s.Cursor = len(s.Pkgs) - 1
		}
		if s.Cursor < 0 {
			s.Cursor = 0
		}
		return s, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			s.showSelection()
			return s, tea.Quit
		case "up":
			if s.Cursor > 0 {
				s.Cursor--
			} else if s.Cursor == 0 {
				upPressedAtStart = true
			}
		case "down":
			if s.Cursor < len(s.Pkgs)-1 {
				s.Cursor++
			} else if s.Cursor == len(s.Pkgs)-1 {
				downPressedAtEnd = true
			}
		case "tab":
			s.ShowDetails = !s.ShowDetails
		case "enter":
			s.showSelection()
			return s, tea.Quit
		}

	}

	ti, cmd := s.Ti.Update(msg)
	s.Ti = ti
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	newQuery := strings.TrimSpace(s.Ti.Value())
	if newQuery != s.query {
		s.query = newQuery
		s.start = 0
		s.next = 0
		s.endReached = false
		s.Sources = nil
		s.Pkgs = nil
		s.Cursor = 0

		s.debounceSeq++
		seq := s.debounceSeq
		cmds = append(cmds, tea.Tick(debounceDelay, func(time.Time) tea.Msg {
			return searchTriggerMsg{seq}
		}))
	}

	fetchDown := s.query != "" && downPressedAtEnd && !s.endReached
	fetchUp := s.query != "" && upPressedAtStart && s.start > 0

	if s.query != "" {
		if fetchDown {
			s.fetchForward(true)
			s.Cursor = 0
		}
		if fetchUp {
			s.fetchBackward()
			s.Cursor = len(s.Pkgs) - 1
		}

		if s.Cursor > len(s.Pkgs)-1 {
			s.Cursor = len(s.Pkgs) - 1
		}
		if s.Cursor < 0 {
			s.Cursor = 0
		}
	}

	return s, tea.Batch(cmds...)
}
