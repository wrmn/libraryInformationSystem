package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func visitorMigration() {
	task := "Table visitor Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Visitor{})
	util.InfoPrint(2, task)
}
