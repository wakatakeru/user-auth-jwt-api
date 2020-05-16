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

func (repo *UserRepository) Update(u domain.User) (count int, err error) {
	result, err := repo.Execute(
		"UPDATE users SET name=?, display_name=?, email=?, password=? WHERE name=?",
		u.Name,
		u.DisplayName,
		u.Email,
		u.Password,
		u.Name,
	)

	if err != nil {
		return
	}

	count64, err := result.RowsAffected()
	if err != nil {
		return
	}
	count = int(count64)
	return
}

func (repo *UserRepository) FindByName(userName string) (user domain.User, err error) {
	row, err := repo.Query(
		"SELECT id, name, display_name, email, password FROM users WHERE name=?", userName,
	)
	defer row.Close()
	if err != nil {
		return
	}

	var id int
	var name string
	var displayName string
	var email string
	var password string

	row.Next()
	err = row.Scan(&id, &name, &displayName, &email, &password)
	if err != nil {
		return
	}

	user.ID = id
	user.Name = name
	user.DisplayName = displayName
	user.Email = email
	user.Password = password
	return
}

func (repo *UserRepository) FindByID(userID int) (user domain.User, err error) {
	row, err := repo.Query(
		"SELECT id, name, display_name, email, password FROM users WHERE id=?", userID,
	)
	defer row.Close()
	if err != nil {
		return
	}

	var id int
	var name string
	var displayName string
	var email string
	var password string

	row.Next()
	err = row.Scan(&id, &name, &displayName, &email, &password)
	if err != nil {
		return
	}

	user.ID = id
	user.Name = name
	user.DisplayName = displayName
	user.Email = email
	user.Password = password
	return
}
