package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/arborlang/ArborGo/internal/parser/scope"
	typevisitor "github.com/arborlang/ArborGo/internal/parser/visitors/types"
	umlvisitor "github.com/arborlang/ArborGo/internal/parser/visitors/umlVisitor"
	"github.com/spf13/cobra"
)

var fOut *string

func init() {
	fOut = build.Flags().StringP("out", "o", "arbor_ast.uml", "File to write the uml file to")
	rootCmd.AddCommand(visualize)
}

var visualize = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize the Arbor AST",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Visualizing the Arbor AST")
		// fName := "../../docs/example/example.ab"
		fName := args[0]
		fIn, err := os.Open(fName)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		node, err := rulesv2.Parse(fIn)
		if err != nil {
			log.Fatalln("Failed to parse:", err)
			os.Exit(255)
		}
		f, err := os.OpenFile(*fOut, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0444)

		w := bufio.NewWriter(f)
		defer f.Close()
		if err != nil {
			log.Fatalln("Failed to get file:", err)
			os.Exit(255)
		}
		vs := GetAllVisitors()
		var symTable *scope.SymbolTable
		for _, v := range vs {
			var err error
			node, err = node.Accept(v)
			symTable, _ = typevisitor.GetScope(v)
			if err != nil {
				log.Fatalln("Failed to walk tree", err)
				os.Exit(255)
			}
		}
		fmt.Println("Here I am")
		fmt.Println(symTable)
		node, err = umlvisitor.Visualize(node, w)
		// umlWriter := umlvisitor.New(w)
		// node, err = node.Accept(umlWriter)
		w.Flush()
		if err != nil {
			log.Fatalln("Failed to get UML:", err)
			os.Exit(255)
		}
		// vis, visitor := visualizer.New()
		// fName := args[0]
		// f, err := os.Open(fName)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	os.Exit(1)
		// }
		// log.Println("Visualizing file", fName)
		// node, err := rulesv2.Parse(f)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	os.Exit(1)
		// }
		// _, err = node.Accept(visitor)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	os.Exit(1)
		// }
		// visualizer.NewServer(vis)
	},
}
