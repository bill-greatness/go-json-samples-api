package comments

import (
	"encoding/json"
	"os"
)

const (
	indexPage  = 1
	chunckSize = 10
)

func getPagination(page int) (start int, end int) {
	if page == 1 {
		return indexPage - 1, chunckSize
	} else {
		startIndex := (indexPage * page) * chunckSize
		endIndex := startIndex + chunckSize
		return startIndex, endIndex
	}

}

func GenerateComments(page int) []*Comment {
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
	start, end := getPagination(page)
	for i := start; i < end; i++ {
		Info = append(Info, comments.Comments[i])
	}

	return Info

}
