package myaurhelper

import (
	"fmt"
	"os"
	"os/exec"
)

func Upgrade(ac AppConfig) error {
	pkgs, err := ac.DS.KnownPackages()
	if err != nil {
		return fmt.Errorf("getting known packages from database: %v", err)
	}

	for _, pkg := range pkgs {
		if !pkg.UpdateAvailable {
			continue
		}

		fmt.Printf("Upgrading %v\n", pkg.Name)

		pkgDir := ac.AppDir + "/" + pkg.Name

		// Display the git diff of PKGBUILD to the user using the $PAGER env var (or manually display if not set)
		if pagerBin, exists := os.LookupEnv("PAGER"); exists {
			cmd := exec.Command("git", "diff", "HEAD~1", "PKGBUILD", "|", pagerBin)

			// Connect console to the pager process so that the user interact with it properly
			cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

			if err := cmd.Run(); err != nil {
				return fmt.Errorf("git diff into $PAGER: %v", err)
			}
		} else {
			fmt.Println("PAGER environment variable not set. Outputting directly to standard output.")

			cmd := exec.Command("git", "diff", "HEAD~1", "PKGBUILD")

			// Connect console to the git process so that it natively displays
			cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

			if err := cmd.Run(); err != nil {
				return fmt.Errorf("git diff: %v", err)
			}
		}

		cmd := exec.Command("makepkg", "-si")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.Dir = pkgDir

		if err = cmd.Run(); err != nil {
			return fmt.Errorf("makepkg: %v", err)
		}

		fmt.Printf("%v upgraded\n", pkg.Name)
	}

	return nil
}
