package database

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/user-auth-jwt-api/domain"
)

func TestStore(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockSqlHandler := NewMockSqlHandler(ctrl)
	mockSqlResult := NewMockSqlResult(ctrl)

	user := domain.User{}
	query := "INSERT INTO users (name, display_name, email, password) VALUES (?,?,?,?)"
	var expectedID int64
	var err error

	mockSqlHandler.EXPECT().Execute(query, user.Name, user.DisplayName, user.Email, user.Password).Return(mockSqlResult, err)
	mockSqlResult.EXPECT().LastInsertedId().Return(expectedID, err)

	userRepository := NewUserRepository(mockSqlHandler)
	resultID, err := userRepository.Store(user)

	if err != nil {
		t.Error("Store is not same as expected")
	}

	if !reflect.DeepEqual(int(expectedID), resultID) {
		t.Error("Store is not same as expected")
	}
}
