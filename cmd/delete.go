package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteIndex int
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := GetTodos().Delete(deleteIndex)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while deleting a task: %s\n", err)
		}
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&deleteIndex, "index", "i", -1, "index of task being deleted")
	rootCmd.AddCommand(deleteCmd)
}
