package comments

import (
	"encoding/json"
	"os"
)

const (
	maxCount = 60
	minCount = 30
)

func GenerateComments(total int) []*Comment {
	dir, _ := os.Getwd()
	fileLink := dir + "/data/comments.json"

	fileInfo, err := os.ReadFile(fileLink)

	if err != nil {
		panic(err)
	}
	var comments AllComments

	err = json.Unmarshal([]byte(fileInfo), &comments)
	if err != nil {
		panic(err)
	}

	Info := []*Comment{}
	for i := 0; i < total; i++ {
		Info = append(Info, comments.Comments[i])
	}

	return Info

}
