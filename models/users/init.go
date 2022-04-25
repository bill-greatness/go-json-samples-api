package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	total, err := strconv.Atoi(c.Query("total"))
	// return users with a total of 5 if total query is not passed.
	if err != nil {
		c.IndentedJSON(http.StatusOK, GetData(5))
		return
	}

	// by default, you may get only 200 users per a query, if total is more than 200, w
	if total < 20 {
		c.JSON(http.StatusOK, GetData(total))
		return
	} else {
		c.JSON(http.StatusOK, GetData(20-total))
		return
	}
}

// get a user with a specific ID, options are the id will change on every call since the data is called at random.
func GetUser(c *gin.Context) {
	data := GetData(10)
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	for _, info := range data {
		if info.ID == ID {
			c.IndentedJSON(http.StatusOK, info)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User Not Found!"})

}

func removeItem(id int) ([]*User, error) {
	data := GetData(10)
	cleanedData := []*User{}

	// by design, Ids are positive numbers and by testing,  1 <= len(data) <= 10.
	if id < 1 || id > len(data) {
		return nil, errors.New("invalid id passed")
	}

	for _, usr := range data {
		if usr.ID != id {
			cleanedData = append(cleanedData, usr)
		}
	}

	return cleanedData, nil
}

func patchUser(id int, body *User) ([]*User, error) {
	data := GetData(10)
	if id < 1 || id > len(data) {
		return nil, errors.New("invalid id passed")
	}
	for idx, usr := range data {
		if usr.ID == id {
			data[idx] = body
			data[idx].ID = id
			return data, nil
		}
	}
	return nil, errors.New("something went wrong")

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

	info, err := removeItem(ID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID passed"})
		return
	}
	c.IndentedJSON(http.StatusOK, info)

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	body := c.Request.Body

	defer body.Close()

	var newBody User
	err := json.NewDecoder(body).Decode(&newBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})

		return
	}

	ID, _ := strconv.Atoi(id)
	info, err := patchUser(ID, &newBody)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, info)
}

func AddUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
