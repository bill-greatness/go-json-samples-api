package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	body := c.Request.Body

	var newBody Post
	err := json.NewDecoder(body).Decode(&newBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := newBody.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newBody.Date = time.Now()
	TempPost = append(TempPost, newBody)

	c.JSON(http.StatusOK, TempPost)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var cleanedData []Post

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Delete for ID"})
		return
	}

	for _, post := range TempPost {
		if post.ID != id {
			cleanedData = append(cleanedData, post)
		}
	}
	c.JSON(http.StatusOK, cleanedData)

}

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, TempPost)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	for _, post := range TempPost {
		if post.ID == id {
			c.JSON(http.StatusOK, post)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Post with ID (%s) not found", id)})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	body := c.Request.Body
	var updatedPost Post
	err := json.NewDecoder(body).Decode(&updatedPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	for idx, post := range TempPost {
		if post.ID == id {
			TempPost[idx] = updatedPost
			c.JSON(http.StatusAccepted, TempPost)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("ID (%s) Not Found!", id)})
}
