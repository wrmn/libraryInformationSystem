package main

import (
	"flag"
	"log"

	"librarySysfo/command"
	"librarySysfo/config"
	"librarySysfo/database"
)

func init() {
	logFile := "log.txt"

	err := config.EnvReader()
	if err != nil {
		config.ErrFatal(err, "")
	}

	config.LogInit(logFile)
	config.InfoPrint(3, ".env file read successfully")

	command.ActionName = flag.String("r", "", "Action name")
	command.TableName = flag.String("t", "", "Table name")
	command.AllCommand = flag.Bool("a", false, "All")
	err = database.InitDatabase()
	if err != nil {
		config.ErrFatal(err, "")
	}

	flag.Parse()
}

func main() {
	log.Println("Starting Program")
	command.ReadCommand()
}
