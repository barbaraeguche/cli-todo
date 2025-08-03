package todo

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Task        string
	IsComplete  bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (ts *Todos) Add(task string) error {
	// append to list
	*ts = append(*ts, Todo{
		Task:      task,
		CreatedAt: time.Now(),
	})
	return nil
}

func (ts *Todos) validateIndex(idx int) error {
	if !(1 <= idx && idx <= len(*ts)) {
		return errors.New("you entered an invalid index")
	}
	return nil
}

func (ts *Todos) Edit(idx int, task string) error {
	todos := *ts

	// validate both index and task
	if err := ts.validateIndex(idx); err != nil {
		return err
	}

	// edit a specific to-do
	todos[idx-1].Task = task
	return nil
}

func (ts *Todos) Complete(idx int) error {
	todos := *ts

	if err := ts.validateIndex(idx); err != nil {
		return err
	}

	idx = idx - 1
	isComplete := todos[idx].IsComplete

	// reset the completion time
	if isComplete {
		todos[idx].CompletedAt = nil
	} else {
		completionTime := time.Now()
		todos[idx].CompletedAt = &completionTime
	}

	// mark as completed
	todos[idx].IsComplete = !isComplete
	return nil
}

func (ts *Todos) Delete(idx int) error {
	todos := *ts

	if err := ts.validateIndex(idx); err != nil {
		return err
	}

	// remove from list
	*ts = append(todos[:idx-1], todos[idx:]...)
	return nil
}

func (ts *Todos) List() {
	t := table.New(os.Stdout)
	t.SetRowLines(false)
	t.SetHeaders("#", "Task", "Completed", "Created At", "Completed At")

	for idx, task := range *ts {
		completed, completedAt := "🌑", ""

		if task.IsComplete {
			completed = "🌕"
			if task.CompletedAt != nil {
				completedAt = task.CompletedAt.Format(time.RFC1123)
			}
		}

		t.AddRow(strconv.Itoa(idx+1), task.Task, completed, task.CreatedAt.Format(time.RFC1123), completedAt)
	}

	t.Render()
}
