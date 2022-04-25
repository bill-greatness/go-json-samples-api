package comments

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type AllComments struct {
	Comments []*Comment `json:"comments"`
}

type Comment struct {
	ID      int    `json:"id" validate:"required"`
	PostID  int    `json:"postID" validate:"required"`
	UserID  int    `json:"userID" validate:"required"`
	Content string `json:"comment" validate:"required"`
}

func (comment *Comment) Jsonize() string {
	data, err := json.MarshalIndent(comment, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (comment *Comment) Validate() error {
	validator := validator.New()

	return validator.Struct(comment)
}
