package infrastructure

import "golang.org/x/crypto/bcrypt"

type CryptHandler struct {
	DefaultCost int
}

func NewCryptHandler() *CryptHandler {
	cryptHandler := new(CryptHandler)
	cryptHandler.DefaultCost = bcrypt.DefaultCost

	return cryptHandler
}

func (handler *CryptHandler) Hash(plain string) (hashed string, err error) {
	cryptHandler := NewCryptHandler()

	hashedByte, err := bcrypt.GenerateFromPassword([]byte(plain), cryptHandler.DefaultCost)
	if err != nil {
		return "", err
	}

	hashed = string(hashedByte)
	return
}

func (handler *CryptHandler) Verify(hashed string, plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
