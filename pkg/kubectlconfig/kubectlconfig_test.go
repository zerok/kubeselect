package kubectlconfig_test

import (
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"gitlab.com/zerok/kubeselect/pkg/kubectlconfig"
)

func TestFindingConfigFiles(t *testing.T) {
	t.Run("find *.yml and *.yaml", func(t *testing.T) {
		ctx := context.Background()
		fs := afero.NewMemMapFs()
		afero.WriteFile(fs, "/subfolder1/test.yaml", []byte("clusters: []"), 0600)
		afero.WriteFile(fs, "/subfolder2/test.yml", []byte("clusters: []"), 0600)
		res, err := kubectlconfig.FindAllFiles(ctx, fs, "/")
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Len(t, res, 2)
	})
}
