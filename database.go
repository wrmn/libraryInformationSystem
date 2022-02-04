package main

import (
	"database/sql"
)

func initDb() (*sql.DB, error) {
	connection, dbSource := getDbSource()

	db, err := sql.Open(connection, dbSource)

	return db, err
}

func pingDB(db *sql.DB) error {
	err := db.Ping()
	return err
}

func runQuery(db *sql.DB, content string) error {
	res, err := db.Query(content)
	if err != nil {
		return err
	}

	defer res.Close()

	infoPrint(3, "Successfully run query")
	return nil
}
