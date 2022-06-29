package users

import (
	"encoding/json"

	"io/ioutil"
	"os"
)

const (
	chunckSize = 10
	indexPage  = 1
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

func GetData(page int) []*User {

	dir, _ := os.Getwd()
	fileLink := dir + "/data/users.json"

	// Unmarshall users to struct.
	// convert json to Array of Bytes.

	info, err := ioutil.ReadFile(fileLink)
	if err != nil {
		panic(err)
	}

	var users Users

	err = json.Unmarshal([]byte(info), &users)
	if err != nil {
		panic(err)
	}

	Info := []*User{}
	// substitute total for page.

	// get users by the total number.
	start, end := getPagination(page)
	for count := start; count < end; count++ {
		Info = append(Info, users.Users[count])
	}

	return Info

}
