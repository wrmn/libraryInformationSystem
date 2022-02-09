package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func visitorMigration() {
	database.DB.AutoMigrate(&models.Visitor{})
}
