package main

import (
	"fmt"
	"os"
	"strings"

	myaurhelper "github.com/BrenekH/my-aur-helper"
	"github.com/BrenekH/my-aur-helper/jsonds"
)

func main() {
	// Commands
	//   - update - Refresh known package versions
	//   - upgrade - Upgrade packages
	//   - install - Install new packages
	//   - remove - Uninstall specified packages

	if len(os.Args) < 2 {
		fmt.Println("expected at least 2 arguments")
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory")
		os.Exit(1)
	}

	jsonDS, err := jsonds.Create(home + "/.mah/db.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ac := myaurhelper.AppConfig{
		DS:     &jsonDS,
		AppDir: home + "/.mah",
	}

	switch strings.ToLower(os.Args[1]) {
	case "update":
		if err := myaurhelper.Update(ac); err != nil {
			fmt.Printf("Update command failed with error: %v\n", err)
		}

	case "upgrade":
		if err := myaurhelper.Upgrade(ac); err != nil {
			fmt.Printf("Upgrade command failed with error: %v\n", err)
		}

	case "install":
		if err := myaurhelper.Install(ac, os.Args[2:]); err != nil {
			fmt.Printf("Install command failed with error: %v\n", err)
		}

	case "remove":
		if err := myaurhelper.Remove(ac, os.Args[2:]); err != nil {
			fmt.Printf("Remove command failed with error: %v\n", err)
		}

	default:
		fmt.Printf("Unknown command: %v\n", os.Args[1])
	}
}
