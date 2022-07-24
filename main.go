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

var todoList = TodoList{}

func addTodo(id int, text string, is_done bool) bool {
	todoList.Todo = append(todoList.Todo, Todo{
		id:   id,
		text: text,
		done: is_done,
	})
	return true
}

func listTodo() {
	for _, todo := range todoList.Todo {
		fmt.Printf("%d : %s is done %t\n", todo.id, todo.text, todo.done)
	}
}

func updateStatus(id int, status bool) bool {
	if id > len(todoList.Todo) || id < 0 {
		return false
	}
	todoList.Todo[id].done = status
	fmt.Printf("Updated status for %d from %t to %t", id, !status, status)
	return true
}

func removeSliceItem(index int) []Todo {
	return append(todoList.Todo[:index], todoList.Todo[index+1:]...)
}

func deleteTodo(id int) bool {
	if id > len(todoList.Todo) || id < 0 {
		return false
	}
	for index, task := range todoList.Todo {
		if task.id == id {
			todoList.Todo = removeSliceItem(index)
			break
		}
	}
	fmt.Println("Deleted successfully")
	return true
}

func main() {

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
	fmt.Println("Added task in todo list")
	listTodo()
	updateStatus(2, true)
	deleteTodo(1)
	listTodo()
}
