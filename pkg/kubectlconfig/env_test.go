package kubectlconfig_test

import (
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"gitlab.com/zerok/kubeselect/pkg/kubectlconfig"
)

func TestLoadEnv(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		env      []string
		expected *kubectlconfig.Environment
		err      bool
		setupFS  func() afero.Fs
	}{
		{
			env:      []string{},
			expected: &kubectlconfig.Environment{},
			err:      false,
			setupFS: func() afero.Fs {
				return afero.NewMemMapFs()
			},
		},
		{
			env: []string{
				"KUBECONFIG=some-path",
				"KUBECTX=some-ctx",
			},
			expected: &kubectlconfig.Environment{
				KubeConfig: "some-path",
				Context:    "some-ctx",
			},
			// Because the path doesn't exist
			err: true,
			setupFS: func() afero.Fs {
				return afero.NewMemMapFs()
			},
		},
		{
			env: []string{
				"KUBECONFIG=/exists.yaml",
				"KUBECTX=some-ctx",
			},
			expected: &kubectlconfig.Environment{
				KubeConfig: "/exists.yaml",
				Context:    "some-ctx",
			},
			err: false,
			setupFS: func() afero.Fs {
				fs := afero.NewMemMapFs()
				afero.WriteFile(fs, "/exists.yaml", []byte("contexts: [{name: some-ctx}]"), 0600)
				return fs
			},
		},
		// If a context is set that doesn't exist in the
		// kubeconfig, then an error is returned.
		{
			env: []string{
				"KUBECONFIG=/exists.yaml",
				"KUBECTX=not-existing-ctx",
			},
			expected: &kubectlconfig.Environment{
				KubeConfig: "/exists.yaml",
				Context:    "not-existing-ctx",
			},
			err: true,
			setupFS: func() afero.Fs {
				fs := afero.NewMemMapFs()
				afero.WriteFile(fs, "/exists.yaml", []byte("contexts: [{name: some-ctx}]"), 0600)
				return fs
			},
		},
		// If no context is specified, then the first found
		// context is used.
		{
			env: []string{
				"KUBECONFIG=/exists.yaml",
				"KUBECTX=",
			},
			expected: &kubectlconfig.Environment{
				KubeConfig: "/exists.yaml",
				Context:    "some-ctx",
			},
			err: false,
			setupFS: func() afero.Fs {
				fs := afero.NewMemMapFs()
				afero.WriteFile(fs, "/exists.yaml", []byte("contexts: [{name: some-ctx}]"), 0600)
				return fs
			},
		},
		// If no context is specified, then the first found
		// context is used.
		{
			env: []string{
				"KUBECONFIG=/exists.yaml",
				"KUBECTX=",
			},
			expected: &kubectlconfig.Environment{
				KubeConfig: "/exists.yaml",
				Context:    "active",
			},
			err: false,
			setupFS: func() afero.Fs {
				fs := afero.NewMemMapFs()
				afero.WriteFile(fs, "/exists.yaml", []byte("current-context: active\ncontexts: [{name: some-ctx},{name:active}]"), 0600)
				return fs
			},
		},
	}

	for _, test := range tests {
		fs := test.setupFS()
		actual, err := kubectlconfig.LoadEnvironment(ctx, fs, test.env)
		if test.err {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, actual, test.expected)
	}
}
