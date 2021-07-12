package jsonds

import (
	"encoding/json"
	"os"

	mah "github.com/BrenekH/my-aur-helper"
)

func Create(filepath string) (JsonDS, error) {
	j := JsonDS{
		filepath: filepath,
	}

	// Attempt to load, and write the file if an error occurs
	err := j.Load()
	if err != nil {
		if err = j.Save(); err != nil {
			return j, err
		}
		err = j.Load()
	}

	return j, err
}

type JsonDS struct {
	filepath      string
	knownPackages []mah.Package
}

func (j *JsonDS) KnownPackages() ([]mah.Package, error) {
	return j.knownPackages, nil
}

func (j *JsonDS) AddKnownPackage(p mah.Package) error {
	j.knownPackages = append(j.knownPackages, p)
	return j.Save()
}

func (j *JsonDS) Load() error {
	// Read file
	b, err := os.ReadFile(j.filepath)
	if err != nil {
		return err
	}

	var jS jsonStruct
	if err = json.Unmarshal(b, &jS); err != nil {
		return err
	}

	j.knownPackages = jS.Packages

	return nil
}

func (j *JsonDS) Save() error {
	// Marshal data
	jS := jsonStruct{
		Packages: j.knownPackages,
	}

	b, err := json.MarshalIndent(jS, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(j.filepath, b, 0777)
}

type jsonStruct struct {
	Packages []mah.Package `json:"packages"`
}
