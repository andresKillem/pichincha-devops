package use_cases

import (
	. "github.com/andresKillem/pichincha-devops/todo/structures"
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	var operationHandler TodoOperations
	g.Describe("MicroMessage CRUD use cases", func() {

		// Runs at the beginning of all tests
		g.Before(func() {
			operationHandler = NewTodoOperationsHandler()
		})

		// Runs before each test
		g.BeforeEach(func() {
			operationHandler = NewTodoOperationsHandler()
		})

		// Runs after each test
		g.AfterEach(func() {
			operationHandler = nil
		})

		// Runs after all tests
		g.After(func() {
			operationHandler = nil
		})

		// Passing Tests
		g.It("Should create a Message ", func() {
			todo := MicroMessage{Message: "1"}
			result := operationHandler.Create(todo)
			g.Assert(todo.Message).Equal(result.Message)
		})
	})
}
