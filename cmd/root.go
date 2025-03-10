package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command{
cmd:= &cobra.Command{
	Use: "task-tracker",
	Short: "Task Tracker is a CLI tool for managing tasks",
	Long: `Task tracker is a CLI tool for managing. It allows user to add, list and delete tasks.
	
	you can also mark task by its status.
	`,
	}
	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewUpdateTask())
	cmd.AddCommand(NewStatusDoneCmd())
	cmd.AddCommand(NewStatusInProgressCmd())
	cmd.AddCommand(NewStatusTodoCmd())

	return cmd
}

