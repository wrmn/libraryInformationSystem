package util

import (
	"fmt"
	"log"

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
	switch status {
	case 1:
		logInfo = "Start   "
		info = color.CyanString(logInfo)
	case 2:
		logInfo = "Done    "
		info = color.GreenString(logInfo)
	case 3:
		logInfo = "Info    "
		info = color.BlueString(logInfo)
	case 4:
		logInfo = "Warning "
		info = color.YellowString(logInfo)
	case 5:
		logInfo = "Error   "
		info = color.RedString(logInfo)
	}
	log.Printf("%s : %s\n", logInfo, msg)
	fmt.Printf("%s : %s\n", info, msg)
}
