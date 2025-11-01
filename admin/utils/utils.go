package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordIsCorrect(pwd, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)) == nil
}

func HashPassword(pwd string) string {
	if hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		log.Println("[ERROR]: Failed to hash password:", err.Error())
		return ""
	} else {
		return string(hash)
	}
}
