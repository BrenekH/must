package myaurhelper

import (
	"fmt"
	"os"
	"os/exec"
)

func Remove(ac AppConfig, toRemove []string) error {
	knownPkgs, err := ac.DS.KnownPackages()
	if err != nil {
		return err
	}

	for _, pkgName := range toRemove {
		// Check that the package was installed using mah.
		var installedByMah bool
		for _, kPkg := range knownPkgs {
			if pkgName == kPkg.Name {
				installedByMah = true
				break
			}
		}

		if !installedByMah {
			return fmt.Errorf("%v was not installed using mah install", pkgName)
		}

		// Remove package using Pacman (pacman -Rns)
		cmd := exec.Command("sudo", "pacman", "-Rns", pkgName)

		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

		if err = cmd.Run(); err != nil {
			return fmt.Errorf("pacman -Rns: %v", err)
		}

		// Remove downloaded AUR repo
		if err = os.RemoveAll(ac.AppDir + "/" + pkgName); err != nil {
			return fmt.Errorf("os.RemoveAll: %v", err)
		}

		if err = ac.DS.RemovePackage(pkgName); err != nil {
			return err
		}
	}

	return nil
}
