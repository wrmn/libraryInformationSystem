package database

import (
	"librarySysfo/util"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (err error) {
	source, dbConfig := util.GetDbSource()

	switch source {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	case "psql":
		DB, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	}

	return
}
