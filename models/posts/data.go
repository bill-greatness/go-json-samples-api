package posts

import (
	"encoding/json"
	"os"
)

func GetData(total int) []*Post {
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

	for i := 0; i < total; i++ {
		Info = append(Info, posts.Posts[i])
	}

	return Info
}
