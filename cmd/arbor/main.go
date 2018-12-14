package main

import (
	"fmt"
	"github.com/radding/ArborGo/lib/plugins"
	"log"
	"os"
	"path"
	"path/filepath"
	"plugin"

	"github.com/urfave/cli"
)

// LoadPlugins load the plugins for the tool chain
func LoadPlugins() []plugins.Command {
	plugs := []plugins.Command{}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dir = path.Join(dir, "plugins")
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			plug, err := plugin.Open(path)
			if err != nil {
				return err
			}
			cmdPlug, err := plug.Lookup("Command")
			if err != nil {
				return err
			}
			cmd, ok := cmdPlug.(plugins.Command)
			if !ok {
				return fmt.Errorf("plugin %s doesn't implement the plugins.Command interface", path)
			}
			plugs = append(plugs, cmd)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return plugs
}

func main() {
	plugs := LoadPlugins()
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

	cmds := []cli.Command{}
	for _, plug := range plugs {
		cmds = append(cmds, cli.Command{
			Name:   plug.GetName(),
			Usage:  plug.Help(),
			Action: plug.Action,
		})
	}
	app.Commands = cmds
	// app.Action = func(c *cli.Context) error {
	// 	subCmd := "help"
	// 	if c.NArg() > 0 {
	// 		subCmd = c.Args().Get(0)
	// 	}
	// 	slice := os.Args[1:]
	// 	if subCmd == "help" {
	// 		return nil
	// 	}
	// 	if len(slice) > 1 {
	// 		slice = slice[1:]
	// 	}
	// 	return commands.Exec(subCmd, slice)
	// }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
