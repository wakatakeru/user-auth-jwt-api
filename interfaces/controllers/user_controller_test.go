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

func TestShow(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockContext := NewMockContext(ctrl)
	mockJWTHandler := NewMockJWTHandler(ctrl)

	var authToken string
	var challengeUser domain.User
	var user domain.User
	var name string
	var err error

	mockContext.EXPECT().GetHeader("Authorization").Return(authToken)
	mockJWTHandler.EXPECT().Verify(authToken).Return(challengeUser, err)
	mockContext.EXPECT().Param("name").Return(name)
	mockContext.EXPECT().JSON(http.StatusOK, user)
}
