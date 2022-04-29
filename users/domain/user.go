package users

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	GroupID  []Group
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
