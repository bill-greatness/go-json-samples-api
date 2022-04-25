package users

import (
	"encoding/json"

	"io/ioutil"
	"os"
)

func GetData(total int) []*User {
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

	// get users by the total number.
	for count := 0; count < total; count++ {
		Info = append(Info, users.Users[count])
	}

	return Info

}
