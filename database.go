package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// initialization database
func initDb() (*sql.DB, error) {
	connection, dbSource := getDbSource()

	db, err := sql.Open(connection, dbSource)

	return db, err
}

// check database connection
func pingDB(db *sql.DB) error {
	err := db.Ping()
	return err
}

// run query for database
func runQuery(db *sql.DB, content string) error {
	res, err := db.Query(content)
	if err != nil {
		return err
	}

	defer res.Close()

	infoPrint(3, "Successfully run query")
	return nil
}

func composeInsert(db string, field interface{}) string {
	fieldType := reflect.TypeOf(field)
	fieldValue := reflect.ValueOf(field)
	var columns, values, value string
	for i := 0; i < fieldType.NumField(); i++ {
		column := fieldType.Field(i)
		value = fieldValue.Field(i).String()

		columns = fmt.Sprintf("%s`%s`", columns, strings.ToLower(column.Name))
		if column.Type == reflect.TypeOf(value) {
			values = fmt.Sprintf("%s'%s'", values, value)
		} else {
			values = fmt.Sprintf("%s%d", values, value)
		}

		if i != fieldType.NumField()-1 {
			columns = fmt.Sprintf("%s,", columns)
			values = fmt.Sprintf("%s,", values)
		}

	}
	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", db, columns, values)
}
