package must

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

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
		cmd := exec.Command("pacman", "-Qi", result.Name)
		b, err := cmd.Output()
		if err != nil {
			return err
		}

		var re = regexp.MustCompile(`Version *: (.+)`)
		regexResult := re.FindStringSubmatch(string(b))

		if len(regexResult) < 2 {
			return fmt.Errorf("failed to parse version for package '%v' from pacman", result.Name)
		}

		currentPkgVersion := regexResult[1]

		if !isPackageUpgradeable(currentPkgVersion, result.Version) {
			continue
		}

		for _, pkg := range pkgs {
			if pkg.Name == result.Name {
				pkg.UpdateAvailable = true
				if err = ac.DS.UpdatePackage(pkg); err != nil {
					return fmt.Errorf("update database: %v", err)
				}
			}
		}
	}

	fmt.Println("[must] Update complete")

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
		fmt.Printf("[must] There are %v packages available to upgrade\n", numPkgs)
	}

	return nil
}

func isPackageUpgradeable(currentVersion, latestVersion string) bool {
	cmd := exec.Command("vercmp", currentVersion, latestVersion)
	b, err := cmd.Output()
	if err != nil {
		fmt.Println("[must] Could not call vercmp. Please check that the Pacman is properly installed.")
		os.Exit(1)
	}

	switch strings.TrimSpace(string(b)) {
	case "0", "1":
		return false
	case "-1":
		return true
	default:
		fmt.Printf("[must] Received an unknown vercmp value %v\n", strings.TrimSpace(string(b)))
		return false
	}
}
