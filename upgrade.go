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

		fmt.Printf("[must] Upgrading %v\n", pkg.Name)
		pkgDir := ac.AppDir + "/" + pkg.Name

		// Run git pull in the package directory instead of the current working directory
		gitPullCMD := exec.Command("git", "pull")
		gitPullCMD.Dir = pkgDir

		// fmt.Printf("[must] Pulling %v\n", pkg.Name)
		if err := gitPullCMD.Run(); err != nil {
			return fmt.Errorf("git pull: %v", err)
		}

		gitDiffCMD := exec.Command("git", "diff", "HEAD~1", "PKGBUILD")
		gitDiffCMD.Dir = pkgDir

		// Connect console to the git process so that it natively displays
		gitDiffCMD.Stdin, gitDiffCMD.Stdout, gitDiffCMD.Stderr = os.Stdin, os.Stdout, os.Stderr

		if err := gitDiffCMD.Run(); err != nil {
			return fmt.Errorf("git diff: %v", err)
		}

		fmt.Print("Continue the upgrade (y/N)? ")
		var response string
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			fmt.Printf("[must] Skipping upgrade of %v\n", pkg)
			continue
		}

		gitDiffCMD = exec.Command("makepkg", "-si")
		gitDiffCMD.Stdin, gitDiffCMD.Stdout, gitDiffCMD.Stderr = os.Stdin, os.Stdout, os.Stderr
		gitDiffCMD.Dir = pkgDir

		if err = gitDiffCMD.Run(); err != nil {
			return fmt.Errorf("makepkg: %v", err)
		}

		pkg.UpdateAvailable = false
		if err = ac.DS.UpdatePackage(pkg); err != nil {
			return fmt.Errorf("mark update as complete: %v", err)
		}

		fmt.Printf("[must] %v upgraded\n", pkg.Name)
	}

	return nil
}
