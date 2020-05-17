package controllers

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/user-auth-jwt-api/domain"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)

	user := domain.User{}
	var err error
	var jwt string

	mockContext := NewMockContext(ctrl)
	mockCryptHandler := NewMockCryptHandler(ctrl)
	mockJWTHandler := NewMockJWTHandler(ctrl)

	mockContext.EXPECT().Bind(&user).Return(err)
	mockCryptHandler.EXPECT().Hash(user.Password).Return(user.Password, err)
	mockJWTHandler.EXPECT().Generate(user).Return(jwt, err)
	mockContext.EXPECT().JSON(http.StatusCreated, jwt)
}
