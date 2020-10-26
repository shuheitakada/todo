package data

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

type Session struct {
	ID        int
	UUID      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	session = Session{
		UUID:   createUUID(),
		UserID: user.ID,
	}
	result := Db.Create(&session)
	err = result.Error
	return
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func FindSessionByUuid(uuid string) (session Session, err error) {
	result := Db.Where("uuid = ?", uuid).First(&session)
	err = result.Error
	return
}

func (session *Session) Delete() (err error) {
	result := Db.Delete(&session)
	err = result.Error
	return
}
