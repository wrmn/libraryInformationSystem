package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logFile := "log.txt"

	err := envReader()
	if err != nil {
		errFatal(err, "")
	}

	logInit(logFile)
	infoPrint(3, ".env file read successfully")

	dbAction = flag.String("d", "", "database command.")
	sysAction = flag.String("r", "", "run action system")
	tableName = flag.String("t", "", "create table")
	allCommand = flag.Bool("a", false, "run all query file")

	flag.Parse()
}

func main() {
	log.Println("Starting Program")
	if *dbAction != "" {
		action := fmt.Sprintf("%s table %s", *dbAction, *tableName)
		if *allCommand {
			action = fmt.Sprintf("%s all table", *dbAction)
		}
		infoPrint(3, fmt.Sprintf("Command Is \"database\", Action is \"%s\" ", action))
		databaseAction()
	}
}

func databaseAction() {
	switch *dbAction {
	case "test":
		task = "Testing database connection"
		infoPrint(1, task)
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

		infoPrint(2, task)
		infoPrint(3, "Database connected")

	case "create":
		db, err := initDb()
		if err != nil {
			errFatal(err, "check database connection! Run test!")
			break
		}
		defer db.Close()

		if *tableName != "" {
			task = fmt.Sprintf("Creating table %s", *tableName)
			infoPrint(1, task)

			content := readFile(fmt.Sprintf("query/create/%s.sql", *tableName))

			err = runQuery(db, content)
			if err != nil {
				errFatal(err, "")
				break
			}

			infoPrint(2, task)
		} else if *allCommand {
			mainTask := "Running query create for all table"
			infoPrint(1, mainTask)

			files, err := ioutil.ReadDir("query/create")
			if err != nil {
				errFatal(err, "")
			}

			for _, f := range files {
				name := f.Name()

				reg, err := regexp.Compile("(.sql)")
				if err != nil {
					log.Fatal(err)
				}

				table := reg.ReplaceAllString(name, "")

				task = fmt.Sprintf("Creating table %s", table)
				infoPrint(1, task)

				content := readFile(fmt.Sprintf("query/create/%s.sql", table))

				err = runQuery(db, content)
				if err != nil {
					errFatal(err, "")
					break
				}
				infoPrint(2, task)
			}
			infoPrint(2, mainTask)
		}

	default:
		log.Println("dbAction is not set ")
	}

}
