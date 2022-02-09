package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func inventoryMigration() {
	task := "Table inventory Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Inventory{})
	util.InfoPrint(2, task)
}
