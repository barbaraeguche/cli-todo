package cmd

import (
	"fmt"
	"os"

	"github.com/barbaraeguche/cli-todo/todo"
	"github.com/spf13/cobra"
)

var todos *todo.Todos

func SetTodos(t *todo.Todos) {
	todos = t
}
func GetTodos() *todo.Todos {
	return todos
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI to manage your tasks",
	Long: `       .__  .__            __             .___
  ____ |  | |__|         _/  |_  ____   __| _/____
_/ ___\|  | |  |  ______ \   __\/  _ \ / __ |/  _ \
\  \___|  |_|  | /_____/  |  | (  <_> ) /_/ (  <_> )
 \___  >____/__|          |__|  \____/\____ |\____/
     \/                                    \/

A simple and minimal command-line todo application written in Go.

This project serves as my introductory to building command-line
applications, and the Go language as a whole.

This minimal app allows users to add, view, and manages pending
todos they may currently have.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
