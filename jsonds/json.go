package jsonds

import (
	mah "github.com/BrenekH/my-aur-helper"
)

func Create() JsonDS {
	return JsonDS{}
}

type JsonDS struct {
}

func (j *JsonDS) KnownPackages() ([]mah.Package, error) {
	return []mah.Package{}, nil
}

func (j *JsonDS) AddKnownPackage(p mah.Package) error {
	return nil
}
