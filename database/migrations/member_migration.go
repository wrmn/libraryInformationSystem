package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func memberMigration() {
	task := "Table member Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Member{})
	util.InfoPrint(2, task)
}
