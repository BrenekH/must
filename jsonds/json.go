package jsonds

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BrenekH/must"
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
	knownPackages []must.Package
}

func (j *JsonDS) KnownPackages() ([]must.Package, error) {
	return j.knownPackages, nil
}

func (j *JsonDS) AddKnownPackage(p must.Package) error {
	j.knownPackages = append(j.knownPackages, p)
	return j.Save()
}

func (j *JsonDS) UpdatePackage(p must.Package) error {
	var pkgFound bool
	for i := 0; i < len(j.knownPackages); i++ {
		if j.knownPackages[i].Name == p.Name {
			j.knownPackages[i] = p
			pkgFound = true
			break
		}
	}

	if !pkgFound {
		return fmt.Errorf("unable to find package %v in database", p.Name)
	}

	return j.Save()
}

func (j *JsonDS) RemovePackage(pname string) error {
	for i := 0; i < len(j.knownPackages); i++ {
		if pname == j.knownPackages[i].Name {
			j.knownPackages = removeFromSlice(j.knownPackages, i)
			break
		}
	}
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
	Packages []must.Package `json:"packages"`
}

// removeFromSlice uses an index to remove an element from a slice. Source: https://stackoverflow.com/a/37335777
func removeFromSlice(s []must.Package, i int) []must.Package {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
