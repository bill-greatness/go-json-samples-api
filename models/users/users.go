package users

import (
	"encoding/json"
)

// cordinates for the address field
type CordinateInfo struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// address Information
type AddressInfo struct {
	Street     string         `json:"street"`
	Cordinates *CordinateInfo `json:"cordinates"`
}

// User Model Information.
type User struct {
	ID              int          `json:"id"`
	Name            string       `json:"name"`
	Occupation      string       `json:"occupation"`
	Age             int          `json:"age"`
	CountryOfOrigin string       `json:"countryOfOrigin"`
	PhotoURL        string       `json:"photoURL"`
	Email           string       `json:"email"`
	Gender          string       `json:"gender"`
	MaritalStatus   string       `json:"maritalStatus"`
	DateOfBirth     string       `json:"dateOfBirth"`
	PhoneLine       string       `json:"phoneLine"`
	IsActive        bool         `json:"isActive"`
	Address         *AddressInfo `json:"address"`
}

type Users struct {
	Users []*User `json:"users"`
}

// convert user into json on method call.
func (usr *User) Jsonize() string {
	info, err := json.MarshalIndent(usr, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(info)
}
