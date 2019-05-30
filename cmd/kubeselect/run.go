package main

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run kubectl with the configured environment",
	Run: func(cmd *cobra.Command, args []string) {
		c := os.Getenv("KUBECTX")
		finalArgs := []string{"--context", c}
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
