package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// check if file exist
func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// date random in time.Time with param string formatted date. format param : yyyy-mm-dd
func dateRandom(minYear string, maxYear string) time.Time {
	min, _ := time.Parse(dmy, minYear)
	max, _ := time.Parse(dmy, maxYear)
	return gofakeit.DateRange(min, max)
}

// Create file
func createFile(file string) error {
	myfile, err := os.Create(file)
	myfile.Close()
	return err
}

// Set Fatal for error
func errFatal(err error, msg string) {
	fmt.Println("ERROR : runnig task! Check log file for more info")
	if msg != "" {
		fmt.Printf("System says : %s\n", msg)
	}
	log.Fatalf("Some error occured. Err : %s", err)
}

// Print task on terminal and log
// 1=TASK;2=DONE;3=INFO;4=WARNING;5=ERROR;
func infoPrint(status int, msg string) {
	var info string
	switch status {
	case 1:
		info = "TASK"
	case 2:
		info = "DONE"
	case 3:
		info = "INFO"
	case 4:
		info = "WARNING"
	case 5:
		info = "ERROR"
	}
	log.Printf("%s : %s\n", info, msg)
	fmt.Printf("%s : %s\n", info, msg)
}

// Read string from file
func readFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		errFatal(err, "")
	}
	return string(content)
}

// Format int to ddc format
func intToDdc(c int) string {
	var num string
	if c < 10 {
		num = fmt.Sprintf("00%d", c)
	} else if c < 100 {
		num = fmt.Sprintf("0%d", c)
	} else {
		num = strconv.Itoa(c)
	}
	return num
}

func randDigit(i int) (s string) {
	var d int
	for c := 0; c < i; c++ {
		if c == 0 {
			d = rand.Intn(8) + 1
		} else {
			d = rand.Intn(9)
		}
		s = fmt.Sprintf("%s%d", s, d)
	}
	return s
}
