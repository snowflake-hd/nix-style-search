package flags

import (
	"flag"
	"fmt"
)

func DefineFlags() bool {
	help := false
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&help, "h", false, "Show help")

	flag.Parse()

	if help {
		fmt.Println("Usage: nix-style-search")
		fmt.Println("A command-line tool for searching Nix packages.")
		return false
	}

	return true
}
