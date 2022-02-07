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

	return nil
}

// make query string for insert
func composeInsert(db string, field interface{}) string {
	fieldType := reflect.TypeOf(field)
	fieldValue := reflect.ValueOf(field)
	var columns, values string

	for i := 0; i < fieldType.NumField(); i++ {
		column := fieldType.Field(i)

		columns = fmt.Sprintf("%s`%s`", columns, strings.ToLower(column.Name))
		if column.Type == reflect.TypeOf(values) {
			value := fieldValue.Field(i).String()
			values = fmt.Sprintf("%s\"%s\"", values, value)
		} else if column.Type == reflect.TypeOf(true) {
			value := fieldValue.Field(i).Bool()
			values = fmt.Sprintf("%s%t", values, value)
		} else if column.Type == reflect.TypeOf(1) {
			value := fieldValue.Field(i).Int()
			if value != 0 {
				values = fmt.Sprintf("%s%d", values, value)
			} else {
				values = fmt.Sprintf("%sNULL", values)
			}
		}

		if i != fieldType.NumField()-1 {
			columns = fmt.Sprintf("%s,", columns)
			values = fmt.Sprintf("%s,", values)
		}
	}

	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", db, columns, values)
}

// count total book and inventory that have same registration id
func totalSerial(i *sql.DB, s SearchParam) int {
	param := QueryParam{
		Db: i,
	}
	paramBook, paramInventory := param, param
	paramBook.TableName, paramInventory.TableName = "book", "inventory"
	bookCount := selectCount(paramBook, s)
	inventCount := selectCount(paramInventory, s)
	return bookCount + inventCount
}

func selectCount(i QueryParam, s SearchParam) (result int) {
	var query string
	if reflect.TypeOf(s.Value) == reflect.TypeOf("") {
		query = fmt.Sprintf("SELECT count(*) FROM %s WHERE %s='%v'", i.TableName, s.Column, s.Value)
	} else if reflect.TypeOf(s.Value) == reflect.TypeOf(1) {
		query = fmt.Sprintf("SELECT count(*) FROM %s WHERE %s=%v", i.TableName, s.Column, s.Value)
	} else if reflect.TypeOf(s.Value) == reflect.TypeOf(true) {
		query = fmt.Sprintf("SELECT count(*) FROM %s WHERE %s=%v", i.TableName, s.Column, s.Value)
	}
	if err := i.Db.QueryRow(query).Scan(&result); err != nil {
		fmt.Println(err.Error())
	}

	return result
}
