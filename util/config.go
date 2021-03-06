package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// initialize log file
func LogInit(file string) {
	if !isFileExist(file) {
		myfile, err := os.Create(file)
		myfile.Close()
		if err != nil {
			ErrFatal(err, "")
		}
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		ErrFatal(err, "")
	}

	log.SetOutput(f)

	InfoPrint(3, "Initializing program")
	InfoPrint(3, "Log file configured")
}

// Read env file
func EnvReader() error {
	err := godotenv.Load(".env")
	return err
}

// Get database connection string
func GetDbSource() (string, string) {
	var source string
	c := getDbConfig()
	switch c.Connection {
	case "mysql":
		source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.Database,
		)
	case "psql":
		source = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			c.Host,
			c.Username,
			c.Password,
			c.Database,
			c.Port)

	}
	b, err := json.Marshal(c)
	if err != nil {
		ErrFatal(err, "")
	}

	InfoPrint(3, fmt.Sprintf("Connection is : %s", c.Connection))
	InfoPrint(3, fmt.Sprintf("Source is : %s", string(b)))
	return c.Connection, source
}

// Get database variable
func getDbConfig() (config dbConfig) {
	config.Connection = os.Getenv(("DB_CONNECTION"))
	config.Host = os.Getenv("DB_HOST")
	config.Port = os.Getenv("DB_PORT")
	config.Database = os.Getenv(("DB_DATABASE"))
	config.Username = os.Getenv(("DB_USERNAME"))
	config.Password = os.Getenv(("DB_PASSWORD"))

	return config
}
