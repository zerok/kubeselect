package main

import (
	"context"
	"os"
	"os/exec"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gitlab.com/zerok/kubeselect/pkg/kubectlconfig"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run kubectl with the configured environment",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		fs := afero.NewOsFs()
		env, err := kubectlconfig.LoadEnvironment(ctx, fs, os.Environ())
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to load environment.")
		}
		finalArgs := []string{}
		if env.Context != "" {
			finalArgs = append(finalArgs, "--context", env.Context)
		}
		finalArgs = append(finalArgs, args...)
		p := exec.Command("kubectl", finalArgs...)
		p.Stdout = os.Stdout
		p.Stdin = os.Stdin
		p.Stderr = os.Stderr
		if err := p.Run(); err != nil {
			if eerr, ok := err.(*exec.ExitError); ok {
				os.Exit(eerr.ExitCode())
			}
			logger.Fatal().Err(err).Msg("kubectl exited.")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
