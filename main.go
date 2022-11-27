package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to serialize json into a struct we can define a third column in the struct to hold the key inside the json
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// this data structure is completely different then json, even if it looks like this
var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Learn GraphQL", Completed: false},
	{ID: "3", Item: "Learn Rust", Completed: false},
}

// context is the http request context
func getTodos(context *gin.Context) {
	// both lines seem to do the same thing
	context.IndentedJSON(http.StatusOK, todos)
	// context.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	// this is a middleware that parses the request body and binds it to the newTodo struct
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	if todo, err := getTodoById(id); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else {
		context.IndentedJSON(http.StatusOK, todo)
	}
}

func main() {
	// router is the server
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:8080")
}
