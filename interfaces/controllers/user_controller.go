package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
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
	if err != nil || id == 0 {
		// [TODO]: Return an appropriate HTTP status code
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	user.ID = id
	jwt, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, jwt)
}

func (controller *UserController) Show(c Context) {
	challengeUser, err := verifyJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	name := c.Param("name")
	if challengeUser.Name != name {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	user, err := controller.Interactor.UserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c Context) {
	user := domain.User{}
	challengeUser, err := verifyJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	err = c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	userName := c.Param("name")
	if challengeUser.Name != user.Name || challengeUser.Name != userName {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	count, err := controller.Interactor.Update(user)
	if err != nil || count <= 0 {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	user.Password = ""
	c.JSON(http.StatusCreated, user)
}

func (controller *UserController) Login(c Context) {
	challengeUser := domain.User{}
	err := c.Bind(&challengeUser)
	if err != nil || challengeUser.Name == "" || challengeUser.Password == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user, err := controller.Interactor.UserByName(challengeUser.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	err = verifyPassword(user.Password, challengeUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	jwt, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, jwt)
}

func hashPassword(plainPassword string) (hashedPassword string, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassword = string(password)
	return
}

func verifyPassword(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func generateJWT(user domain.User) (jwtString string, err error) {
	signBytes, err := ioutil.ReadFile(os.Getenv("JWT_PRIV_KEY_PATH"))
	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}

	// Apply method is RS256
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = os.Getenv("JWT_ISSUER")
	claims["sub"] = strconv.Itoa(user.ID)
	claims["name"] = user.Name
	claims["display_name"] = user.DisplayName
	claims["email"] = user.Email

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyJWT(jwtString string) (user domain.User, err error) {
	user = domain.User{}

	type CustomClaims struct {
		Name string `json:name`
		jwt.StandardClaims
	}

	verifyBytes, err := ioutil.ReadFile(os.Getenv("JWT_PUB_KEY_PATH"))
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return user, err
	}

	token, err := jwt.ParseWithClaims(jwtString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		user.ID, _ = strconv.Atoi(claims.Subject)
		user.Name = claims.Name
		return user, nil
	}

	return user, errors.New("Failed Parse JWT")
}
