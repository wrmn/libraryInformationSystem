package command

import (
	"librarySysfo/database/migrations"
	"librarySysfo/database/seeds"
	"librarySysfo/util"
)

func ReadCommand() {
	switch *util.ActionName {
	case "migrate":
		if *util.TableName != "" {
			migrations.MigrateTable(*util.TableName)
		} else {
			migrations.MigrateAll()
		}
	case "seed":
		if *util.TableName != "" {
			// seeds.MigrateTable(*util.TableName)
		} else {
			seeds.SeedAll()
		}
	}
}
