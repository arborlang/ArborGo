package main

import (
	"fmt"
	"github.com/radding/ArborGo/internal/environment"
	"github.com/urfave/cli"
	// "log"
	"os"
)

// Run is the entrypoint for building
type Run struct{}

// GetName returns the name of the command
func (b Run) GetName() string { return "run" }

// Category returns the catagory
func (b Run) Category() string { return "Run" }

// Action builds the project
func (b Run) Action(c *cli.Context) {
	file := c.Args()[0]
	var content []byte
	entryPoint := c.String("entrypoint")
	content, isWasm, err := environment.LoadFile(file, c.Bool("wasm"))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if isWasm {
		ret, err := environment.RunWasm(content, entryPoint, "stdlib.so")
		if err != nil {
			fmt.Println(err)
			ret = -1
		}
		os.Exit(int(ret))
	}
	ret, err := environment.RunArbor(content, entryPoint)
	if err != nil {
		fmt.Println(err)
		ret = -1
	}
	os.Exit(int(ret))
}

// Flags returns the Flags for the command
func (b Run) Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "wasm",
			Usage: "The file is already compiled to wasm",
		},
		cli.StringFlag{
			Name:  "entrypoint",
			Usage: "The entrypoint for the application",
			Value: "main",
		},
	}
}

// Help Describe the command
func (b Run) Help() map[string]string {
	mp := map[string]string{}
	mp["description"] = "Run an arbor project"
	mp["usage"] = "arbor run <options> [files]"
	return mp
}

// Command is the plugin
var Command Run

func main() {

}
