package main 

import (
	"encoding/json"
	"fmt"
	"os"
)

const TasksFile = "tasks.json"

func LoadTasks() (TaskList, error) {
	var taskList TaskList	

	if _,err := os.Stat(TasksFile); os.IsNotExist(err) {
		return taskList, nil
	}
	
	data, err := os.ReadFile(TasksFile)
	if err != nil {
		return taskList, fmt.Errorf("error reading tasks file: %w", err)
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &taskList)
		if err != nil {
			return taskList, fmt.Errorf("error parsing task file: %w", err)
		}
	}


	return taskList, nil
}

func SaveTasks(taskList TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("error converting task to json: %w", err)
	}

	err = os.WriteFile(TasksFile, data,0644)
	if err != nil {
		return fmt.Errorf("error writing tasks to file: %w", err)
	}
	return nil
}