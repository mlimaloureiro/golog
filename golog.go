package golog

import (
	"fmt"
	"os"
	"regexp"

	"github.com/codegangsta/cli"
)

const alphanumericRegex = "^[a-zA-Z0-9_]*$"

var commands = []cli.Command{
	{
		Name:   "start",
		Usage:  "start tracking a given task",
		Action: Start,
	},
	{
		Name:   "stop",
		Usage:  "Stop tracking a given task",
		Action: Stop,
	},
	{
		Name:   "status",
		Usage:  "Give status of all tasks",
		Action: Status,
	},
	{
		Name:   "list",
		Usage:  "List all tasks",
		Action: List,
	},
}

// Start a given task
func Start(context *cli.Context) {
	if !IsValidIdentifier(context.Args().First()) {
		cli.ShowCommandHelp(context, context.Command.FullName())
	}
	fmt.Println("starting", context.Args().First())
}

// Pause a given task
func Stop(context *cli.Context) {
	if !IsValidIdentifier(context.Args().First()) {
		cli.ShowCommandHelp(context, context.Command.FullName())
	}
	fmt.Println("pause", context.Args().First())
}

// Status display tasks being tracked
func Status(context *cli.Context) {
	fmt.Print("status")
}

// List lists all tasks
func List(context *cli.Context) {
	fmt.Print("list")
}

// IsValidIdentifier checks if the string passed is a valid task identifier
func IsValidIdentifier(identifier string) bool {
	re := regexp.MustCompile(alphanumericRegex)
	return len(identifier) > 0 && re.MatchString(identifier)
}

func main() {
	app := cli.NewApp()
	app.Name = "Golog"
	app.Usage = "Easy CLI time tracker for your tasks"
	app.Version = "0.1"
	app.Commands = commands

	app.Run(os.Args)
}
