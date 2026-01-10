package search

import (
	"NixStyleSearch/internal/search"
	"strings"
	"testing"
)

func TestFirstNonEmpty(t *testing.T) {
	if got := firstNonEmpty("", "  ", "alpha", "beta"); got != "alpha" {
		t.Fatalf("expected alpha, got %q", got)
	}

	if got := firstNonEmpty("", ""); got != "" {
		t.Fatalf("expected empty string when all values blank, got %q", got)
	}
}

func TestExtractPackageName(t *testing.T) {
	src := search.PackageSource{PackageAttrName: "hello", PackagePversion: "1.2.3"}
	if got := extractPackageName(src); got != "hello (1.2.3)" {
		t.Fatalf("expected name with version, got %q", got)
	}

	src = search.PackageSource{PackagePname: "pkg"}
	if got := extractPackageName(src); got != "pkg" {
		t.Fatalf("expected fallback pname, got %q", got)
	}

	src = search.PackageSource{}
	if got := extractPackageName(src); got != "" {
		t.Fatalf("expected empty when no names, got %q", got)
	}
}

func TestGetSearchPackageString(t *testing.T) {
	pkgs := []string{"alpha", "beta"}
	out := getSearchPackageString(pkgs, 1, 5)
	if !strings.Contains(out, "  ▶  7. beta") {
		t.Fatalf("expected cursor line with numbering, got %q", out)
	}
	if !strings.Contains(out, "     6. alpha") {
		t.Fatalf("expected non-cursor line, got %q", out)
	}
}
