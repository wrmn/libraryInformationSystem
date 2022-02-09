package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func ddcMigration() {
	task := "Table ddc Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Ddc{})
	util.InfoPrint(2, task)

}
