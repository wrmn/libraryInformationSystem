package main

import (
	"flag"
	"log"

	"librarySysfo/command"
	"librarySysfo/database"
	"librarySysfo/util"
)

func init() {
	logFile := "log.txt"

	err := util.EnvReader()
	if err != nil {
		util.ErrFatal(err, "")
	}

	util.LogInit(logFile)
	util.InfoPrint(3, ".env file read successfully")

	util.ActionName = flag.String("r", "", "Action name")
	util.TableName = flag.String("t", "", "Table name")
	util.AllCommand = flag.Bool("a", false, "All")
	err = database.InitDatabase()
	if err != nil {
		util.ErrFatal(err, "")
	}

	flag.Parse()
}

func main() {
	log.Println("Starting Program")
	command.ReadCommand()
}
