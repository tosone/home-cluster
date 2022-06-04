package types

// Source ...
type Source struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	URL       string `yaml:"url"`
	Branch    string `yaml:"branch"`
	Tag       string `yaml:"tag"`
}

// HelmRelease ...
type HelmRelease struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Chart     string `yaml:"chart"`
	Source    string `yaml:"source"`
	Values    string `yaml:"values"`
}

// Deploy ...
type Deploy struct {
	Sources      map[string][]Source `yaml:"sources"`
	HelmReleases []HelmRelease       `yaml:"helmreleases"`
}
