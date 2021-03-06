package util

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

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
	var info, logInfo string
	t := time.Now().Format(Dmyhms)
	switch status {
	case 1:
		logInfo = "START"
		info = color.HiCyanString(logInfo)
	case 2:
		logInfo = "DONE"
		info = color.HiGreenString(logInfo)
	case 3:
		logInfo = "INFO"
		info = color.HiBlueString(logInfo)
	case 4:
		logInfo = "WARNING"
		info = color.HiYellowString(logInfo)
	case 5:
		logInfo = "ERROR"
		info = color.HiRedString(logInfo)
	}
	log.Printf("%s : %s\n", logInfo, msg)
	fmt.Printf("%s [%s] %s\n", t, info, msg)
}
