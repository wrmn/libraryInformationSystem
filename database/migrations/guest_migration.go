package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func guestMigration() {
	database.DB.AutoMigrate(&models.Guest{})
}
