package database

import (
	"librarySysfo/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (err error) {
	source, dbConfig := util.GetDbSource()

	if source == "mysql" {
		DB, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	}

	return
}
