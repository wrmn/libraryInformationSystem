package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.Println("Starting")

	logFile := "log.txt"

	err := envReader()
	if err != nil {
		errFatal(err, "")
	}

	logInit(logFile)
	log.Println(".env file read successfully")

	dbAction = flag.String("d", "", "database command.")
	fileQuery = flag.String("f", "", "file that contain query to run")

	flag.Parse()
}

func main() {

	if *dbAction != "" {
		log.Printf("dbAction is \"%s\" ", *dbAction)
	}

	switch *dbAction {
	case "test":
		db, err := initDb()
		if err != nil {
			errFatal(err, "check .env file")
			break
		}

		defer db.Close()

		err = pingDB(db)
		if err != nil {
			errFatal(err, "check connection or env file")
			break
		}

		fmt.Println("testing database connection")
		logAndPrint("database successfully connect")
	case "run":
		if *fileQuery == "" {
			logAndPrint("Need file that contain query to run.")
			break
		}

		logAndPrint(fmt.Sprintf("file query  is %s ", *fileQuery))
		content, err := ioutil.ReadFile(*fileQuery)
		if err != nil {
			fmt.Println(err)
			errFatal(err, "")
		}
		logMsg := fmt.Sprintf("query : \n%s", string(content))
		logAndPrint(logMsg)

		db, err := initDb()
		if err != nil {
			errFatal(err, "check .env file")
			break
		}
		defer db.Close()

		err = runQuery(db, string(content))
		if err != nil {
			errFatal(err, "check .env file")
			break
		}

	default:
		log.Println("dbAction is not set ")
	}

}
