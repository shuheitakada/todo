package data

import (
	"time"
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
