package database

import (
	"librarySysfo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (err error) {
	source, dbConfig := config.GetDbSource()

	if source == "mysql" {
		DB, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	}

	return
}
