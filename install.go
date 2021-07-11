package myaurhelper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Install(pkgs []string) error {
	for _, pkg := range pkgs {
		// Download PKGBUILD and other files from the AUR
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		//! If the folder is already populated, error out or automatically use the update code instead

		cloneDir := fmt.Sprintf("%s/.mah/%s", home, pkg)

		cmd := exec.Command("git", "clone", fmt.Sprintf("https://aur.archlinux.org/%s.git", pkg), cloneDir)
		if err = cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Downloaded AUR repo")

		// Display the PKGBUILD to the user using the $PAGER env var (or manually display if not set)
		if pagerBin, exists := os.LookupEnv("PAGER"); exists {
			cmd = exec.Command(pagerBin, cloneDir+"/PKGBUILD")

			// Connect console to makepkg process so that the user can provide their password for elevation and allow pacman to install
			cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

			if err = cmd.Run(); err != nil {
				return err
			}
		} else {
			fmt.Println("PAGER environment variable not set. Outputting directly to standard output.")

			if b, err := os.ReadFile(cloneDir + "/PKGBUILD"); err == nil {
				fmt.Println("---------- BEGIN FILE ----------")
				fmt.Println(string(b))
				fmt.Println("----------- END FILE -----------")
			} else {
				return err
			}
		}

		fmt.Print("Continue the installation (y/N)? ")
		var response string
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			fmt.Printf("Skipping installation of %v\n", pkg)
			continue
		}

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
