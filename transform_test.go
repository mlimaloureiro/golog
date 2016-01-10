package golog

import "testing"

const secondsInHour = 3600

func TestTrackingToSeconds(t *testing.T) {
	tasks := Tasks{
		Items: []Task{
			// 2 hours for identifier-1
			{"identifier-1", "start", "2016-01-02 15:04:00"},
			{"identifier-1", "stop", "2016-01-02 17:04:00"},
			// 1 hour for identifier-2
			{"identifier-2", "start", "2016-01-02 15:04:00"},
			{"identifier-2", "stop", "2016-01-02 16:04:00"},
		},
	}
	transformer := Transformer{LoadedTasks: tasks}
	if transformer.TrackingToSeconds("identifier-1") != secondsInHour*2 {
		t.Errorf(
			"Transformation for identifier-1 should be 120 seconds, got %d.",
			transformer.TrackingToSeconds("identifier-1"),
		)
	}
}
