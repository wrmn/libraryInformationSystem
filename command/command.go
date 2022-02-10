package command

import (
	"fmt"
	"librarySysfo/database/migrations"
	"librarySysfo/database/seeds"
	"librarySysfo/server"
	"librarySysfo/util"
)

// Read command flag from running program
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
			seeds.SeedTable(*util.TableName)
		} else {
			seeds.SeedAll()
		}
	case "serve":
		server.Serve()
	default:
		info := fmt.Sprintf("Command %s not available", *util.ActionName)
		util.InfoPrint(5, info)
	}
}
