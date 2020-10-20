package data

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := "dbname=todo sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return
}
