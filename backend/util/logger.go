package util

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime)
	warnLogger  = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime)
	errorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime)
)

func LogInfo(msg string, args ...interface{}) {
	infoLogger.Printf(msg+"\n", args...)
}

func LogWarning(msg string, args ...interface{}) {
	warnLogger.Printf(msg+"\n", args...)
}

func LogError(msg string, args ...interface{}) {
	errorLogger.Printf(msg+"\n", args...)
}

func LogPanic(err error) {
	LogError(err.Error())
	os.Exit(1)
}
