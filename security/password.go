package security

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(p string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

func VerifyPassword(p, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
