package seeds

import (
	"fmt"
	"librarySysfo/util"
)

func SeedAll() {
	ddcSeed()
	guestSeed()
	userSeed()

	assetRecordSeed()
	bookSeed()

	borrowSeed()
	employeeSeed()
	inventorySeed()
	memberSeed()
	visitorSeed()
}

func SeedTable(table string) {
	switch table {
	case "assetRecord":
		assetRecordSeed()
	case "book":
		bookSeed()
	case "borrow":
		borrowSeed()
	case "ddc":
		ddcSeed()
	case "employee":
		employeeSeed()
	case "guest":
		guestSeed()
	case "inventory":
		inventorySeed()
	case "member":
		memberSeed()
	case "user":
		userSeed()
	case "visitor":
		visitorSeed()
	default:
		info := fmt.Sprintf("No migration for table : %s", table)
		util.InfoPrint(5, info)
	}
}
