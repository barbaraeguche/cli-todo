# cli-todo 📋
the classic todo app; built to deepen my understanding of core go concepts.

## tech stack ✨
- [cobra](https://github.com/spf13/cobra) for cli structure
- [aqua security's table](https://github.com/aquasecurity/table) for clean table rendering

## features 👾
- `task add "..."` — add a new task
- `task edit -i=X "..."` — edit a task by index
- `task toggle -i=X` — toggle task completion
- `task delete -i=X` — delete a task by index
- `task list` — view all tasks

`i` refers to the index of the task you want to work with.

## what i learned 💭
- working with pointers and value receivers
- building with cobra
- function and variable scoping in go

## running the project 🏁
- run `go mod init github.com/yourusername/cli-todo` to initialize the module
- run `go build -o task` to rebuild after changes
- `go run .` to run the application

*note: task data is saved and updated automatically in a local json file*