package main

import (
	"fmt"
	"os"
	"strings"

	myaurhelper "github.com/BrenekH/my-aur-helper"
)

func main() {
	// Commands
	//   - update - Refresh known package versions
	//   - upgrade - Upgrades packages
	//   - install - Install new packages

	if len(os.Args) < 2 {
		fmt.Println("expected at least 2 arguments")
		os.Exit(1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "update":
		if err := myaurhelper.Update(); err != nil {
			fmt.Printf("Update command failed with error: %v\n", err)
		}
	case "upgrade":
		if err := myaurhelper.Upgrade(); err != nil {
			fmt.Printf("Upgrade command failed with error: %v\n", err)
		}
	case "install":
		if err := myaurhelper.Install(); err != nil {
			fmt.Printf("Install command failed with error: %v\n", err)
		}
	default:
		fmt.Printf("Unknown command: %v\n", os.Args[1])
	}
}
