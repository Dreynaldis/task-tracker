package cmd

import (
	"fmt"
	"strconv"

	"github.com/dreynaldis/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		Short: "Delete a task",
		Long: `Delete a task by providing the task ID
		
		Example: 
		task-tracker delete 1
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}
	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) != 1{
		return fmt.Errorf("please provide a task ID")
	}

	taskId := args[0]
	taskIdInt, err := strconv.Atoi(taskId)
	if err != nil {
		return err
	}
	return task.DeleteTask(taskIdInt)
}