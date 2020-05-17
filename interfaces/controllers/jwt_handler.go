package controllers

import (
	"github.com/wakatakeru/user-auth-jwt-api/domain"
)

type JWTHandler interface {
	Generate(domain.User) (string, error)
	Verify(string) (domain.User, error)
}
