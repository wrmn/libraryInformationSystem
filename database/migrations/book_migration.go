package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func bookMigration() {
	database.DB.AutoMigrate(&models.Book{})
}
