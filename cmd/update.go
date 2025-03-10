package cmd

import (
	"fmt"
	"strconv"

	"github.com/dreynaldis/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func NewUpdateTask() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update",
		Short: "Update a task",
		Long: `Update a task by providing ID and the new description
		
		Exmpale:
		task-tracker update 1 'new description' 
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}
	return cmd
}


func RunUpdateTaskCmd(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("please provide task ID and new description")
	}
	taskID := args[0]
	TaskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		return err
	}

	description := args[1]
	return task.UpdateTaskDescription(TaskIDInt, description)
}

func NewStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mark-done",
		Short: "mark a task status as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.TASK_STATUS_DONE)
		},
	}
	return cmd
}
func NewStatusInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mark-done",
		Short: "mark a task status as in progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.TASK_STATUS_IN_PROGRESS)
		},
	}
	return cmd
}
func NewStatusTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mark-done",
		Short: "mark a task status as todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.TASK_STATUS_TODO)
		},
	}
	return cmd
}
func RunUpdateStatusCmd(args []string, status task.TaskStatus) error {
	if len(args) == 0 {
		return fmt.Errorf("task id is required")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	return task.UpdateTaskStatus(id, status)
}