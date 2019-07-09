
package main

import (
	. "github.com/andresKillem/pichincha-devops/todo/structures"
	. "github.com/andresKillem/pichincha-devops/todo/use_cases"
)

func main() {
	todoOperations := NewTodoOperationsHandler()

	todoOperations.Create(MicroMessage{})
}

