package controllers

type CryptHandler interface {
	Hash(string) (string, error)
	Verify(string, string) error
}
