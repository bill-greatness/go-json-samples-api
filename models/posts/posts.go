package posts

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	Title   string `json:"title" validate:"required"`
	ID      int    `json:"postID"`
	UserID  int    `json:"userID" validate:"required"`
	Content string `json:"content" validate:"required"`
	// some some reasons, instead of time.Time, using string.
	Date string `json:"datePosted"`
}

type Posts struct {
	Posts []*Post `json:"posts"`
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
