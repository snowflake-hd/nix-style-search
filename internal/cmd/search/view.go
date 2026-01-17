package search

import (
	"fmt"
	"strings"

	"nix-style-search/internal"
	"nix-style-search/internal/search"

	"github.com/atotto/clipboard"
)

func (s *Screen) View() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Search in index: %s\n\n", s.Index))
	b.WriteString(s.Ti.View())
	b.WriteString("\n\n")

	if s.Err != "" {
		b.WriteString(fmt.Sprintf(internal.SEARCH_ERROR_TEMPLATE, s.Err))
	}

	switch {
	case strings.TrimSpace(s.Ti.Value()) == "":
		b.WriteString(internal.SEARCH_PROMPT)
	case len(s.Pkgs) == 0:
		b.WriteString(internal.SEARCH_NO_RESULTS)
	default:
		b.WriteString(internal.SEARCH_RESULTS_TITLE)
		b.WriteString(getSearchPackageString(s.Pkgs, s.Cursor, s.start))

		if s.ShowDetails && len(s.Sources) > 0 {
			detail := s.Sources[s.Cursor]
			b.WriteString(internal.SEARCH_DETAILS_TITLE)
			b.WriteString(renderPackageDetail(detail))
		} else if !s.ShowDetails {
			b.WriteString(internal.SEARCH_DETAILS_PROMPT)
		}
	}

	b.WriteString(internal.INDEX_HELP)
	return b.String()
}

func getSearchPackageString(pkgs []string, cursor int, offset int) string {
	var b strings.Builder
	for idx, name := range pkgs {
		num := fmt.Sprintf("%2d.", offset+idx+1)
		if idx == cursor {
			b.WriteString(fmt.Sprintf(internal.CURSOR_DISPLAY, num, name))
		} else {
			b.WriteString(fmt.Sprintf("    %s %s\n", num, name))
		}
	}
	return b.String()
}

func renderPackageDetail(src search.PackageSource) string {
	var b strings.Builder

	if name := extractPackageName(src); name != "" {
		b.WriteString(fmt.Sprintf(internal.SEARCH_NAME_TEMPLATE, name))
	}
	if src.PackageDescription != "" {
		b.WriteString(fmt.Sprintf(internal.SEARCH_SUMMARY_TEMPLATE, src.PackageDescription))
	}
	if src.PackagePversion != "" {
		b.WriteString(fmt.Sprintf(internal.SEARCH_VERSION_TEMPLATE, src.PackagePversion))
	}
	if src.PackageAttrName != "" {
		b.WriteString(fmt.Sprintf(internal.SEARCH_ATTR_TEMPLATE, src.PackageAttrName))
	}
	if len(src.PackageHomepage) > 0 {
		b.WriteString(fmt.Sprintf(internal.SEARCH_HOMEPAGE_TEMPLATE, src.PackageHomepage[0]))
	}
	if len(src.PackageLicenseSet) > 0 {
		b.WriteString(fmt.Sprintf(internal.SEARCH_LICENSE_TEMPLATE, strings.Join(src.PackageLicenseSet, ", ")))
	}
	return b.String()
}

func extractPackageName(src search.PackageSource) string {
	name := firstNonEmpty(
		src.PackageAttrName,
		src.PackagePname,
		src.PackageMainProgram,
		src.PackageAttrSet,
	)
	if name == "" {
		return ""
	}
	if version := src.PackagePversion; version != "" {
		return fmt.Sprintf("%s (%s)", name, version)
	}
	return name
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func (s *Screen) showSelection() {
	if len(s.Pkgs) == 0 {
		fmt.Println(internal.SEARCH_NO_PACKAGE_SELECTED)
		return
	}
	if s.Cursor < 0 || s.Cursor >= len(s.Pkgs) {
		fmt.Println(internal.SEARCH_SELECTION_OOB)
		return
	}

	identifier := s.selectedPackageIdentifier()
	if identifier == "" {
		identifier = s.Pkgs[s.Cursor]
	}

	if err := clipboard.WriteAll(identifier); err != nil {
		fmt.Printf(internal.SEARCH_COPY_FAILED, err)
		return
	}
	fmt.Printf(internal.SEARCH_PACKAGE_SELECTED, identifier)
}

func (s *Screen) selectedPackageIdentifier() string {
	if s.Cursor < 0 || s.Cursor >= len(s.Sources) {
		return ""
	}
	src := s.Sources[s.Cursor]
	return firstNonEmpty(src.PackageAttrName, src.PackagePname, src.PackageMainProgram)
}
