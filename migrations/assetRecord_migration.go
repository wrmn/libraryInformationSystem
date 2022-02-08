package migrations

import (
	"librarySysfo/database"
	"librarySysfo/models"
)

func assetRecordMigration() {
	database.DB.AutoMigrate(&models.AssetRecord{})
}
