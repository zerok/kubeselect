package kubectlconfig

type ConfigFile struct {
	Path           string    `yaml:"-"`
	APIVersion     string    `yaml:"apiVersion"`
	Kind           string    `yaml:"kind"`
	CurrentContext string    `yaml:"current-context"`
	Clusters       []Cluster `yaml:"clusters"`
	Contexts       []Context `yaml:"contexts"`
}

type Cluster struct{}

type Context struct {
	Name string `yaml:"name"`
}
