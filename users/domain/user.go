package users

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	GroupID  string //[]Group `json:"groupid"`
}

type Group struct {
	Id          int
	Name        string
	Description string
	Permissions []Permissions
}

type Permissions struct {
	Id          int
	Name        string
	Description string
}
