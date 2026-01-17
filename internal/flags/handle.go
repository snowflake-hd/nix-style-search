package flags

import (
	"flag"
	"fmt"
)

// HandleFlags parses command-line flags and returns true if the application should run.
func HandleFlags() bool {
	help := false
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&help, "h", false, "Show help")

	flag.Parse()

	if help {
		printHelp()
		return false
	}

	return true
}

func printHelp() {
	fmt.Println("Usage: nix-style-search [options]")
	fmt.Println("A command-line tool for searching Nix packages.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help\t\tShow help")
	fmt.Println()
	fmt.Println("While searching for packages use <tab> to see more details about the selected package.")
}
