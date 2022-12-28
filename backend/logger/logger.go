package logger

import (
	"log"
	"os"
)

var (
	infoLogger    = log.New(log.Writer(), "INFO:: ", log.Ldate|log.Ltime)
	warningLogger = log.New(log.Writer(), "WARNING:: ", log.Ldate|log.Ltime)
	errorLogger   = log.New(log.Writer(), "ERROR:: ", log.Ldate|log.Ltime)
	fatalLogger   = log.New(log.Writer(), "FATAL:: ", log.Ldate|log.Ltime)
	DefaultLogger = infoLogger
)

func Info(msg string, args ...interface{}) {
	infoLogger.Printf(msg+"\n", args...)
}

func Warning(msg string, args ...interface{}) {
	warningLogger.Printf(msg+"\n", args...)
}

func Error(msg string, args ...interface{}) {
	errorLogger.Printf(msg+"\n", args...)
}

func FatalOnError(err error) {
	if err != nil {
		fatalLogger.Println(err.Error())
		os.Exit(1)
	}
}
