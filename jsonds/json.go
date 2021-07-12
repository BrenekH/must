package jsonds

import (
	mah "github.com/BrenekH/my-aur-helper"
)

func Create(filepath string) JsonDS {
	return JsonDS{
		filepath: filepath,
	}
}

type JsonDS struct {
	filepath string
}

func (j *JsonDS) KnownPackages() ([]mah.Package, error) {
	return []mah.Package{}, nil
}

func (j *JsonDS) AddKnownPackage(p mah.Package) error {
	return nil
}
