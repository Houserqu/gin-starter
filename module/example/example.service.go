package example

import (
	"houserqu.com/gin-starter/lib"
)

type Hello struct {
	First string `json:"hello"`
	Last  string `json:"world"`
}

func GetHello() User {
	var user User
	lib.DB.Find(&user, "id = ?", 1)

	return user
}
