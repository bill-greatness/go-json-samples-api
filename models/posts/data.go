package posts

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

func GetData(page int) []*Post {
	dir, _ := os.Getwd()
	fileLink := dir + "/data/posts.json"

	fileInfo, err := os.ReadFile(fileLink)

	if err != nil {
		panic(err)
	}

	var posts Posts
	err = json.Unmarshal(fileInfo, &posts)

	if err != nil {
		panic(err)
	}
	Info := []*Post{}

	start, end := getPagination(page)
	for i := start; i < end; i++ {
		Info = append(Info, posts.Posts[i])
	}

	return Info
}
