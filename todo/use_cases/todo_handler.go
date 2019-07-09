package use_cases

import (
	. "github.com/bixlabs/go-layout/todo/structures"
)

type TodoOperations interface {
	Create(todo Todo) *Todo
	Read(id string) *Todo
	Update(todo Todo) *Todo
	Delete(id string) bool
}

type TodoOperationsHandler struct{}

func NewTodoOperationsHandler() TodoOperationsHandler {
	return TodoOperationsHandler{}
}

func (handler TodoOperationsHandler) Create(todo Todo) *Todo {
	println("A Todo was created")
	return &todo
}

func (handler TodoOperationsHandler) Read(id string) *Todo {
	println("A Todo was retrieved")
	return &Todo{ID:id}
}

func (handler TodoOperationsHandler) Update(todo Todo) *Todo {
	println("A Todo was updated")
	return &todo
}

func (handler TodoOperationsHandler) Delete(id string) bool {
	println("A Todo was deleted")
	return true
}
