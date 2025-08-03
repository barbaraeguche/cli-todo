package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current tasks",
	Long:  "List displays all tasks in your task list with their status.  e.g: task list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		GetTodos().List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
