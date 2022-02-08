package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func userMigration() {
	database.DB.AutoMigrate(&models.User{})
}
