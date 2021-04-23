package models

import (
	"github.com/kamva/mgm/v3"
)

//User is the model that defines a user entry

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name string `json:"name" bson:"name`
	Address string `json:"address bson:"address"` 
	Description string `json:"description" bson:"description"`
}

//Create User is a wrapper that creates a new user entry
func CreateUser(name, address, description string) *User {
	return & User {
		Name: name,
		Address: address,
		Description: description,
	}
}