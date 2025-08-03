package cmd

import (
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
	Use:   "cli-todo",
	Short: "CLI Todo",
	Long: `A simple and minimal command-line todo application written in Go.

This project serves as an introduction to building CLI tools using the Cobra library.
It allows users to add, view, and manage tasks directly from the terminal.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
