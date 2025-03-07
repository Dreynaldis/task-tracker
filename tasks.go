package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func NewTask(id int, description string) Task {
	now := time.Now()
	return Task {
		ID: id,
		Description: description,
		Status: "pending",
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (tl *TaskList) GetNextID() int {
	if len(tl.Tasks) == 0 {
		return 1
	}

	maxID:= 0

	for _, task := range tl.Tasks {
		if task.ID > maxID {
			maxID = task.ID
			}
		}
		return maxID + 1	
}

func (tl *TaskList) AddTask(description string) (Task, error) {
	if description == "" {
		return Task{}, fmt.Errorf("description cannot be empty")
	}

	nextID := tl.GetNextID()
	task := NewTask(nextID, description)
	tl.Tasks = append(tl.Tasks, task)

	return task, nil
}

func (tl *TaskList) UpdateTask(id int, description string) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Description = description
			tl.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("Task with ID %d not found", id)
}

func (tl *TaskList) DeleteTask(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i] = tl.Tasks[len(tl.Tasks)-1]
			tl.Tasks = tl.Tasks[:len(tl.Tasks)-1]
			return nil
		}
	}
	return fmt.Errorf("Task with ID %d not found", id)
}

func (tl *TaskList) ChangeTaskStatus(id int, status string) error {
	validStatuses := map[string]bool {
		"todo" : true,
		"in-progress": true,
		"done": true,
	}

	if !validStatuses[status] {
		return fmt.Errorf("invalid status: %s", status)
	}

	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Status = status
			tl.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}

	return fmt.Errorf("Task with ID %d not found", id)
}

func (tl *TaskList) GetTaskByStatus(status string) []Task {
	var filteredTasks []Task

	for _, task := range tl.Tasks {
		if status == "all" || task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}