package utils

import "golang.org/x/crypto/bcrypt"

func PasswordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}
