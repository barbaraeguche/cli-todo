package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var editIndex int
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a task",
	Long:  "Edit modifies an existing task. e.g: task edit -i=1 \"submit phil paper 2\"",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.TrimSpace(args[0])
		if task == "" {
			_, _ = fmt.Fprintf(os.Stderr, "task cannot be empty\n")
			return
		}

		err := GetTodos().Edit(editIndex, task)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while editing task: %s\n", err)
		}
	},
}

func init() {
	editCmd.Flags().IntVarP(&editIndex, "index", "i", -1, "index of task being edited")
	rootCmd.AddCommand(editCmd)
}
