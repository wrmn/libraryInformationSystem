package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func employeeMigration() {
	database.DB.AutoMigrate(&models.Employee{})
}
