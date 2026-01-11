package flags

import (
	"flag"
	"fmt"
)

// DefineFlags parses command-line flags and returns true if the application
// should continue running, or false if it should exit early (for example,
// after displaying help text).
func DefineFlags() bool {
	help := false
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&help, "h", false, "Show help")

	flag.Parse()

	if help {
		fmt.Println("Usage: nix-style-search [options]")
		fmt.Println("A command-line tool for searching Nix packages.")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		return false
	}

	return true
}
