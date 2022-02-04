package main

var (
	dbAction   *string
	sysAction  *string
	tableName  *string
	allCommand *bool
	task       string
)

type DbConfig struct {
	Connection string `json:"connection"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Database   string `json:"database"`
}
