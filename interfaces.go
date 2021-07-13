package myaurhelper

type DataStorer interface {
	AddKnownPackage(Package) error
	KnownPackages() ([]Package, error)
	UpdatePackage(Package) error
	RemovePackage(name string) error
}
