package model

import "errors"

var (
	//ErrAccessDenied returns when a user do not have access
	ErrAccessDenied = errors.New("access denied")
)

//User user struct
type User struct {
	//Id user id
	Id int `json:"id"`
	//Name user name
	Name string `json:"name"`
	//LastName user lastname
	LastName string `json:"lastname"`
	//User user
	User string `json:"user,omitempty"`
	//Password password
	Password string `json:"password,omitempty"`
	//Root it is true if user is a root
	Root     bool   `json:"root"`
}

//UserFinder interface implemented by all structs to manage user access
type UserFinder interface {
	FindUser(*User) error
}
