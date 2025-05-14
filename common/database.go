package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewMySQLDatabase() *gorm.DB {
	dsn := EnvLookup("DB_HOST")

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
