package search

import (
	"NixStyleSearch/internal/search"
	"errors"
	"testing"
)

func TestFetchForwardNoOverlap(t *testing.T) {
	original := searchPageFn
	defer func() { searchPageFn = original }()

	searchPageFn = func(index, query string, from, size int) ([]search.PackageSource, int, error) {
		return []search.PackageSource{{PackageAttrName: "alpha"}, {PackageAttrName: "beta"}}, 2, nil
	}

	s := &Screen{Index: "idx", query: "foo", next: 21}
	s.fetchForward(false)

	if s.start != 21 {
		t.Fatalf("expected start to match previous next, got %d", s.start)
	}
	if s.next != 23 {
		t.Fatalf("expected next to advance by fetched items, got %d", s.next)
	}
	if !s.endReached {
		t.Fatalf("expected endReached since fetched < pageSize")
	}
	if len(s.Pkgs) != 2 || s.Pkgs[0] != "alpha" {
		t.Fatalf("expected packages to be rebuilt, got %#v", s.Pkgs)
	}
}

func TestFetchForwardOverlap(t *testing.T) {
	original := searchPageFn
	defer func() { searchPageFn = original }()

	searchPageFn = func(index, query string, from, size int) ([]search.PackageSource, int, error) {
		return []search.PackageSource{{PackageAttrName: "beta"}}, 1, nil
	}

	s := &Screen{
		Index:   "idx",
		query:   "foo",
		next:    10,
		Sources: []search.PackageSource{{PackageAttrName: "alpha"}},
		Pkgs:    []string{"alpha"},
	}

	s.fetchForward(true)

	if len(s.Sources) != 2 {
		t.Fatalf("expected overlap plus new item, got %d", len(s.Sources))
	}
	if s.Sources[0].PackageAttrName != "alpha" || s.Sources[1].PackageAttrName != "beta" {
		t.Fatalf("unexpected source order: %#v", s.Sources)
	}
	if s.start != 9 {
		t.Fatalf("expected start to move back by one, got %d", s.start)
	}
}

func TestFetchForwardError(t *testing.T) {
	original := searchPageFn
	defer func() { searchPageFn = original }()

	searchPageFn = func(string, string, int, int) ([]search.PackageSource, int, error) {
		return nil, 0, errors.New("boom")
	}

	s := &Screen{Index: "idx", query: "foo"}
	s.fetchForward(false)

	if s.Err == "" {
		t.Fatalf("expected error to be captured")
	}
}

func TestFetchBackward(t *testing.T) {
	original := searchPageFn
	defer func() { searchPageFn = original }()

	searchPageFn = func(index, query string, from, size int) ([]search.PackageSource, int, error) {
		return []search.PackageSource{{PackageAttrName: "delta"}}, 1, nil
	}

	s := &Screen{
		Index:   "idx",
		query:   "foo",
		start:   5,
		next:    7,
		Sources: []search.PackageSource{{PackageAttrName: "gamma"}},
		Pkgs:    []string{"gamma"},
	}

	s.fetchBackward()

	if s.start != 0 {
		t.Fatalf("expected start to move backward to 0, got %d", s.start)
	}
	if s.next != len(s.Sources)+s.start {
		t.Fatalf("expected next to follow new window, got %d", s.next)
	}
	if s.Sources[len(s.Sources)-1].PackageAttrName != "gamma" {
		t.Fatalf("expected overlap to keep prior first item, got %#v", s.Sources)
	}
}
