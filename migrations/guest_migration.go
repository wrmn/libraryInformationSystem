package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func guestMigration() {
	database.DB.AutoMigrate(&models.Guest{})
}
