package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func employeeMigration() {
	database.DB.AutoMigrate(&models.Employee{})
}
