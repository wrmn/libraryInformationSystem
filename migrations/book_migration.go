package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func bookMigration() {
	database.DB.AutoMigrate(&models.Book{})
}
