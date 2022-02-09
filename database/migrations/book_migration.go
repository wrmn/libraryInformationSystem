package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func bookMigration() {
	task := "Table book Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Book{})
	util.InfoPrint(2, task)
}
