package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPassword(hashed, raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
	return err == nil
}
