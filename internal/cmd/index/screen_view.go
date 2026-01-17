package index

import (
	"fmt"
	"strings"

	"nix-style-search/internal"
)

func (i Screen) View() string {
	if i.Err != "" {
		return fmt.Sprintf(internal.INDEX_ERROR_TEMPLATE, i.Err)
	}

	if len(i.Indices) == 0 {
		return internal.INDEX_LOADING_MESSAGE
	}

	var b strings.Builder
	b.WriteString(internal.INDEX_TITLE)

	for idx, name := range i.Indices {
		num := fmt.Sprintf("%2d.", idx+1)
		if idx == i.Cursor {
			b.WriteString(fmt.Sprintf(internal.CURSOR_DISPLAY, num, name))
		} else {
			b.WriteString(fmt.Sprintf("    %s %s\n", num, name))
		}
	}

	b.WriteString(internal.INDEX_HELP)
	return b.String()
}
