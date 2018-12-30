package main

import (
	"fmt"
	"github.com/urfave/cli"
)

// Build is the entrypoint for building
type Build struct{}

// GetName returns the name of the command
func (b Build) GetName() string { return "build" }

// Category returns the catagory
func (b Build) Category() string { return "Build" }

// Action builds the project
func (b Build) Action(c *cli.Context) {
	fmt.Println("I guess I'll build")
}

// Flags returns the Flags for the command
func (b Build) Flags() []cli.Flag {
	return []cli.Flag{}
}

// Help Describe the command
func (b Build) Help() map[string]string {
	mp := map[string]string{}
	mp["description"] = "Build an arbor project into an executable"
	mp["usage"] = "arbor build <options> [files]"
	return mp
}

// Command is the plugin
var Command Build

func main() {

}
