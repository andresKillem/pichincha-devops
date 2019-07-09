package use_cases

import (
	. "github.com/andresKillem/pichincha-devops/todo/structures"
)

type TodoOperations interface {
	Create(todo MicroMessage) *MicroMessage
}

type TodoOperationsHandler struct{}

func NewTodoOperationsHandler() TodoOperationsHandler {
	return TodoOperationsHandler{}
}

func (handler TodoOperationsHandler) Create(todo MicroMessage) *MicroMessage {
	println("A MicroMessage was created")
	return &todo
}
