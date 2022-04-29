package users

import (
	"database/sql"
	"log"
	"sabasy/internal/db-schema/querys"
	domain "sabasy/users/domain"
	users "sabasy/users/ports"
	"strconv"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) users.UserRepositories {
	return &userRepository{db: db}
}

func (u userRepository) GetUsers() ([][]string, error) {
	var users [][]string

	rows, err := u.db.Query(querys.UsersGet)
	if err != nil {
		log.Printf("cannot execute select query %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var userQuery domain.User

		err := rows.Scan(
			&userQuery.Id,
		)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, []string{
			strconv.Itoa(userQuery.Id),
		})
	}

	return users, nil
}
