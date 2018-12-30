package plugins

import (
	"github.com/urfave/cli"
)

// Command Describes the basis of all plugins in the system
type Command interface {
	// GetName returns the name of the Plugin
	GetName() string
	Help() map[string]string
	Action(c *cli.Context)
	Flags() []cli.Flag
	Category() string
}
