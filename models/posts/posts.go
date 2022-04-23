package posts

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	Title   string    `json:"title" validate:"required"`
	ID      string    `json:"id"`
	UserID  int       `json:"userID" validate:"required"`
	Content string    `json:"content" validate:"required"`
	Date    time.Time `json:"datePosted"`
}

type Posts struct {
	Posts []Post
}

func (post *Post) Jsonize() string {
	data, err := json.MarshalIndent(post, "", " ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (post *Post) Validate() error {
	validator := validator.New()
	return validator.Struct(post)
}
