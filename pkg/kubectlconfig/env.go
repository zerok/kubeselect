package kubectlconfig

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// Environment is an abstraction to the environment settings that are
// supported by kubeselect.
type Environment struct {
	KubeConfig string
	Context    string
}

// Loads an environment based on the given slice of environment
// variables. Internally, this also does some validation for which
// filesystem access is required.
func LoadEnvironment(ctx context.Context, filesystem afero.Fs, environ []string) (*Environment, error) {
	e := &Environment{}
	for _, env := range environ {
		if strings.HasPrefix(env, "KUBECONFIG=") {
			e.KubeConfig = env[11:]
		} else if strings.HasPrefix(env, "KUBECTX=") {
			e.Context = env[8:]
		}
	}

	if e.KubeConfig == "" {
		return e, nil
	}

	// Now also ensure that the context is actually present in the
	// specified config file. If it isn't, reset the context to
	// the one defined as currently active inside the kubeconfig.
	cfg, err := LoadConfigFile(ctx, filesystem, e.KubeConfig)
	if err != nil {
		return e, err
	}

	if e.Context == "" {
		if cfg.CurrentContext != "" {
			e.Context = cfg.CurrentContext
		} else {
			if cfg.Contexts != nil {
				for _, c := range cfg.Contexts {
					e.Context = c.Name
					break
				}
			}
		}
	} else {
		if !cfg.HasContext(e.Context) {
			return e, errors.Errorf("context '%s' doesn't exist", e.Context)
		}
	}
	return e, nil
}
