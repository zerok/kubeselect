package kubectlconfig

import (
	"context"

	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v2"
)

// ConfigFile represents the content of a kubeconfig
// file. Additionally, this also holds the path of that file inside
// the Path field.
type ConfigFile struct {
	Path           string    `yaml:"-"`
	APIVersion     string    `yaml:"apiVersion"`
	Kind           string    `yaml:"kind"`
	CurrentContext string    `yaml:"current-context"`
	Clusters       []Cluster `yaml:"clusters"`
	Contexts       []Context `yaml:"contexts"`
}

// HasContext checks all the contexts defined in the configfile for
// one with the given name.
func (c *ConfigFile) HasContext(name string) bool {
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			return true
		}
	}
	return false
}

// LoadConfigFile tries to load a ConfigFile from the given path
// inside a filesystem.
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

// Cluster represents a single cluster definition inside a ConfigFile.
type Cluster struct{}

// Context represents a single context definition inside a ConfigFile.
type Context struct {
	Name string `yaml:"name"`
}
