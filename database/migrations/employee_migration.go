package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func employeeMigration() {
	task := "Table employee Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.Employee{})
	util.InfoPrint(2, task)
}
