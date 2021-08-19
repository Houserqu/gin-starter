package example

import "houserqu.com/gin-starter/internal"

type Hello struct {
	First string `json:"hello"`
	Last  string `json:"world"`
}

func GetHello() Hello {
	user := Hello{First: "hello", Last: "world"}
	return user
}

func GetModelByID(id int) Example {
	var example Example
	internal.DB.Take(&example, id)

	return example
}

func GetModelAll() (examples []Example, err error) {
	res := internal.DB.Find(&examples)
	err = res.Error
	return
}
