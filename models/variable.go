package models

import (
	"database/sql"
	"net/http"
)

type Response struct {
	Writer http.ResponseWriter
	Data   []byte
	Code   int
}

type DbConfig struct {
	Connection string `json:"connection"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Database   string `json:"database"`
}

var (
	DbAction   *string
	SysAction  *string
	TableName  *string
	AllCommand *bool
	Task       string
	Dmy        string = "2006-01-02 "
	Dmyhms     string = "2006-01-02 15:04:05"
	Gender            = []string{"M", "F"}
)

type QueryParam struct {
	Db        *sql.DB
	TableName string
}

type SearchParam struct {
	Column string
	Value  interface{}
}
