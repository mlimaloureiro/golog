package main

// Tasks Struct with tasks items
type Tasks struct {
	Items []Task
}

// Task define the data structure for
type Task struct {
	Identifier string
	Action     string
	At         string
}

// AddItem add a new task to Tasks items field
func (tasks *Tasks) addItem(item Task) []Task {
	tasks.Items = append(tasks.Items, item)
	return tasks.Items
}

func (tasks Tasks) getByIdentifier(identifier string) Tasks {
	tasksWithIdentifier := Tasks{}
	for _, task := range tasks.Items {
		if task.getIdentifier() != identifier {
			continue
		}
		tasksWithIdentifier.addItem(task)
	}

	return tasksWithIdentifier
}

func (task Task) getIdentifier() string {
	return task.Identifier
}

func (task Task) getAction() string {
	return task.Action
}

func (task Task) getAt() string {
	return task.At
}

func (task Task) toArrayString() []string {
	return []string{task.Identifier, task.Action, task.At}
}
