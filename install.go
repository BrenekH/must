package myaurhelper

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		cmd = exec.Command("makepkg")
		cmd.Dir = cloneDir
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err = cmd.Run(); err != nil {
			return err
		}

		fmt.Println("makepkg command completed")

		// Install .pkg.tar.zst using Pacman
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.pkg.tar.zst", cloneDir))
		if err != nil {
			return err
		}

		fmt.Printf("Installing archive for %v\n", pkg)
		cmd = exec.Command("sudo", "pacman", "--noconfirm", "-U", matches[0]) // This hard coded index is not great -_-

		// Connect the user's terminal stdout and stdin to the pacman process (this allows the user to enter a password for sudo)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		if err = cmd.Run(); err != nil {
			return err
		}

		fmt.Printf("Package %v installed\n", pkg)
	}

	return nil
}
