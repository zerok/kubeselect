package kubectlconfig

import (
	"context"
	"os"
	"strings"

	"github.com/spf13/afero"
)

func FindAllFiles(ctx context.Context, filesystem afero.Fs, root string) ([]ConfigFile, error) {
	result := make([]ConfigFile, 0, 10)
	err := afero.Walk(filesystem, root, func(path string, info os.FileInfo, err error) error {
		if !(strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".conf")) {
			return nil
		}
		cfg, err := LoadConfigFile(ctx, filesystem, path)
		if err != nil {
			return err
		}
		result = append(result, *cfg)
		return nil
	})
	return result, err
}
