package main

import (
	"log"
	"os"

	"github.com/radding/ArborGo/internal/commands"
	"github.com/urfave/cli"
)

func main() {
	cli.AppHelpTemplate = `The Arbor programming language tool chain!

Usage:
	arbor [--version] [--help] <subcommand> [<args>]

Subcommands:
	build:
		Usage:
			arbor build [files] [--o|--output outputname]
		Flags:
			-o | --output 		The output of the file
`
	app := cli.NewApp()
	app.Flags = []cli.Flag{}
	app.Action = func(c *cli.Context) error {
		subCmd := "help"
		if c.NArg() > 0 {
			subCmd = c.Args().Get(0)
		}
		slice := os.Args[1:]
		if subCmd == "help" {
			return nil
		}
		if len(slice) > 1 {
			slice = slice[1:]
		}
		return commands.Exec(subCmd, slice)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
