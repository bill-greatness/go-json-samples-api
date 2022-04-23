package todos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func newID() int {
	return len(Toos) + 1
}

func getID(id string) int {
	newID, err := strconv.Atoi(id)

	// return invalid ID id is not a number.
	if err != nil {
		return -1
	}

	return newID
}

func GetTodos(c *gin.Context) {

	c.JSON(http.StatusOK, Toos)
}

func AddTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := newTodo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newTodo.ID = newID()
	Toos = append(Toos, newTodo)
	c.JSON(http.StatusCreated, Toos)
}

func GetTodo(c *gin.Context) {
	id := getID(c.Param("id"))

	// check ID from getID.
	if id == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id specified, should be <type:int>"})
		return
	}

	// loop through todos to find to with the passed ID
	for _, todo := range Toos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Todo with ID (%d) not found!", id)})
}

func UpdateTodo(c *gin.Context) {
	id := getID(c.Param("id"))
	body := c.Request.Body

	var newBody Todo
	if id == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "todo id specified does not exist"})
		return
	}
	err := json.NewDecoder(body).Decode(&newBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	for idx, todo := range Toos {
		if todo.ID == id {
			todo.ID = id
			Toos[idx] = newBody
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "update successful"})
}

func DeleteTodo(c *gin.Context) {
	// using getID method to convert and check ID
	id := getID(c.Param("id"))

	var cleanedData []Todo
	if id == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "todo id specified does not exist"})
		return
	}

	for _, todo := range Toos {
		if todo.ID != id {
			cleanedData = append(cleanedData, todo)
		}
	}
	c.JSON(http.StatusOK, cleanedData)
}
