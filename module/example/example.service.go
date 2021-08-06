package example

type Hello struct {
	First string `json:"hello"`
	Last  string `json:"world"`
}

func GetHello() Hello {
	user := Hello{First: "hello", Last: "world"}
	return user
}
