package main

import (
	"golang.org/x/crypto/bcrypt"
)

func cryptePassword(password string) (string, error) {
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(cryptedPassword), nil
}
