package todos

import (
	"encoding/json"
	"os"
)

func GetData(total int) []*Todo {
	dir, _ := os.Getwd()
	fileLink := dir + "/data/todos.json"

	fileInfo, err := os.ReadFile(fileLink)
	if err != nil {
		panic(err)
	}

	var todos Todos
	err = json.Unmarshal(fileInfo, &todos)

	if err != nil {
		panic(err)
	}
	Info := []*Todo{}

	for i := 0; i < total; i++ {
		Info = append(Info, todos.Todos[i])
	}

	return Info
}
