package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bill-greatness/goxide/models/posts"
	"github.com/bill-greatness/goxide/models/todos"
	"github.com/bill-greatness/goxide/models/users"
	"github.com/gin-gonic/gin"
)

func getFormat() {
	logger := log.New(io.MultiWriter(os.Stdout), "[goxide] ", log.LUTC)
	headerStrings := "	 Method  |	Path |	IP Address | Time |	Status	\n"
	formatString := fmt.Sprintf("%s\t\t%s \t %d \t %s", "Hello", "World", 1, "Hello")
	fmt.Print(headerStrings)
	logger.Print(formatString)
}

func main() {

	getFormat()
	// customizing the gin Logger

	file, _ := os.Create("goxide-logs.log")

	// disable color format when logging.
	gin.DisableConsoleColor()

	//logs to standard out and file.
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()

	// user router group.
	user := router.Group("/users")
	{
		// Get all users by default return 20 users.
		user.GET("/", users.GetUsers)

		// Get a specific user by ID
		user.GET("/:id", users.GetUser)

		//Add new User

		user.POST("/", users.AddUser)

		// Delete a user by a specific ID
		user.DELETE("/:id", users.DeleteUser)

		// put/update specific user information.
		user.PUT("/:id", users.UpdateUser)
	}

	todo := router.Group("/todos")
	{
		// Get all Todos by default 20, when total paramters is not specificied.
		todo.GET("/", todos.GetTodos)

		// Get todo with a specific ID
		todo.GET("/:id", todos.GetTodo)

		// Add Todo to collection
		todo.POST("/", todos.AddTodo)

		//Delete Todo
		todo.DELETE("/:id", todos.DeleteTodo)

		// Update Todo
		todo.PATCH("/:id", todos.UpdateTodo)
	}

	post := router.Group("/posts")
	{
		// Get All Posts
		post.GET("/", posts.GetPosts)

		// Get a single post by ID
		post.GET("/:id", posts.GetPost)

		// Delete Post by ID
		post.DELETE("/:id", posts.DeletePost)
	}

	router.Run("localhost:5050")
}
