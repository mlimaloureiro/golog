package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/codegangsta/cli"
)

const alphanumericRegex = "^[a-zA-Z0-9_]*$"

var repository = TaskCsvRepository{Path: "db.csv"}
var commands = []cli.Command{
	{
		Name:   "start",
		Usage:  "Start tracking a given task",
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
	identifier := context.Args().First()
	if !IsValidIdentifier(identifier) {
		cli.ShowCommandHelp(context, context.Command.FullName())
	}

	repository.save(Task{Identifier: identifier, Action: "start", At: time.Now().Format(time.RFC3339)})

	fmt.Println("Started tracking ", identifier)
}

// Stop a given task
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
