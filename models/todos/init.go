package todos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxCount = 30
	minCount = 10
)

func newID() int {
	return len(GetData(maxCount)) + 1
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
	total, err := strconv.Atoi(c.Query("total"))
	if err != nil {
		// if quest comes with no total query, return a total of 10 Todos
		c.IndentedJSON(http.StatusOK, GetData(minCount))
		return
	}

	c.JSON(http.StatusOK, GetData(total))
}

func AddTodo(c *gin.Context) {
	Toos := GetData(maxCount)
	var newTodo *Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := newTodo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newTodo.ID = newID()
	newTodo.Time = time.Now().String()
	Toos = append(Toos, newTodo)
	c.JSON(http.StatusCreated, Toos)
}

func GetTodo(c *gin.Context) {
	id := getID(c.Param("id"))
	Toos := GetData(maxCount)
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
	Toos := GetData(maxCount)

	var newBody *Todo
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
	Toos := GetData(maxCount)

	var cleanedData []*Todo
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
