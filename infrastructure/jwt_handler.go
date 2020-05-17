package infrastructure

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wakatakeru/user-auth-jwt-api/domain"
)

type JWTHandler struct {
	PrivKey *rsa.PrivateKey
	PubKey  *rsa.PublicKey
}

type JSONWebToken struct {
	Token jwt.Token
}

type CustomClaims struct {
	Name string `json:name`
	jwt.StandardClaims
}

func NewJWTHandler() *JWTHandler {
	privBytes, err := ioutil.ReadFile(os.Getenv("JWT_PRIV_KEY_PATH"))
	if err != nil {
		log.Panic(err)
	}

	pubBytes, err := ioutil.ReadFile(os.Getenv("JWT_PUB_KEY_PATH"))
	if err != nil {
		log.Panic(err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		log.Panic(err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Panic(err)
	}

	jwtHandler := new(JWTHandler)
	jwtHandler.PrivKey = privKey
	jwtHandler.PubKey = pubKey

	return jwtHandler
}

func (handler *JWTHandler) Generate(user domain.User) (token string, err error) {
	jwtHandler := NewJWTHandler()
	jsonWebToken := jwt.New(jwt.SigningMethodRS256)

	timeNow := int(time.Now().Unix())
	duraExp, _ := strconv.Atoi(os.Getenv("JWT_EXP_DURA_SEC"))
	expTime := timeNow + duraExp

	claims := jsonWebToken.Claims.(jwt.MapClaims)
	claims["iss"] = os.Getenv("JWT_ISSUER")
	claims["sub"] = strconv.Itoa(user.ID)
	claims["exp"] = expTime
	claims["name"] = user.Name
	claims["display_name"] = user.DisplayName
	claims["email"] = user.Email

	token, err = jsonWebToken.SignedString(jwtHandler.PrivKey)
	if err != nil {
		return "", err
	}
	return
}

func (handler *JWTHandler) Verify(token string) (user domain.User, err error) {
	jwtHandler := NewJWTHandler()
	jsonWebToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtHandler.PubKey, nil
	})

	if claims, ok := jsonWebToken.Claims.(*CustomClaims); ok && jsonWebToken.Valid {
		user.ID, _ = strconv.Atoi(claims.Subject)
		user.Name = claims.Name
		return user, nil
	}

	return user, errors.New("Failed Parse JWT")
}
