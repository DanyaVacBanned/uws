package internal

import (
	"context"
	"fmt"

	"github.com/urfave/cli"
)

// commands definiton
const (
	Add     = "add"
	Execute = "execute"
)

// flags definition
const (
	Script  = "script"
	Command = "command"
)

func DeclareCommands() []cli.Command {
	var commands []cli.Command
	addCommand := &cli.Command{
		Name: Add,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: Script, Usage: "Adds a script. Usage: uws add script `script_name`"},
			&cli.StringFlag{Name: Command, Usage: "Adds a command, Usage: uws add command `command_name`"},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			fmt.Printf("Hello %[1]\n", Script)
			return nil
		},
	}
	commands = append(commands, *addCommand)
	return commands
}
