package myaurhelper

import (
	"fmt"
	"os"
	"os/exec"
)

func Install(pkgs []string) error {
	for _, pkg := range pkgs {
		// Download PKGBUILD and other files from the AUR
		//! If the folder is already populated, git pull should be run instead of clone
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		cloneDir := fmt.Sprintf("%s/.mah/%s", home, pkg)

		cmd := exec.Command("git", "clone", fmt.Sprintf("https://aur.archlinux.org/%s.git", pkg), cloneDir)
		if err = cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Downloaded AUR repo")

		// TODO: Display the PKGBUILD to the user using the $PAGER env var (or use less if not set)

		// Run makepkg against the downloaded files
		cmd = exec.Command("makepkg", "-si")
		cmd.Dir = cloneDir

		// Connect console to makepkg process so that the user can provide their password for elevation and allow pacman to install
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

		if err = cmd.Run(); err != nil {
			return err
		}

		fmt.Printf("Package %v installed\n", pkg)
	}

	return nil
}
