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

func (u userRepository) CreateUser(du *domain.User) (*domain.User, error) {
	log.Println("creating a new patient")
	res, err := u.db.Exec("insert into patient (first_name, last_name, address, phone, email) values (?,?,?,?,?)",
		du.Name, du.Username, du.Password, du.GroupID)

	if err != nil {
		log.Printf("cannot save the patient, %s", err.Error())
		return nil, err
	}

	id, err := res.LastInsertId()

	return &domain.User{
		Id:       int(id),
		Name:     du.Name,
		Username: du.Username,
		Password: du.Password,
		GroupID:  du.GroupID,
	}, nil
}
