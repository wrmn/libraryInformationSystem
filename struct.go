package main

import "database/sql"

var (
	dbAction *string
	// sysAction  *string
	tableName  *string
	allCommand *bool
	task       string
)

type InsertParam struct {
	Db        *sql.DB
	TableName string
}

type Ddcs struct {
	Ddc   string
	Group string
	Name  string
}

type DbConfig struct {
	Connection string `json:"connection"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Database   string `json:"database"`
}
