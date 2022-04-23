package todos

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type Todos struct {
	Todos []*Todo `json:"todos"`
}

type Todo struct {
	ID       int    `json:"id" `
	UserID   int    `json:"userID" validate:"required"`
	Activity string `json:"activity" validate:"required"`
	Time     string `json:"time"`
	Date     string `json:"date"`
	Status   string `json:"status"`
}

func (todo *Todo) Jsonize() string {
	info, err := json.MarshalIndent(todo, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(info)
}

func (todo *Todo) Validate() error {

	validator := validator.New()
	return validator.Struct(todo)
}
