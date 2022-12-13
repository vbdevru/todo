package todo

type User struct {
	Id int `json:"-"`
	Name string `json:"name"`
	Username string `jsom:"username"`
	Password string `json:"password"`
}