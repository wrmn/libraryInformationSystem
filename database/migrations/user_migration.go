package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func userMigration() {
	database.DB.AutoMigrate(&models.User{})
}
