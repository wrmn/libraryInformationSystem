package migrations

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
)

func assetRecordMigration() {
	database.DB.AutoMigrate(&models.AssetRecord{})
}
