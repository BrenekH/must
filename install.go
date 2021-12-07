package must

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Install(ac AppConfig, pkgs []string) error {
	for _, pkg := range pkgs {
		// Download PKGBUILD and other files from the AUR
		cloneDir := fmt.Sprintf("%s/%s", ac.AppDir, pkg)
		//! If the folder is already populated, error out or automatically use the update code instead

		cmd := exec.Command("git", "clone", fmt.Sprintf("https://aur.archlinux.org/%s.git", pkg), cloneDir)

		// Connect console to the git process so that the user can see any issues and provide any input it requires
		cmd.Stdout, cmd.Stdin, cmd.Stderr = os.Stdout, os.Stdin, os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("git clone: %v", err)
		}

		fmt.Println("[must] Downloaded AUR repo")

		// Display the PKGBUILD to the user using the $PAGER env var (or manually display if not set)
		if pagerBin, exists := os.LookupEnv("PAGER"); exists {
			cmd = exec.Command(pagerBin, cloneDir+"/PKGBUILD")

			// Connect console to the pager process so that the user interact with it properly
			cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

			if err := cmd.Run(); err != nil {
				return fmt.Errorf("pager: %v", err)
			}
		} else {
			fmt.Println("[must] PAGER environment variable not set. Outputting directly to standard output.")

			if b, err := os.ReadFile(cloneDir + "/PKGBUILD"); err == nil {
				fmt.Println("---------- BEGIN FILE ----------")
				fmt.Println(string(b))
				fmt.Println("----------- END FILE -----------")
			} else {
				return fmt.Errorf("reading PKGBUILD: %v", err)
			}
		}

		fmt.Print("Continue the installation (y/N)? ")
		var response string
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			fmt.Printf("[must] Skipping installation of %v\n", pkg)
			// TODO: Remove cloned repository
			continue
		}

		// Run makepkg against the downloaded files
		cmd = exec.Command("makepkg", "-si")
		cmd.Dir = cloneDir

		// Connect console to makepkg process so that the user can provide their password for elevation and allow pacman to install
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("makepkg: %v", err)
		}

		p := Package{Name: pkg, UpdateAvailable: false}
		if err := ac.DS.AddKnownPackage(p); err != nil {
			return fmt.Errorf("storing in database: %v", err)
		}

		fmt.Printf("[must] Package %v installed\n", pkg)
	}

	return nil
}
