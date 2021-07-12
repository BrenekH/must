package myaurhelper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Update(ac AppConfig) error {
	pkgs, err := ac.DS.KnownPackages()
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		pkgDir := fmt.Sprintf("%v/%v", ac.AppDir, pkg.Name)

		cmd := exec.Command("git", "pull")

		// Connect user to git process so they know what is going on
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

		// Run git pull in the package directory instead of the current working directory
		cmd.Dir = pkgDir

		b, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("git pull: %v", err)
		}

		if strings.TrimSpace(string(b)) == "Already up to date." {
			return nil
		}

		pkg.UpdateAvailable = true
		if err = ac.DS.UpdatePackage(pkg); err != nil {
			return fmt.Errorf("update database: %v", err)
		}
	}

	return nil
}
