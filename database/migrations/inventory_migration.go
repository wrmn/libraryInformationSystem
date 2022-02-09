package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func inventoryMigration() {
	database.DB.AutoMigrate(&models.Inventory{})
}
