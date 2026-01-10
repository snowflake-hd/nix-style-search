package index

import (
	"errors"
	"testing"

	"NixStyleSearch/internal/search"
)

func TestCleanAliasFiltersHidden(t *testing.T) {
	in := []search.Alias{{Alias: ".hidden", Index: "idx"}, {Alias: "shown", Index: "visible"}, {Alias: "", Index: ""}}
	out := cleanAlias(&in)
	if len(*out) != 1 || (*out)[0].Alias != "shown" {
		t.Fatalf("expected only visible alias, got %#v", out)
	}
}

func TestLoadIndicesSuccess(t *testing.T) {
	original := fetchAliases
	defer func() { fetchAliases = original }()

	fetchAliases = func() ([]search.Alias, error) {
		return []search.Alias{{Alias: "one", Index: "idx"}}, nil
	}

	scr := Screen{}
	scr = scr.loadIndices()

	if len(scr.Indices) != 1 || scr.Indices[0] != "one" {
		t.Fatalf("expected indices populated, got %#v", scr.Indices)
	}
	if scr.Err != "" {
		t.Fatalf("did not expect error, got %s", scr.Err)
	}
}

func TestLoadIndicesError(t *testing.T) {
	original := fetchAliases
	defer func() { fetchAliases = original }()

	fetchAliases = func() ([]search.Alias, error) { return nil, errors.New("boom") }

	scr := Screen{}
	scr = scr.loadIndices()

	if scr.Err == "" {
		t.Fatalf("expected error message when fetch fails")
	}
}

func TestMoveCursorClamps(t *testing.T) {
	scr := Screen{Indices: []string{"a", "b", "c"}}
	scr.moveCursor(2)
	if scr.Cursor != 2 {
		t.Fatalf("expected cursor to move within bounds, got %d", scr.Cursor)
	}
	scr.moveCursor(1)
	if scr.Cursor != 2 {
		t.Fatalf("expected cursor to clamp at upper bound, got %d", scr.Cursor)
	}
	scr.moveCursor(-5)
	if scr.Cursor != 0 {
		t.Fatalf("expected cursor to clamp at zero, got %d", scr.Cursor)
	}
}
