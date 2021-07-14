package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BrenekH/must"
	"github.com/BrenekH/must/jsonds"
)

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
