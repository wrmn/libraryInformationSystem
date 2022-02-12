package util

import (
	"net/http"
)

var (
	ActionName *string
	TableName  *string
	AllCommand *bool
	Dmy        string = "2006-01-02"
	Dmyhms     string = "15:04:05.000"
	Gender            = []string{"M", "F"}
)

type response struct {
	Writer http.ResponseWriter
	Data   []byte
	Code   int
}

type dbConfig struct {
	Connection string `json:"connection"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Database   string `json:"database"`
}
