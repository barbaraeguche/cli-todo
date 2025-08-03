package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var toggleIndex int
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Mark/Unmark a task as completed",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := GetTodos().Complete(toggleIndex)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while toggling a task: %s\n", err)
		}
	},
}

func init() {
	toggleCmd.Flags().IntVarP(&toggleIndex, "index", "i", 4, "index of task being toggled")
	rootCmd.AddCommand(toggleCmd)
}
