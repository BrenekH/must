package must

import (
	"fmt"
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

		// Run git pull in the package directory instead of the current working directory
		cmd.Dir = pkgDir

		fmt.Printf("Pulling %v\n", pkg.Name)
		b, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("git pull: %v", err)
		}

		if strings.TrimSpace(string(b)) == "Already up to date." {
			continue
		}

		pkg.UpdateAvailable = true
		if err = ac.DS.UpdatePackage(pkg); err != nil {
			return fmt.Errorf("update database: %v", err)
		}
	}

	fmt.Println("Update complete")

	return outputNumPkgsNeedingUpgrade(ac)
}

func outputNumPkgsNeedingUpgrade(ac AppConfig) error {
	pkgs, err := ac.DS.KnownPackages()
	if err != nil {
		return err
	}

	numPkgs := 0
	for _, pkg := range pkgs {
		if pkg.UpdateAvailable {
			numPkgs++
		}
	}

	if numPkgs != 0 {
		fmt.Printf("There are %v packages available to upgrade\n", numPkgs)
	}

	return nil
}
