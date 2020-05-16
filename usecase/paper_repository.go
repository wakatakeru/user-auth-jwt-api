package usecase

import (
	"github.comm/wakatakeru/user-auth-jwt-api/domain"
)

type PaperRepository interface {
	Store(domain.User) (int, error)
	Update(domain.User) (domain.User, error)
	FindByName(string) (domain.User, error)
	FindByID(int) (domain.User, error)
}
