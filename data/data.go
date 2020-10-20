package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var Db *sql.DB
var DbGorm *gorm.DB

func init() {
	var err error
	dsn := "dbname=todo sslmode=disable"
	DbGorm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db, err = sql.Open("postgres", "dbname=todo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
