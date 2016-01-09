package golog

import (
	"encoding/csv"
	"fmt"
	"os"
)

// RepositoryInterface interface is used by a RepositoryHandler to do the necessary
// operations. Example implementations of an RepositoryProvider might be a simple
// csv file, sql, mongo...
type RepositoryInterface interface {
	save(tasks Tasks) bool
	load() Tasks
}

// TaskCsv is a type with the path of the file to be readed
type TaskCsv struct {
	Path string
}

// FileRepository implementation of RepositoryInterface for simple .csv files
// each line: identifier,action,at
func (csvPath TaskCsv) load() Tasks {
	csvFile, err := os.Open(csvPath.Path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	rawCsvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tasks := Tasks{}
	// sanity check, display to standard output
	for _, line := range rawCsvData {
		tasks.AddItem(Task{Identifier: line[0], Action: line[1], At: line[2]})
	}
	defer csvFile.Close()

	return tasks
}
