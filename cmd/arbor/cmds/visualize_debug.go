package cmd

import (
	"log"
	"os"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/arborlang/ArborGo/internal/parser/visitors/visualizer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(visualize)
}

var visualize = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize the Arbor AST",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Visualizing the Arbor AST")
		vis, visitor := visualizer.New()
		fName := args[0]
		f, err := os.Open(fName)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		log.Println("Visualizing file", fName)
		node, err := rulesv2.Parse(f)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		_, err = node.Accept(visitor)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		visualizer.NewServer(vis)
	},
}
