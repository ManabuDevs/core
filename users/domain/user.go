package users

type user struct {
	id       int
	name     string
	username string
	password string
	GroupID  []group
}

type group struct {
	id          int
	name        string
	description string
	permissions []permissions
}

type permissions struct {
	id          int
	name        string
	description string
}
