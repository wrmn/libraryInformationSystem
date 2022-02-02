package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting")

	err := envReader()
	if err != nil {
		errFatal(err)
	}

	logInit("log.txt")
	connection, dbSource := getDbSource()

	db, err := sql.Open(connection, dbSource)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

}
