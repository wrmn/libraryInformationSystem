package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func visitorMigration() {
	database.DB.AutoMigrate(&models.Visitor{})
}
