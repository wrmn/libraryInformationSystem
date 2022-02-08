package config

import (
	"encoding/json"
	"fmt"
	"librarySysfo/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// initialize log file
func LogInit(file string) {
	if !isFileExist(file) {
		err := createFile(file)
		if err != nil {
			ErrFatal(err, "")
		}
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		ErrFatal(err, "")
	}

	log.SetOutput(f)

	log.Println("Initializing program")
	InfoPrint(3, "Log file configured")
}

// Read env file
func EnvReader() error {
	err := godotenv.Load(".env")
	return err
}

// Get database connection string
func GetDbSource() (string, string) {
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
		ErrFatal(err, "")
	}

	InfoPrint(3, fmt.Sprintf("Connection is : %s", c.Connection))
	InfoPrint(3, fmt.Sprintf("Source is : %s", string(b)))
	return c.Connection, source
}

// Get database variable
func getDbConfig() (config models.DbConfig) {
	config.Connection = os.Getenv(("DB_CONNECTION"))
	config.Host = os.Getenv("DB_HOST")
	config.Port = os.Getenv("DB_PORT")
	config.Database = os.Getenv(("DB_DATABASE"))
	config.Username = os.Getenv(("DB_USERNAME"))
	config.Password = os.Getenv(("DB_PASSWORD"))

	return config
}
