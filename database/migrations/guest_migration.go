package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func guestMigration() {
	task := "Table guest Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Guest{})
	util.InfoPrint(2, task)
}
