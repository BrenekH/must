package must

type AppConfig struct {
	DS     DataStorer
	AppDir string
}

type Package struct {
	Name            string `json:"name"`
	UpdateAvailable bool   `json:"update_available"`
}
