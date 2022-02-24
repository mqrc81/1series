package util

import (
	"log"
	"os"
)

const (
	reset  = "\033[0m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
	red    = "\033[31m"
)

var (
	infoLogger  = log.New(log.Writer(), cyan+"INFO: "+reset, log.Ldate|log.Ltime)
	warnLogger  = log.New(log.Writer(), yellow+"WARN: "+reset, log.Ldate|log.Ltime)
	errorLogger = log.New(log.Writer(), red+"ERROR: "+reset, log.Ldate|log.Ltime)
)

func LogInfo(msg string, args ...interface{}) {
	infoLogger.Printf(msg+"\n", args...)
}

func LogWarn(msg string, args ...interface{}) {
	warnLogger.Printf(msg+"\n", args...)
}

func LogError(msg string, args ...interface{}) {
	errorLogger.Printf(msg+"\n", args...)
}

func LogPanic(err error) {
	errorLogger.Printf(err.Error() + "\n")
	os.Exit(1)
}
