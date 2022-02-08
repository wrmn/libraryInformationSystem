package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func borrowMigration() {
	database.DB.AutoMigrate(&models.Borrow{})
}
