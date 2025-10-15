package main

import (
	"log"
	"os"
)

// log level
const (
	InfoLevel = iota
	WarningLevel
	ErrorLevel
)


type Logger struct {
	Level int 
	infoLogger  *log.Logger
	warnLogger  *log.Logger 
    errorLogger *log.Logger
}

var logger *Logger 

func init() {
	 logger = &Logger{
		Level: InfoLevel,
		infoLogger:  log.New(os.Stdout , "INFO" , log.LstdFlags),
		warnLogger:  log.New(os.Stdout , "WARN" , log.LstdFlags),
		errorLogger: log.New(os.Stdout , "ERROR" , log.LstdFlags | log.Lshortfile),
	 }
}

// set log level 

func SelLevel(level int) {
	logger.Level = level 
}

// Methods to add log (at different level)


func Info(message string) {
	if logger.Level <= InfoLevel {
		logger.infoLogger.Println(message)
	}
}

func Error(message string) {
	if logger.Level <= ErrorLevel {
		logger.errorLogger.Println(message)
	}
}

func Warn(message string) {
	if logger.Level <= WarningLevel {
		logger.warnLogger.Println(message)
	}
}