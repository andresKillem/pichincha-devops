package main

import (
	"github.com/andresKillem/pichincha-devops/api"
	"github.com/andresKillem/pichincha-devops/todo/use_cases"
)

func main() {
	todoOperations := use_cases.NewTodoOperationsHandler()
	api.NewTodoRestConfigurator(todoOperations)
}