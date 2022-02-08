package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func memberMigration() {
	database.DB.AutoMigrate(&models.Member{})
}
