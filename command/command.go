package command

import (
	"librarySysfo/migrations"
)

func ReadCommand() {
	switch *ActionName {
	case "migrate":
		if *TableName != "" {
			migrations.MigrateTable(*TableName)
		} else {
			migrations.MigrateAll()
		}
	case "seed":
	}
}
