package cmd

import (
	"github.com/dreynaldis/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Short: "List All Task",
		Long: `List all task or filter tasks by status
		
		Example:
		task-tracker list todo
		task-tracker list
		task-tracker list done
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) error {
	if len(args) > 0 {
		status := task.TaskStatus(args[0])
		return task.ListTask(status)
	}
	return task.ListTask("all")
}