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
		Name:   "clear",
		Usage:  "Clear all data",
		Action: Clear,
	},
	{
		Name:   "list",
		Usage:  "List all tasks",
		Action: List,
	},
}

// Start a given task
func Start(context *cli.Context) error {
	identifier := context.Args().First()
	if !IsValidIdentifier(identifier) {
		return invalidIdentifier(identifier)
	}

	err := repository.save(Task{Identifier: identifier, Action: "start", At: time.Now().Format(time.RFC3339)})

	if err == nil {
		fmt.Println("Started tracking ", identifier)
	}
	return err
}

// Stop a given task
func Stop(context *cli.Context) error {
	identifier := context.Args().First()
	if !IsValidIdentifier(identifier) {
		return invalidIdentifier(identifier)
	}

	err := repository.save(Task{Identifier: identifier, Action: "stop", At: time.Now().Format(time.RFC3339)})

	if err == nil {
		fmt.Println("Stopped tracking ", identifier)
	}
	return err
}

// Status display tasks being tracked
func Status(context *cli.Context) error {
	identifier := context.Args().First()
	if !IsValidIdentifier(identifier) {
		return invalidIdentifier(identifier)
	}

	tasks, err := repository.load()
	if err != nil {
		return err
	}
	transformer.LoadedTasks = tasks.getByIdentifier(identifier)
	fmt.Println(transformer.Transform()[identifier])
	return nil
}

// List lists all tasks
func List(context *cli.Context) error {
	var err error
	transformer.LoadedTasks, err = repository.load()
	if err != nil {
		return err
	}

	for _, task := range transformer.Transform() {
		fmt.Println(task)
	}
	return nil
}

// Clear all data
func Clear(context *cli.Context) error {
	err := repository.clear()
	if err == nil {
		fmt.Println("All tasks deleted")
	}
	return err
}

// AutocompleteTasks loads tasks from repository and show them for completion
func AutocompleteTasks(context *cli.Context) {
	var err error
	transformer.LoadedTasks, err = repository.load()
	// This will complete if no args are passed
	//   or there is problem with tasks repo
	if len(context.Args()) > 0 || err != nil {
		return
	}

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
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func invalidIdentifier(identifier string) error {
	return fmt.Errorf("identifier %q is invalid", identifier)
}
