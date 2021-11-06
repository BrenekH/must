package must

import (
	"fmt"

	"github.com/mikkeloscar/aur"
)

func Update(ac AppConfig) error {
	pkgs, err := ac.DS.KnownPackages()
	if err != nil {
		return err
	}

	pkgsToFind := make([]string, 0)
	for _, pkg := range pkgs {
		pkgsToFind = append(pkgsToFind, pkg.Name)
	}

	results, err := aur.Info(pkgsToFind)
	if err != nil {
		return err
	}

	for _, result := range results {
		_ = result
		// TODO: Query current version from database
		// TODO: Update db with update available flag if db version is less than current version
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
