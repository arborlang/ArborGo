package cmd

import (
	"fmt"
	"os"

	"github.com/arborlang/ArborGo/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "arbor",
	Short:   "Arbor is a modern programming langauge for the web",
	Long:    `Arbor is a hyper portable and fast modern langauge for the web`,
	Version: config.Version,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
