
package main

import (
	. "github.com/andresKillem/pichincha-devops/todo/use_cases"
	. "github.com/andresKillem/pichincha-devops/todo/structures"
)

func main() {
	todoOperations := NewTodoOperationsHandler()

	todoOperations.Create(Todo{})
	todoOperations.Read("1")
	todoOperations.Update(Todo{})
	todoOperations.Delete("1")
}

