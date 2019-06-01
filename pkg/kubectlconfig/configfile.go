package kubectlconfig

import (
	"context"

	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v2"
)

type ConfigFile struct {
	Path           string    `yaml:"-"`
	APIVersion     string    `yaml:"apiVersion"`
	Kind           string    `yaml:"kind"`
	CurrentContext string    `yaml:"current-context"`
	Clusters       []Cluster `yaml:"clusters"`
	Contexts       []Context `yaml:"contexts"`
}

func (c *ConfigFile) HasContext(name string) bool {
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			return true
		}
	}
	return false
}

func LoadConfigFile(ctx context.Context, filesystem afero.Fs, path string) (*ConfigFile, error) {
	fp, err := filesystem.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	cfg := ConfigFile{Path: path}
	if err := yaml.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Cluster struct{}

type Context struct {
	Name string `yaml:"name"`
}
