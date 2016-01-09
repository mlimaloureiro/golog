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
func (tasks *Tasks) AddItem(item Task) []Task {
	tasks.Items = append(tasks.Items, item)
	return tasks.Items
}
