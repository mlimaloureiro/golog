package golog

import (
	"fmt"
	"os"
	"time"
)

const timeFormat = "2006-01-02 15:04:00"

// Transformer is a type that has loaded all Tasks entries from storage
type Transformer struct {
	LoadedTasks Tasks
}

// TrackingToSeconds get entries from storage by identifier and calculate
// time between each start/stop for a single identifier
func (transformer *Transformer) TrackingToSeconds(identifier string) int {
	nextAction := "start"
	var durationInSeconds float64
	var start, stop time.Time

	tasks := transformer.LoadedTasks.getByIdentifier(identifier)
	for _, task := range tasks.Items {
		if task.getAction() == "start" && nextAction == "start" {
			nextAction = "stop"
			start = parseTime(task.getAt())
		}
		if task.getAction() == "stop" && nextAction == "stop" {
			nextAction = "start"
			stop = parseTime(task.getAt())
			durationInSeconds += stop.Sub(start).Seconds()
		}
	}

	return int(durationInSeconds)
}

func parseTime(at string) time.Time {
	then, err := time.Parse(timeFormat, at)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return then
}
