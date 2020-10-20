package data

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DbGorm *gorm.DB

func init() {
	var err error
	dsn := "dbname=todo sslmode=disable"
	DbGorm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return
}
