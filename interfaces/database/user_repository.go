package database

import (
	"github.com/wakatakeru/user-auth-jwt-api/domain"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) UserRepository {
	userRepository := UserRepository{SqlHandler: sqlHandler}
	return userRepository
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (name, display_name, email, password) VALUES (?,?,?,?)", u.Name, u.DisplayName, u.Email, u.Password,
	)

	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}
