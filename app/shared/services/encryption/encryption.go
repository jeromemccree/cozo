package encryption

import (

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	passwordError := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if passwordError != nil {

		return false
	}
	return passwordError == nil
}
