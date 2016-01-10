package golog

import "testing"

const hourInSeconds = 3600
const hourInMinutes = 60

func TestTransform(t *testing.T) {
	tasks := Tasks{
		Items: []Task{
			// 2 hours 2 mins and 5 seconds for identifier-1
			{"identifier-1", "start", "2016-01-02T15:04:00Z"},
			{"identifier-1", "stop", "2016-01-02T17:04:02Z"},
			{"identifier-1", "start", "2016-12-29T19:04:00Z"},
			{"identifier-1", "stop", "2016-12-29T19:06:02Z"},
			// 1 hour for identifier-2
			{"identifier-2", "start", "2016-01-02T15:04:00Z"},
			{"identifier-2", "stop", "2016-01-02T16:04:00Z"},
		},
	}
	transformer := Transformer{LoadedTasks: tasks}
	transformedTasks := transformer.Transform()
	expectedString1 := "2h:2m:4s  identifier-1"
	expectedString2 := "1h:0m:0s  identifier-2"
	if transformedTasks["identifier-1"] != expectedString1 {
		t.Errorf("Expected %s, got %s.", expectedString1, transformedTasks["identifier-1"])
	}
	if transformedTasks["identifier-2"] != expectedString2 {
		t.Errorf("Expected %s, got %s.", expectedString2, transformedTasks["identifier-2"])
	}
}

func TestSecondsToHuman(t *testing.T) {
	transformer := Transformer{}
	secondsCase1 := 1432
	secondsCase2 := 4432
	if transformer.SecondsToHuman(secondsCase1) != "0h:23m:52s" {
		t.Errorf(
			"1432 Seconds to human should be 0h:23m:52s, got %s.",
			transformer.SecondsToHuman(secondsCase1),
		)
	}
	if transformer.SecondsToHuman(secondsCase2) != "1h:13m:52s" {
		t.Errorf(
			"4432 Seconds to human should be 1h:13m:52s, got %s.",
			transformer.SecondsToHuman(secondsCase2),
		)
	}
}

func TestTrackingToSeconds(t *testing.T) {
	tasks := Tasks{
		Items: []Task{
			// 2 hours 2 mins and 5 seconds for identifier-1
			{"identifier-1", "start", "2016-01-02T15:04:00Z"},
			{"identifier-1", "stop", "2016-01-02T17:04:02Z"},
			{"identifier-1", "start", "2016-12-29T19:04:00Z"},
			{"identifier-1", "stop", "2016-12-29T19:06:02Z"},
			// 1 hour for identifier-2
			{"identifier-2", "start", "2016-01-02T15:04:00Z"},
			{"identifier-2", "stop", "2016-01-02T16:04:00Z"},
			// identifier-1 again to check positions
			{"identifier-1", "start", "2017-01-01T19:06:02Z"},
			{"identifier-1", "stop", "2017-01-01T19:06:03Z"},
		},
	}
	transformer := Transformer{LoadedTasks: tasks}
	if transformer.TrackingToSeconds("identifier-1") != hourInSeconds*2+hourInMinutes*2+5 {
		t.Errorf(
			"Transformation for identifier-1 should be 7325 seconds, got %d.",
			transformer.TrackingToSeconds("identifier-1"),
		)
	}
	if transformer.TrackingToSeconds("identifier-2") != hourInSeconds*1 {
		t.Errorf(
			"Transformation for identifier-1 should be 3600 seconds, got %d.",
			transformer.TrackingToSeconds("identifier-2"),
		)
	}
}

func TestIsActive(t *testing.T) {
	if !isActive(stop) {
		t.Error("When the next entry to look is the stop action it should mean the task is active.")
	}
}
