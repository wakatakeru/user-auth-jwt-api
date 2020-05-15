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
	mockSqlResult.EXPECT().LastInsertId().Return(expectedID, err)

	userRepository := NewUserRepository(mockSqlHandler)
	resultID, err := userRepository.Store(user)

	if err != nil {
		t.Error("Store is not same as expected")
	}

	if !reflect.DeepEqual(int(expectedID), resultID) {
		t.Error("Store is not same as expected")
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockSqlHandler := NewMockSqlHandler(ctrl)
	mockRow := NewMockSqlRow(ctrl)

	query := "UPDATE users SET name=?, display_name=?, email=?, password=? WHERE id=?"
	user := domain.User{}
	var err error
	var id int
	var name string
	var display_name string
	var email string
	var password string

	mockSqlHandler.EXPECT().Query(
		query,
		user.Name,
		user.DisplayName,
		user.Email,
		user.Password,
		user.ID,
	).Return(mockRow, err)
	mockRow.EXPECT().Next()
	mockRow.EXPECT().Scan(&id, &name, &display_name, &email, &password)
	mockRow.EXPECT().Close()

	userRepository := NewUserRepository(mockSqlHandler)
	_, err = userRepository.Update(user)

	if err != nil {
		t.Error("Update is not same as expected")
	}
}
