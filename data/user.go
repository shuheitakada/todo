package data

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int
	Name            string
	Email           string
	CryptedPassword string
	Salt            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (user *User) Create() (err error) {
	result := Db.Create(&user)
	err = result.Error
	return
}

func FindUserByEmail(email string) (user User, err error) {
	result := Db.Where("email = ?", email).First(&user)
	err = result.Error
	return
}

func (user *User) Authenticate(password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(user.CryptedPassword), []byte(password))
	return
}
