package cmd

import (
	"log"
	"os"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	typevisitor "github.com/arborlang/ArborGo/internal/parser/visitors/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(build)
}

var visitors []ast.Visitor = []ast.Visitor{
	typevisitor.New(),
}

var build = &cobra.Command{
	Use:   "build",
	Short: "Build an Arbor program",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fName := args[0]
		f, err := os.Open(fName)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		node, err := rulesv2.Parse(f)
		if err != nil {
			log.Fatalln(err)
			os.Exit(-1)
		}
		for _, visitor := range visitors {
			node, err = node.Accept(visitor)
			if err != nil {
				log.Fatalln(err)
				os.Exit(1)
			}
		}
	},
}
