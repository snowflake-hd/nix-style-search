package search

import (
	"NixStyleSearch/internal/search"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestScreenUpdateCursorMovement(t *testing.T) {
	s := NewSearchScreen("idx")
	s.Pkgs = []string{"alpha", "beta"}
	s.Sources = []search.PackageSource{{PackageAttrName: "alpha"}, {PackageAttrName: "beta"}}

	s.Update(tea.KeyMsg{Type: tea.KeyDown})
	if s.Cursor != 1 {
		t.Fatalf("expected cursor to move down to 1, got %d", s.Cursor)
	}

	s.Update(tea.KeyMsg{Type: tea.KeyUp})
	if s.Cursor != 0 {
		t.Fatalf("expected cursor to move up to 0, got %d", s.Cursor)
	}

	s.Update(tea.KeyMsg{Type: tea.KeyUp})
	if s.Cursor != 0 {
		t.Fatalf("expected cursor to remain at 0, got %d", s.Cursor)
	}
}

func TestScreenUpdateNewQuerySchedulesTick(t *testing.T) {
	s := NewSearchScreen("idx")
	s.query = "old"
	s.Ti.SetValue("new")

	model, cmd := s.Update(nil)
	if model.(*Screen).query != "new" {
		t.Fatalf("expected query to update to new value, got %q", model.(*Screen).query)
	}
	if cmd == nil {
		t.Fatalf("expected debounce tick command")
	}

	msg := cmd()
	if _, ok := msg.(searchTriggerMsg); !ok {
		t.Fatalf("expected searchTriggerMsg, got %T", msg)
	}
}

func TestScreenUpdateDownTriggersFetch(t *testing.T) {
	original := searchPageFn
	defer func() { searchPageFn = original }()

	called := 0
	searchPageFn = func(index, query string, from, size int) ([]search.PackageSource, int, error) {
		called++
		return []search.PackageSource{{PackageAttrName: "beta"}}, 1, nil
	}

	s := NewSearchScreen("idx")
	s.query = "foo"
	s.Ti.SetValue("foo")
	s.Pkgs = []string{"alpha"}
	s.Sources = []search.PackageSource{{PackageAttrName: "alpha"}}

	s.Update(tea.KeyMsg{Type: tea.KeyDown})

	if called == 0 {
		t.Fatalf("expected fetchForward to run when pressing down at end")
	}
	if len(s.Pkgs) != 2 || s.Pkgs[1] != "beta" {
		t.Fatalf("expected new item appended with overlap, got %#v", s.Pkgs)
	}
}
