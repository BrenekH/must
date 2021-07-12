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

		// TODO: Show git diff of PKGBUILD

		cmd := exec.Command("makepkg", "-si")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.Dir = ac.AppDir + "/" + pkg.Name

		if err = cmd.Run(); err != nil {
			return fmt.Errorf("makepkg: %v", err)
		}

		fmt.Printf("%v upgraded\n", pkg.Name)
	}

	return nil
}
