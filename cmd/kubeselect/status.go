package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gitlab.com/zerok/kubeselect/pkg/kubectlconfig"
)

var statusCmd = &cobra.Command{
	Use: "status",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		fs := afero.NewOsFs()
		env, err := kubectlconfig.LoadEnvironment(ctx, fs, os.Environ())
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to determine environment.")
		}
		dotkube, _ := cmd.Flags().GetString("dotkube-dir")
		prefix := os.ExpandEnv(dotkube) + "/"
		fmt.Printf("%s // %s\n", strings.TrimPrefix(env.KubeConfig, prefix), env.Context)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
