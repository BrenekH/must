package must

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

		cmd := exec.Command("git", "diff", "HEAD~1", "PKGBUILD")
		cmd.Dir = pkgDir

		// Connect console to the git process so that it natively displays
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("git diff: %v", err)
		}

		fmt.Print("Continue the upgrade (y/N)? ")
		var response string
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			fmt.Printf("Skipping upgrade of %v\n", pkg)
			continue
		}

		cmd = exec.Command("makepkg", "-si")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.Dir = pkgDir

		if err = cmd.Run(); err != nil {
			return fmt.Errorf("makepkg: %v", err)
		}

		fmt.Printf("%v upgraded\n", pkg.Name)
	}

	return nil
}
