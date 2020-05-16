package controllers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/wakatakeru/user-auth-jwt-api/domain"
	"github.com/wakatakeru/user-auth-jwt-api/interfaces/database"
	"github.com/wakatakeru/user-auth-jwt-api/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	user := domain.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id, err := controller.Interactor.Add(user)
	// [TODO]: Handler key dupulication
	if err != nil || id == 0 {
		// [TODO]: Return an appropriate HTTP status code
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	user.ID = id
	// [TODO]: Return JSON Web Token
	c.JSON(http.StatusCreated, user.ID)
}

func hashPassword(plainPassword string) (hashedPassword string, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassword = string(password)
	return
}
