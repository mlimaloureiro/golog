package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// RepositoryInterface interface is used by a RepositoryHandler to do the necessary
// operations. Example implementations of an RepositoryProvider might be a simple
// csv file, sql, mongo...
type RepositoryInterface interface {
	save(task Task) bool
	load() Tasks
}

// TaskCsvRepository is a type with the path of the file to be readed
type TaskCsvRepository struct {
	Path string
}

// FileRepository implementation of RepositoryInterface for simple .csv files
// each line: identifier,action,at
func (csvRepository TaskCsvRepository) load() Tasks {
	csvFile, err := os.Open(csvRepository.Path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	rawCsvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tasks := Tasks{}
	for _, line := range rawCsvData {
		tasks.addItem(Task{Identifier: line[0], Action: line[1], At: line[2]})
	}

	return tasks
}

func (csvRepository TaskCsvRepository) save(task Task) bool {
	csvFile, err := os.OpenFile(csvRepository.Path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Write(task.toArrayString())
	writer.Flush()

	return true
}

func (csvRepository TaskCsvRepository) clear() bool {
	csvFile, err := os.OpenFile(csvRepository.Path, os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	return true
}
