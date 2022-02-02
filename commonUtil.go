package main

import (
	"errors"
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
			log.Fatal("program exit due error :", err.Error())
		}
	}
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	log.SetOutput(f)

	log.Output(1, "this is an event")
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
	return c.Connection, source
}

// Read env file
func envReader() error {
	log.Println("Reading .env file")
	err := godotenv.Load(".env")
	return err
}

// check if file exist
func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// Create file
func createFile(file string) error {
	myfile, err := os.Create(file)
	myfile.Close()
	return err
}

// Set Fatal for error
func errFatal(err error) {
	log.Fatalf("Some error occured. Err : %s", err)
}
