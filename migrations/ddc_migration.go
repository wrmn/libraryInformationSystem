package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func ddcMigration() {
	database.DB.AutoMigrate(&models.Ddc{})
}
