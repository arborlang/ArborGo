package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// Exec executes a sub command
func Exec(subCmd string, args []string) error {
	cmd := exec.Command(fmt.Sprintf("arbor-%s", subCmd), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
