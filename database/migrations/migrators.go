package migrations

import (
	"fmt"
	"librarySysfo/util"
)

// Run all migration
func MigrateAll() {
	ddcMigration()
	guestMigration()
	userMigration()

	assetRecordMigration()
	bookMigration()

	borrowMigration()
	employeeMigration()
	inventoryMigration()
	memberMigration()
	visitorMigration()

}

// Run Specific Migration
func MigrateTable(table string) {
	switch table {
	case "assetRecord":
		assetRecordMigration()
	case "book":
		bookMigration()
	case "borrow":
		borrowMigration()
	case "ddc":
		ddcMigration()
	case "employee":
		employeeMigration()
	case "guest":
		guestMigration()
	case "inventory":
		inventoryMigration()
	case "member":
		memberMigration()
	case "user":
		userMigration()
	case "visitor":
		visitorMigration()
	default:
		info := fmt.Sprintf("No migration for table : %s", table)
		util.InfoPrint(5, info)
	}
}
