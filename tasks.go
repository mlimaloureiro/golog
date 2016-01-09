package golog

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

func (task Task) toArrayString() []string {
	return []string{task.Identifier, task.Action, task.At}
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
