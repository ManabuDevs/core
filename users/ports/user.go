package users

type UserServices interface {
	Get() ([][]string, error)
}

type UserRepositories interface {
	Get() ([][]string, error)
}
