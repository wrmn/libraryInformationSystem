package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func inventoryMigration() {
	database.DB.AutoMigrate(&models.Inventory{})
}
