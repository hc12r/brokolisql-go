package errors

import (
	"log"
	"os"
)

// For testing purposes, we can override these variables
var osExit = os.Exit
var logFatalf = log.Fatalf

func CheckError(err error) {
	if err != nil {
		logFatalf("error: %v", err)
	}
}
