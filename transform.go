package golog

import (
	"fmt"
	"math"
	"os"
	"time"
)

const (
	start = "start"
	stop  = "stop"
)

// Transformer is a type that has loaded all Tasks entries from storage
type Transformer struct {
	LoadedTasks Tasks
}

// SecondsToHuman returns an human readable string from seconds
func (transformer *Transformer) SecondsToHuman(totalSeconds int) string {
	hours := math.Floor(float64(((totalSeconds % 31536000) % 86400) / 3600))
	minutes := math.Floor(float64((((totalSeconds % 31536000) % 86400) % 3600) / 60))
	seconds := (((totalSeconds % 31536000) % 86400) % 3600) % 60

	return fmt.Sprintf("%dh:%dm:%ds", int(hours), int(minutes), int(seconds))
}

// TrackingToSeconds get entries from storage by identifier and calculate
// time between each start/stop for a single identifier
func (transformer *Transformer) TrackingToSeconds(identifier string) int {
	nextAction := "start"
	var durationInSeconds float64
	var startTime, stopTime time.Time

	tasks := transformer.LoadedTasks.getByIdentifier(identifier)
	for _, task := range tasks.Items {
		if task.getAction() == start && nextAction == start {
			nextAction = stop
			startTime = parseTime(task.getAt())
		}
		if task.getAction() == stop && nextAction == stop {
			nextAction = start
			stopTime = parseTime(task.getAt())
			durationInSeconds += stopTime.Sub(startTime).Seconds()
		}
	}

	if isActive(nextAction) {
		durationInSeconds += time.Since(startTime).Seconds()
	}

	return int(durationInSeconds)
}

// we can check if a task is active if we reach the end of the loop
// without finding the last stop action
func isActive(nextAction string) bool {
	return nextAction == stop
}

func parseTime(at string) time.Time {
	then, err := time.Parse(time.RFC3339, at)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return then
}
