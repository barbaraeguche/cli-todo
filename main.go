package main

import (
	"fmt"
	"os"

	"github.com/barbaraeguche/cli-todo/cmd"
	st "github.com/barbaraeguche/cli-todo/storage"
	"github.com/barbaraeguche/cli-todo/todo"
)

func main() {
	todos := todo.Todos{}
	storage := st.New[todo.Todos]("todos.json")

	err := storage.Load(&todos)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading todos: %s", err)
		return
	}

	cmd.SetTodos(&todos)
	cmd.Execute()

	err = storage.Store(todos)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error storing todos: %s", err)
		return
	}
}
