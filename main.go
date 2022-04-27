package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/bill-greatness/goxide/models/comments"
	"github.com/bill-greatness/goxide/models/posts"
	"github.com/bill-greatness/goxide/models/todos"
	"github.com/bill-greatness/goxide/models/users"
	"github.com/gin-gonic/gin"
)

func main() {

	// customizing the gin Logger

	file, _ := os.Create("goxide-logs.log")

	// disable color format when logging.
	gin.DisableConsoleColor()

	//logs to standard out and file.
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()

	// load HTML Templates
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	// Trying custom log formats with gin.LoggerWithFormatter
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// serve static files.
	router.Static("/static", "./static")
	router.StaticFile("/static", "./static/images/_.png")

	//Default Routes to Page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl.html", gin.H{"title": "Fake JSON Samples Generator"})

	})

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

	comment := router.Group("/comments")
	{
		// Get Random Comments, by default 30 comments, pass a ?total={int} for the total number of comments. max 60.
		comment.GET("/:id", comments.GetComments)

		// pass a query to this path with postID /comments?postID=4
		comment.GET("/post", comments.GetComment)

		comment.POST("/", comments.CreateComment)

		// Delete By Comment ID
		comment.DELETE("/:id", comments.DeleteComment)
	}

	router.Run("localhost:5050")
}
