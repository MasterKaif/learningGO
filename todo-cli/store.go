package main

import(
	"encoding/json";
	"os"
)

const fileName = "tasks.json"

func loadTasks() ([]Tasks, error) {
	var tasks []Tasks

	file, err := os.ReadFile(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}

		return nil, err
	}

	err = json.Unmarshal(file, &tasks)

	return tasks, err
}

func saveTasks(tasks []Tasks) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
