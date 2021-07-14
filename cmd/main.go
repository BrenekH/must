package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BrenekH/must"
	"github.com/BrenekH/must/jsonds"
)

var Version = "dev"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected at least 2 arguments")
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		os.Exit(1)
	}

	// Ensure ~/.must exists
	if err = os.MkdirAll(home+"/.must", 0644); err != nil {
		fmt.Printf("Error creating ~/.must: %v\n", err)
		os.Exit(1)
	}

	jsonDS, err := jsonds.Create(home + "/.must/db.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ac := must.AppConfig{
		DS:     &jsonDS,
		AppDir: home + "/.must",
	}

	switch strings.ToLower(os.Args[1]) {
	case "--help", "-h":
		displayHelpMessage()

	case "--version", "-v":
		fmt.Printf("Must v%v\n", Version)

	case "update":
		if err := must.Update(ac); err != nil {
			fmt.Printf("Update command failed with error: %v\n", err)
		}

	case "upgrade":
		if err := must.Upgrade(ac); err != nil {
			fmt.Printf("Upgrade command failed with error: %v\n", err)
		}

	case "install":
		if err := must.Install(ac, os.Args[2:]); err != nil {
			fmt.Printf("Install command failed with error: %v\n", err)
		}

	case "remove":
		if err := must.Remove(ac, os.Args[2:]); err != nil {
			fmt.Printf("Remove command failed with error: %v\n", err)
		}

	default:
		fmt.Printf("Unknown command: %v\n", os.Args[1])
	}
}

func displayHelpMessage() {
	fmt.Printf(`Must v%v Help

--help, -h    - Show this message
--version, -v - Show the application version

Commands:
    update                        - Pull the underlying git repos and check if any have new upgrades available.
    upgrade                       - Based on the results of update, upgrade any packages which are known to have new versions.
    install <packages to install> - Download and install each package in a space-delimited list of packages available in the AUR.
    remove <packages to remove>   - Uninstall and remove the build files for a space-delimited list of packages that were installed my Must.
`, Version)
	fmt.Println("") // Print an empty string to append a newline (the backtick style of strings doesn't allow for escape chars like '\n')
}
