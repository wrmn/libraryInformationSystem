package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"librarySysfo/models"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// FILE
// ===========

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

// Read string from file
func ReadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		ErrFatal(err, "")
	}
	return string(content)
}

// Print Error information
func ErrFatal(err error, msg string) {
	fmt.Println("ERROR : runnig task! Check log file for more info")
	if msg != "" {
		fmt.Printf("System says : %s\n", msg)
	}
	log.Fatalf("Some error occured. Err : %s", err)
}

// Print task on terminal and log
// 1=TASK;2=DONE;3=INFO;4=WARNING;5=ERROR;
func InfoPrint(status int, msg string) {
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

// DATA FORMATTER
// ===========

// make response for http request
func ResponseFormatter(r models.Response) {
	r.Writer.Header().Set("Content-Type", "application/json")
	r.Writer.WriteHeader(r.Code)
	r.Writer.Write(r.Data)
}

// Format int to ddc format
func IntToDdc(c int) string {
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

// BAD REQUEST FORMATTER
// =====================

// make response param for wrong application type
func BadType(w http.ResponseWriter) models.Response {
	return models.Response{
		Writer: w,
		Data:   []byte("Content Type is not application/json"),
		Code:   http.StatusUnsupportedMediaType,
	}
}

func BadRequest(w http.ResponseWriter, msg string) models.Response {
	return models.Response{
		Writer: w,
		Data:   []byte(msg),
		Code:   http.StatusBadRequest,
	}
}

// RANDOM GENERATOR
// ================

// give string of random digit with n lenght
func RandDigit(n int) (s string) {
	var d int
	for c := 0; c < n; c++ {
		if c == 0 {
			d = rand.Intn(8) + 1
		} else {
			d = rand.Intn(9)
		}
		s = fmt.Sprintf("%s%d", s, d)
	}
	return s
}

// date random in time.Time with param string formatted date. format param : yyyy-mm-dd
func DateRandom(minYear string, maxYear string) time.Time {
	min, _ := time.Parse(models.Dmy, minYear)
	max, _ := time.Parse(models.Dmy, maxYear)
	return gofakeit.DateRange(min, max)
}
