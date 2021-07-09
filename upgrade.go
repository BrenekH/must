package myaurhelper

import "errors"

func Upgrade() error {
	// TODO: Identify all installed packages
	// TODO: Check installed versions against local PKGBUILD
	// TODO: If version is out-of-date, run makepkg and install
	return errors.New("not implemented")
}
