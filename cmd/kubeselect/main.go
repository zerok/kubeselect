package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var logger zerolog.Logger

var rootCmd = &cobra.Command{
	Use: "kubeselect",
	Run: func(cmd *cobra.Command, args []string) {
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	},
}

func init() {
	rootCmd.PersistentFlags().String("dotkube-dir", os.ExpandEnv("$HOME/.kube"), "Path to the .kube directory")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Err(err).Msg("Command failed.")
	}
}
