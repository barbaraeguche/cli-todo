package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Long:  "Add creates a new task and stores it in your task list. e.g: task add \"buy groceries\"",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.TrimSpace(args[0])
		if task == "" {
			_, _ = fmt.Fprintf(os.Stderr, "task cannot be empty\n")
			return
		}

		err := GetTodos().Add(task)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while adding task: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
