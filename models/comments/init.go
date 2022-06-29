package comments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	Comments   = GenerateComments(1)
)

func GetComments(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.IndentedJSON(http.StatusOK, GenerateComments(1))
		return
	}


	c.IndentedJSON(http.StatusOK, GenerateComments(page))
}

func CreateComment(c *gin.Context) {
	body := c.Request.Body
	var comment *Comment

	err := json.NewDecoder(body).Decode(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// validate comment with validator.
	err = comment.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	Comments = append(Comments, comment)
	c.JSON(http.StatusCreated, Comments)
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Delete for empty ID"})
		return
	}
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse ID"})
		return
	}
	var cleanComments []*Comment
	for _, c := range Comments {
		if c.ID != parsedID {
			cleanComments = append(cleanComments, c)
		}
	}
	c.IndentedJSON(http.StatusOK, cleanComments)
}

func GetComment(c *gin.Context) {
	postID := c.Query("postID")

	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Empty postID"})
		return
	}
	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: parsing postID"})
		return
	}
	comments := []Comment{}
	for _, c := range Comments {
		if c.PostID == id {
			comments = append(comments, *c)
		}
	}
	if len(comments) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("couldn't get comments for postID %d", id)})
		return
	}
	c.IndentedJSON(http.StatusOK, comments)
}
