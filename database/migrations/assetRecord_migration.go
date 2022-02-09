package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
)

func assetRecordMigration() {
	task := "Table assetRecord Migrate"
	util.InfoPrint(1, task)
	database.DB.AutoMigrate(&models.AssetRecord{})
	util.InfoPrint(2, task)
}
