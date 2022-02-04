package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

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
func errFatal(err error, msg string) {
	if msg != "" {
		log.Printf("System says : %s", msg)
	}
	log.Fatalf("Some error occured. Err : %s", err)
}

//set output for log and fmt
func logAndPrint(msg string) {
	fmt.Println(msg)
	log.Println(msg)
}
