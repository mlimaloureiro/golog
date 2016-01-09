package golog

import "testing"

func TestCsvRepositoryLoad(t *testing.T) {
	taskCsv := TaskCsv{Path: "test.csv"}
	tasks := taskCsv.load()

	if !(len(tasks.Items) == 2) {
		t.Error("Expected to have 2 items in csv file.")
	}

	expectedTasks := Tasks{}
	expectedTasks.AddItem(Task{Identifier: "track search", Action: "start", At: "2015-08-10"})
	expectedTasks.AddItem(Task{Identifier: "live demo", Action: "stop", At: "2015-09-10"})

	if expectedTasks.Items[0] != tasks.Items[0] || expectedTasks.Items[1] != tasks.Items[1] {
		t.Error("Tasks loaded are different from tasks expected.")
	}
}
