package api

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/andresKillem/pichincha-devops/todo/structures"
	. "github.com/andresKillem/pichincha-devops/todo/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoRestConfigurator struct {
	handler TodoOperations
}

func NewTodoRestConfigurator(handler TodoOperations) {
	todoOperations := TodoRestConfigurator{handler}
	router := gin.Default()
	router.POST("/DevOps", todoOperations.createTodo)
	router.Run(":3000")
}


func (config TodoRestConfigurator) createTodo(c *gin.Context) {
	var request TodoRequest
	if err := c.ShouldBind(&request); err == nil {

		err := config.verifyApiKey(request, c, err)
		if err == nil{
			config.createResponse(request, c)
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (config TodoRestConfigurator) createResponse(request TodoRequest, c *gin.Context) {
	var todo *MicroMessage
	todo = config.handler.Create(TodoPostToBusinessTodo(request))
	responseMessage := TodoResponse{}
	responseMessage.Message = fmt.Sprintf("Hello %s your message will be send", todo.To)
	b, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, string(b))
}

func (config TodoRestConfigurator) verifyApiKey(request TodoRequest, c *gin.Context, err error) error{
	fmt.Println(request.From)
	auth := c.GetHeader("X-Parse-REST-API-Key")
	if auth != "2f5ae96c-b558-4c7b-a590-a501ae1c3f6c" {
		c.String(http.StatusProxyAuthRequired, "unauthorized")
		return errors.New("unauthorized")
	}
	fmt.Println("With key"+ auth)
	return nil

}


type TodoRequest struct {
	Message string    `json:"message"`
	To      string    `json:"to"`
	From    string    `json:"from"`
	TimeToLifeSec int `json:"timeToLifeSec"`
}

type TodoResponse struct {
	Message string    `json:"message"`
}


func TodoPostToBusinessTodo(request TodoRequest) MicroMessage {
	return MicroMessage{Message: request.Message, To: request.To, From: request.From, TimeToLifeSec: request.TimeToLifeSec}
}
