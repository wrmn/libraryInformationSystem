package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func borrowMigration() {
	task := "Table borrow Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Borrow{})
	util.InfoPrint(2, task)
}
