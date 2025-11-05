package router

import (
	"golang.org/x/crypto/bcrypt"
)

func getHashPassword(passw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
