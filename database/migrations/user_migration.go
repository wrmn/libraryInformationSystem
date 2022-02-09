package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func userMigration() {
	task := "Table user Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.User{})
	util.InfoPrint(2, task)
}
