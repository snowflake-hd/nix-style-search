package search

import (
	"NixStyleSearch/internal/search"
	"strings"
)

var searchPageFn = searchPage

func (s *Screen) fetchForward(overlap bool) {
	sources, fetched, err := searchPageFn(s.Index, s.query, s.next, pageSize)
	if err != nil {
		s.Err = err.Error()
		return
	}

	s.Err = ""

	if overlap && len(s.Sources) > 0 {
		last := s.Sources[len(s.Sources)-1]
		sources = append([]search.PackageSource{last}, sources...)
		s.start = s.next - 1
	} else {
		s.start = s.next
	}

	s.Sources = sources
	s.Pkgs = buildPackageNames(s.Sources)

	s.next += fetched
	if fetched < pageSize {
		s.endReached = true
	}
}

func (s *Screen) fetchBackward() {
	if s.start == 0 {
		return
	}

	fetchStart := s.start - (pageSize - 1)
	if fetchStart < 0 {
		fetchStart = 0
	}

	sources, fetched, err := searchPageFn(s.Index, s.query, fetchStart, pageSize)
	if err != nil {
		s.Err = err.Error()
		return
	}
	if fetched == 0 {
		s.endReached = true
		return
	}

	if len(s.Sources) > 0 {
		sources = append(sources, s.Sources[0])
	}

	s.Err = ""
	s.Sources = sources
	s.Pkgs = buildPackageNames(s.Sources)
	s.start = fetchStart
	s.next = s.start + len(s.Sources)
	s.endReached = false
}

func searchPage(index, query string, from, size int) ([]search.PackageSource, int, error) {
	sources, err := searchPackages(index, query, from, size)
	if err != nil {
		return nil, 0, err
	}
	return sources, len(sources), nil
}

func searchPackages(index, query string, from, size int) ([]search.PackageSource, error) {
	if strings.TrimSpace(query) == "" {
		return nil, nil
	}

	result, err := search.Query(query, index, from, size)
	if err != nil {
		return nil, err
	}

	hits := result.Hits.Hits
	sources := make([]search.PackageSource, 0, len(hits))
	for _, h := range hits {
		sources = append(sources, h.Source)
	}

	return sources, nil
}

func buildPackageNames(sources []search.PackageSource) []string {
	names := make([]string, 0, len(sources))
	for _, src := range sources {
		names = append(names, extractPackageName(src))
	}
	return names
}
