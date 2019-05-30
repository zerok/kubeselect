package kubectlconfig

import (
	"context"
	"os"
	"strings"

	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v2"
)

func FindAllFiles(ctx context.Context, filesystem afero.Fs, root string) ([]ConfigFile, error) {
	result := make([]ConfigFile, 0, 10)
	err := afero.Walk(filesystem, root, func(path string, info os.FileInfo, err error) error {
		if !(strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".conf")) {
			return nil
		}
		fp, err := filesystem.Open(path)
		if err != nil {
			return err
		}
		defer fp.Close()
		cfg := ConfigFile{Path: path}
		if err := yaml.NewDecoder(fp).Decode(&cfg); err != nil {
			return err
		}
		result = append(result, cfg)
		return nil
	})
	return result, err
}
