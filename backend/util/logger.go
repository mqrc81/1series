package util

import (
	"log"
	"os"
)

const logArgs = log.Ldate | log.Ltime | log.Lshortfile

const (
	reset  = "\033[0m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
	red    = "\033[31m"
	purple = "\033[35m"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Panic(err error)
}

type logger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	panic *log.Logger
}

func NewLogger() Logger {
	return &logger{
		info:  log.New(log.Writer(), cyan+"INFO: "+reset, logArgs),
		warn:  log.New(log.Writer(), yellow+"WARN: "+reset, logArgs),
		error: log.New(log.Writer(), red+"ERROR: "+reset, logArgs),
		panic: log.New(log.Writer(), purple+"PANIC: "+reset, logArgs),
	}
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.info.Printf(msg+"\n", args...)
}

func (l *logger) Warn(msg string, args ...interface{}) {
	l.warn.Printf(msg+"\n", args...)
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.error.Printf(msg+"\n", args...)
}

func (l *logger) Panic(err error) {
	l.panic.Printf(err.Error() + "\n")
	os.Exit(1)
}
