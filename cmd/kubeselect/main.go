package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var logger zerolog.Logger
var (
	version = "dev"
	commit  = ""
	date    = ""
)

var rootCmd = &cobra.Command{
	Use: "kubeselect",
	Run: func(cmd *cobra.Command, args []string) {
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
		if v, _ := cmd.Flags().GetBool("version"); v {
			printVersion()
			os.Exit(0)
		}
	},
}

func printVersion() {
	output := strings.Builder{}
	output.WriteString(version)
	if commit != "" || date != "" {
		output.WriteString("  (")
		if commit != "" {
			output.WriteString("commit: " + commit)
			if date != "" {
				output.WriteString(", ")
			}
		}
		if date != "" {
			output.WriteString(fmt.Sprintf("built at %s", date))
		}
		output.WriteString(")")
	}
	fmt.Print(output.String())
}

func init() {
	rootCmd.PersistentFlags().String("dotkube-dir", os.ExpandEnv("$HOME/.kube"), "Path to the .kube directory")
	rootCmd.PersistentFlags().Bool("version", false, "Print version information")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Err(err).Msg("Command failed.")
	}
}
