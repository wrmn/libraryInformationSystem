package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// initialize log file
func logInit(file string) {
	if !isFileExist(file) {
		err := createFile(file)
		if err != nil {
			errFatal(err, "")
		}
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		errFatal(err, "")
	}

	log.SetOutput(f)

	log.Println("Initializing program")
	infoPrint(3, "Log file configured")
}

// Get database variable
func getDbConfig() (config DbConfig) {
	config.Connection = os.Getenv(("DB_CONNECTION"))
	config.Host = os.Getenv("DB_HOST")
	config.Port = os.Getenv("DB_PORT")
	config.Database = os.Getenv(("DB_DATABASE"))
	config.Username = os.Getenv(("DB_USERNAME"))
	config.Password = os.Getenv(("DB_PASSWORD"))

	return config
}

func getDbSource() (string, string) {
	c := getDbConfig()
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	b, err := json.Marshal(c)
	if err != nil {
		errFatal(err, "")
	}

	infoPrint(3, fmt.Sprintf("Connection is : %s", c.Connection))
	infoPrint(3, fmt.Sprintf("Source is : %s", string(b)))
	return c.Connection, source
}

// Read env file
func envReader() error {
	infoPrint(3, "Reading .env file")
	err := godotenv.Load(".env")
	return err
}
