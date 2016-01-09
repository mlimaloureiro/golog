package main

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
		Name:   "pause",
		Usage:  "Pause tracking a given task",
		Action: Pause,
	},
	{
		Name:   "continue",
		Usage:  "Continue tracking a given task",
		Action: Continue,
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
	{
		Name:   "delete",
		Usage:  "Delete a task",
		Action: Delete,
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
func Pause(context *cli.Context) {
	if !IsValidIdentifier(context.Args().First()) {
		cli.ShowCommandHelp(context, context.Command.FullName())
	}
	fmt.Println("pause", context.Args().First())
}

// Continue continue to track paused task
func Continue(context *cli.Context) {
	if !IsValidIdentifier(context.Args().First()) {
		cli.ShowCommandHelp(context, context.Command.FullName())
	}
	fmt.Println("continuing", context.Args().First())
}

// Status display tasks being tracked
func Status(context *cli.Context) {
	fmt.Print("status")
}

// List lists all tasks
func List(context *cli.Context) {
	fmt.Print("list")
}

// Delete a task
func Delete(context *cli.Context) {
	fmt.Print("delete")
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
