package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/user-auth-jwt-api/domain"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedUser domain.User
	var expectedID int
	var err error

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Store(expectedUser).Return(expectedID, err)

	userInteractor := NewUserInteractor(mockUserRepository)
	expectedID, err = userInteractor.Add(expectedUser)

	if err != nil {
		t.Error("Add is not same as expected")
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedUser domain.User
	var err error

	MockUserRepository := NewMockUserRepository(ctrl)
	MockUserRepository.EXPECT().Update(expectedUser).Return(expectedUser, err)

	userInteractor := NewUserInteractor(MockUserRepository)
	expectedUser, err = userInteractor.Update(expectedUser)

	if err != nil {
		t.Error("Update is not same as expected")
	}
}

// UserByName(u string) (user domain.User, err error)

func TestUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedUser domain.User
	var userName string
	var err error

	MockUserRepository := NewMockUserRepository(ctrl)
	MockUserRepository.EXPECT().FindByName(userName).Return(expectedUser, err)

	userInteractor := NewUserInteractor(MockUserRepository)
	expectedUser, err = userInteractor.UserByName(userName)

	if err != nil {
		t.Error("UserByName is not same as expected")
	}
}

func TestUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedUser domain.User
	var userID int
	var err error

	MockUserRepository := NewMockUserRepository(ctrl)
	MockUserRepository.EXPECT().FindByID(userID).Return(expectedUser, err)

	userInteractor := NewUserInteractor(MockUserRepository)
	expectedUser, err = userInteractor.UserByID(userID)

	if err != nil {
		t.Error("UserByID is not same as expected")
	}
}
