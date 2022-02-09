package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func ddcMigration() {
	database.DB.AutoMigrate(&models.Ddc{})
}
