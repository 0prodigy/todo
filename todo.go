package main

import (
	"fmt"
)

const (
	ADD_TODO  = "add"
	LIST_TODO = "list"
	DONE_TODO = "done"
)

type TodoList struct {
	Todo []Todo
}

type Todo struct {
	id   int
	text string
	done bool
}

var todoList TodoList

func addTodo(id int, text string, is_done bool) bool {
	_ = append(todoList.Todo, Todo{
		id:   len(todoList.Todo) + 1,
		text: text,
		done: is_done,
	})

	return true
}

func main() {
	fmt.Println("TODO: implement todo.go")

	tasks := TodoList{
		Todo: []Todo{
			{
				id:   1,
				text: "Learn Go",
				done: false,
			},
			{
				id:   2,
				text: "Learn Rust",
				done: false,
			},
			{
				id:   3,
				text: "Learn Javascript",
				done: false,
			},
		},
	}

	for _, task := range tasks.Todo {
		addTodo(task.id, task.text, task.done)
	}
	fmt.Print("Successfully added todo")
	fmt.Print(todoList)
}
