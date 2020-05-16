package usecase

import (
	"github.com/wakatakeru/user-auth-jwt-api/domain"
)

type UserRepository interface {
	Store(domain.User) (int, error)
	Update(domain.User) (domain.User, error)
	FindByName(string) (domain.User, error)
	FindByID(int) (domain.User, error)
}
