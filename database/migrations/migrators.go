package migrations

import "fmt"

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
		fmt.Printf("Table not found :  %s", table)
	}
}
