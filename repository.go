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
	save(task Task) error
	load() (Tasks, error)
	clear() error
}

// TaskCsvRepository is a type with the path of the file to be readed
type TaskCsvRepository struct {
	Path string
}

// FileRepository implementation of RepositoryInterface for simple .csv files
// each line: identifier,action,at
func (csvRepository TaskCsvRepository) load() (Tasks, error) {
	tasks := Tasks{}
	csvFile, err := os.Open(csvRepository.Path)
	if err != nil {
		return tasks, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	rawCsvData, err := reader.ReadAll()
	if err != nil {
		return tasks, err
	}

	for _, line := range rawCsvData {
		if len(line) != 3 {
			if err != nil {
				return tasks, fmt.Errorf("csvfile: malformed line: %q", line)
			}
		}
		tasks.addItem(Task{Identifier: line[0], Action: line[1], At: line[2]})
	}

	return tasks, nil
}

func (csvRepository TaskCsvRepository) save(task Task) error {
	csvFile, err := os.OpenFile(csvRepository.Path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	err = writer.Write(task.toArrayString())
	writer.Flush()
	return err
}

func (csvRepository TaskCsvRepository) clear() error {
	csvFile, err := os.OpenFile(csvRepository.Path, os.O_TRUNC|os.O_WRONLY, 0600)
	if err == nil {
		csvFile.Close()
	}
	return err
}
