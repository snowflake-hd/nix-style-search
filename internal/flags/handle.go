package flags

import (
	"flag"
	"fmt"
)

func HandleFlags() bool {
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		fmt.Println("Usage: nix-style-search")
		fmt.Println("A command-line tool for searching Nix packages.")
		return false
	}

	return true
}
