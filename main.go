package main

import (
	"fmt"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" task-cli add \"Task description\"")
	fmt.Println(" task-cli update <id> \"Update desciprtin\"")
	fmt.Println(" task-cli delete <id>")
	fmt.Println(" task-cli mark-in-progress <id>")
	fmt.Println(" task-cli mark-done <id>")
	fmt.Println(" task-cli list [all|todo|in-progress|done]")
}

func main (){
	if(len(os.Args) < 2) {
		printUsage()
		return
	}

	command := os.Args[1]

	taskList, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: %w", err)
		return
	}

    switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: missing task description")
			return
		}
		
		description := os.Args[2]
		task, err := taskList.AddTask(description)
		if err != nil {
			fmt.Printf("Error adding task: %s", err)
			return
		}

		
		err = SaveTasks(taskList)
		if err != nil {
			fmt.Println("Error saving tasks: %w", err)
			return	
		}

		fmt.Printf("Tasks added successfully (ID: %d)", task.ID)

	case "update": 
		if len(os.Args) < 4 {
			fmt.Println("Error: missing task ID or description")
			return 
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error: invalid task ID")
			return 
		}

		description := os.Args[3]
		err = taskList.UpdateTask(id, description)
		if err != nil {
			fmt.Println("Error updating task: %w", err)
			return 
		}

		err = SaveTasks(taskList)
		if err != nil {
			fmt.Println("error saving tasks: %w", err)
			return
		}
		fmt.Printf("Task ID:%d updated successfully\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("error: missing task Id")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error: invalid task ID")
			return
		}
		
		err = taskList.DeleteTask(id)
		if err != nil {
			fmt.Println("error deletting task: %w", err)
			return
		}

		err = SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error saving tasks: %s\n", err)
			return 
		}
		fmt.Printf("Task ID:%d deleted successfully", id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("error: missing task ID")
			return
		}
		
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error: invalid task ID")
			return
		}
		
		err = taskList.ChangeTaskStatus(id, "in-progress")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return 
		}

		err = SaveTasks(taskList)
		if err != nil {
			fmt.Printf("Error saving tasks: %s", err)
			return
		}

		fmt.Printf("Task ID:%d marked as in-progress", id)
	
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("error: missing task ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error: invalid task ID")
			return
		}
		
		err = taskList.ChangeTaskStatus(id, "done")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		err = SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error saving tasks: %s", err)
			return 
		}
		fmt.Printf("Tasks ID:%d marked as done", id)

	case "list":
		status := "all"
		if len(os.Args) >= 3 {
			status = os.Args[2]
			if status != "all" && status != "todo" && status != "in-progress" && status != "done" {
				fmt.Printf("Error: invalid status %s\n", status)
				return
			}
		}

		tasks := taskList.GetTaskByStatus(status)
		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		fmt.Printf("Tasks with status: %s\n", status)
		for _, task := range tasks {
			fmt.Printf("[%d] %s (Status: %s, Created: %s)\n",
		task.ID, 
		task.Description,
		task.Status, 
		task.CreatedAt.Format("2006-01-02 15:04:05"))
		}
	
	default: 
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}	

}