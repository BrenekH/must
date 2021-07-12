package myaurhelper

type DataStorer interface {
	AddKnownPackage(Package) error
	KnownPackages() ([]Package, error)
}
