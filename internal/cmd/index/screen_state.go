package index

import (
	"NixStyleSearch/internal/search"
)

var fetchAliases = search.GetAvailableIndices

func (i Screen) clampCursor() Screen {
	switch {
	case len(i.Indices) == 0:
		i.Cursor = 0
	case i.Cursor < 0:
		i.Cursor = 0
	case i.Cursor > len(i.Indices)-1:
		i.Cursor = len(i.Indices) - 1
	}
	return i
}

func (i Screen) loadIndices() Screen {
	aliases, err := fetchAliases()
	if err != nil {
		i.Err = err.Error()
		i.Indices = nil
		return i
	}

	cleanedAlias := cleanAlias(&aliases)

	indices := make([]string, 0, len(*cleanedAlias))
	for _, a := range *cleanedAlias {
		indices = append(indices, a.Alias)
	}

	i.Indices = indices
	i.Err = ""
	return i.clampCursor()
}

func cleanAlias(indices *[]search.Alias) *[]search.Alias {
	cleanedAlias := make([]search.Alias, 0, len(*indices))

	for _, alias := range *indices {
		if alias.Alias == "" || alias.Index == "" {
			continue
		}
		// Indices / aliases starting with a '.' are hidden
		if alias.Alias[0] == '.' || alias.Index[0] == '.' {
			continue
		}

		cleanedAlias = append(cleanedAlias, alias)
	}

	return &cleanedAlias
}

func (i *Screen) moveCursor(delta int) {
	if len(i.Indices) == 0 {
		i.Cursor = 0
		return
	}
	i.Cursor += delta
	if i.Cursor < 0 {
		i.Cursor = 0
	}
	if i.Cursor > len(i.Indices)-1 {
		i.Cursor = len(i.Indices) - 1
	}
}
