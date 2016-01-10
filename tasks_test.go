package main

import "testing"

func TestAddItem(t *testing.T) {
	tasks := Tasks{}
	tasks.addItem(Task{"identifier-1", "start", "2016-01-02T15:04:00Z"})
	tasks.addItem(Task{"identifier-2", "start", "2016-01-02T15:04:00Z"})
	if !(len(tasks.Items) == 2) {
		t.Errorf("Expected 2 tasks, only %d.", len(tasks.Items))
	}
}

func TestItemToArrayString(t *testing.T) {
	task := Task{"identifier-1", "start", "2016-01-02T15:04:00Z"}
	expectedArrayString := []string{"identifier-1", "start", "2016-01-02T15:04:00Z"}
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

func TestGetByIdentifier(t *testing.T) {
	tasks := Tasks{
		Items: []Task{
			{"identifier-1", "start", "2016-01-02T15:04:00Z"},
			{"identifier-2", "start", "2016-01-02T15:04:00Z"},
			{"identifier-2", "stop", "2016-01-02T15:04:00Z"},
		},
	}

	identifierTasks := tasks.getByIdentifier("identifier-2")
	if !(len(identifierTasks.Items) == 2) {
		t.Errorf("Expected 2 tasks with identifier-2, got %d.", len(identifierTasks.Items))
	}
}

func TestGetIdentifier(t *testing.T) {
	task := Task{"identifier-1", "start", "2016-01-02T15:04:00Z"}
	if task.getIdentifier() != "identifier-1" {
		t.Errorf("Expected identifier-1, got %s.", task.getIdentifier())
	}
}

func TestGetAction(t *testing.T) {
	task := Task{"identifier-1", "start", "2016-01-02T15:04:00Z"}
	if task.getAction() != "start" {
		t.Errorf("Expected start, got %s.", task.getAction())
	}
}

func TestGetAt(t *testing.T) {
	task := Task{"identifier-1", "start", "2016-01-02T15:04:00Z"}
	if task.getAt() != "2016-01-02T15:04:00Z" {
		t.Errorf("Expected 2016-01-02T15:04:00Z, got %s.", task.getAt())
	}
}
