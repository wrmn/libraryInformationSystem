package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func borrowMigration() {
	database.DB.AutoMigrate(&models.Borrow{})
}
