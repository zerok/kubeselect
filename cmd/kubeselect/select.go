package main

import (
	"context"
	"fmt"
	"strconv"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gitlab.com/zerok/kubeselect/pkg/kubectlconfig"
)

type option struct {
	file    string
	context string
}

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select the environment to be used with kubectl",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		fs := afero.NewOsFs()
		dotkube, _ := cmd.Flags().GetString("dotkube-dir")
		files, err := kubectlconfig.FindAllFiles(ctx, fs, dotkube)
		if err != nil {
			logger.Fatal().Err(err).Msgf("Failed to load kubectl config files.")
		}
		options := make([]option, 0, 10)
		for _, file := range files {
			for _, c := range file.Contexts {
				options = append(options, option{file: file.Path, context: c.Name})
			}
		}
		selection, err := fuzzyfinder.Find(options, func(i int) string {
			return fmt.Sprintf("%s @ %s", options[i].context, options[i].context)
		})
		if err != nil {
			logger.Fatal().Err(err).Msg("Selection failed")
		}
		if selection > len(options)-1 || selection < 0 {
			logger.Fatal().Err(err).Msgf("Invalid selection: %d", selection)
		}
		opt := options[selection]
		fmt.Printf("export KUBECONFIG=%s\nexport KUBECTX=%s\n", strconv.Quote(opt.file), strconv.Quote(opt.context))
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
