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

// Version is the the toolchain version
var Version = "0.0.0-rc0"

func main() {
	plugs := LoadPlugins()
	app := cli.NewApp()
	app.Flags = []cli.Flag{}
	app.Author = "Yoseph Radding"
	app.Description = "The Arbor Language Tool Chain"
	app.Name = "The Tool chain to manage Arbor code"
	app.Version = Version
	cmds := []cli.Command{}
	for _, plug := range plugs {
		cmds = append(cmds, cli.Command{
			Name:        plug.GetName(),
			Action:      plug.Action,
			Description: plug.Help()["description"],
			UsageText:   plug.Help()["usage"],
		})
	}
	app.Commands = cmds
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
