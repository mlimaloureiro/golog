package golog

import "testing"

func TestAddItem(t *testing.T) {
	tasks := Tasks{}
	tasks.addItem(Task{"identifier-1", "start", "timestamp"})
	tasks.addItem(Task{"identifier-2", "start", "timestamp"})
	if !(len(tasks.Items) == 2) {
		t.Errorf("Expected 2 tasks, only %d.", len(tasks.Items))
	}
}

func TestItemToArrayString(t *testing.T) {
	task := Task{"identifier-1", "start", "timestamp"}
	expectedArrayString := []string{"identifier-1", "start", "timestamp"}
	toArrayString := task.toArrayString()
	if toArrayString[0] != expectedArrayString[0] {
		t.Errorf("Expected identifier %s, given %s.", expectedArrayString[0], toArrayString[0])
	}
	if toArrayString[1] != expectedArrayString[1] {
		t.Errorf("Expected action %s, given %s.", expectedArrayString[1], toArrayString[1])
	}
	if toArrayString[2] != expectedArrayString[2] {
		t.Errorf("Expected at %s, given %s.", expectedArrayString[2], toArrayString[2])
	}
}
