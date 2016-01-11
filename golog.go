package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/codegangsta/cli"
	"github.com/mitchellh/go-homedir"
)

const alphanumericRegex = "^[a-zA-Z0-9_-]*$"
const dbFile = "~/.golog"

var dbPath, _ = homedir.Expand(dbFile)
var repository = TaskCsvRepository{Path: dbPath}
var transformer = Transformer{}
var commands = []cli.Command{
	{
		Name:         "start",
		Usage:        "Start tracking a given task",
		Action:       Start,
		BashComplete: AutocompleteTasks,
	},
	{
		Name:         "stop",
		Usage:        "Stop tracking a given task",
		Action:       Stop,
		BashComplete: AutocompleteTasks,
	},
	{
		Name:         "status",
		Usage:        "Give status of all tasks",
		Action:       Status,
		BashComplete: AutocompleteTasks,
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
	identifier := context.Args().First()
	if !IsValidIdentifier(identifier) {
		cli.ShowCommandHelp(context, context.Command.FullName())
	}

	repository.save(Task{Identifier: identifier, Action: "stop", At: time.Now().Format(time.RFC3339)})

	fmt.Println("Stopped tracking ", identifier)
}

// Status display tasks being tracked
func Status(context *cli.Context) {
	identifier := context.Args().First()
	if !IsValidIdentifier(identifier) {
		cli.ShowCommandHelp(context, context.Command.FullName())
	}

	transformer.LoadedTasks = repository.load().getByIdentifier(identifier)
	fmt.Println(transformer.Transform()[identifier])
}

// List lists all tasks
func List(context *cli.Context) {
	transformer.LoadedTasks = repository.load()
	for _, task := range transformer.Transform() {
		fmt.Println(task)
	}
}

// AutocompleteTasks loads tasks from repository and show them for completion
func AutocompleteTasks(context *cli.Context) {
	// This will complete if no args are passed
	if len(context.Args()) > 0 {
		return
	}
	transformer.LoadedTasks = repository.load()
	for _, task := range transformer.LoadedTasks.Items {
		fmt.Println(task.getIdentifier())
	}
}

// IsValidIdentifier checks if the string passed is a valid task identifier
func IsValidIdentifier(identifier string) bool {
	re := regexp.MustCompile(alphanumericRegex)
	return len(identifier) > 0 && re.MatchString(identifier)
}

func checkInitialDbFile() {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Create(dbPath)
	}
}

func main() {
	// @todo remove this from here, should be in file repo implementation
	checkInitialDbFile()
	app := cli.NewApp()
	app.Name = "Golog"
	app.Usage = "Easy CLI time tracker for your tasks"
	app.Version = "0.1"
	app.EnableBashCompletion = true
	app.Commands = commands
	app.Run(os.Args)
}
