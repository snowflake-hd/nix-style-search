package index

import (
	"strings"
	"testing"
)

func TestScreenViewError(t *testing.T) {
	scr := Screen{Err: "boom"}
	out := scr.View()
	if want := "Failed to load indices"; !contains(out, want) {
		t.Fatalf("expected error message, got %q", out)
	}
}

func TestScreenViewLoading(t *testing.T) {
	scr := Screen{}
	out := scr.View()
	if want := "Loading indices"; !contains(out, want) {
		t.Fatalf("expected loading message, got %q", out)
	}
}

func TestScreenViewNormal(t *testing.T) {
	scr := Screen{Indices: []string{"alpha", "beta"}, Cursor: 1}
	out := scr.View()
	if want := "beta"; !contains(out, want) {
		t.Fatalf("expected to list indices, got %q", out)
	}
	if want := "▶"; !contains(out, want) {
		t.Fatalf("expected cursor marker, got %q", out)
	}
}

func contains(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}
