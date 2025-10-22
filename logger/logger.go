package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"
	"fmt"
)

var ENVIRONMENT = os.Getenv("ENVIRONMENT")

const (
	MaxSizeBytes = 5 * 1024 * 1024 // 5 MB
	DateLayout   = "2006-01-02"
	TimeLayout   = "2006-01-02_15-04-05"
)

func SetupLogger() *slog.Logger {
    logsDir := "logs"
	archiveDir := filepath.Join(logsDir, "archive")

    if err := os.MkdirAll(logsDir, 0755); err != nil {
        log.Fatalf("Failed to create logs directory: %v", err)
    }
    
    logPath := filepath.Join(logsDir, "current.log")
    rotateIfNeeded(logsDir,archiveDir,logPath)

    logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }

    multiWriter := io.MultiWriter(logFile, os.Stdout)
    
	var level slog.Level

	if ENVIRONMENT == "production" {
		level = slog.LevelError
	} else {
		level = slog.LevelInfo
	}

    handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
        Level:     level,
        AddSource: true,
    })

    logger := slog.New(handler)
    // Log a test message
 //   logger.Info("Logger initialized", "file", logPath)

    return logger
}


func rotateIfNeeded(logsDir , archiveDir,logFilePath string) {
	info, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		return // no file yet
	}

	if err != nil {
		slog.Warn("Unable to stat log file", "error", err)
		return
	}

	// Get last modified date
	modTime := info.ModTime().Format(DateLayout)
	currentDate := time.Now().Format(DateLayout)

	// Check for rotation conditions
	sizeExceeded := info.Size() >= MaxSizeBytes
	newDay := currentDate != modTime

	if sizeExceeded || newDay {
		timestamp := time.Now().Format(TimeLayout)
		// newName := fmt.Sprintf("Current_log_%s.log", timestamp)
		// newPath := filepath.Join(logsDir, newName)

		// if err := os.Rename(logFilePath, newPath); err != nil {
		// 	slog.Warn("Failed to rotate log file", "error", err)
		// 	return
		// }
		archivedName := fmt.Sprintf("log_%s.log", timestamp)
		archivedPath := filepath.Join(archiveDir, archivedName)

		// Close before renaming
		if file, err := os.OpenFile(logFilePath, os.O_WRONLY, 0644); err == nil {
			file.Close()
		}

		// Move to archive folder
		if err := os.Rename(logFilePath, archivedPath); err != nil {
			slog.Warn("Failed to move log to archive", "error", err)
			return
		}
	}
}


// package main

// import (
// 	"log"
// 	"os"
// )

// // log level
// const (
// 	InfoLevel = iota
// 	WarningLevel
// 	ErrorLevel
// )


// type Logger struct {
// 	Level int 
// 	infoLogger  *log.Logger
// 	warnLogger  *log.Logger 
//     errorLogger *log.Logger
// }

// var logger *Logger 

// func init() {
// 	 logger = &Logger{
// 		Level: InfoLevel,
// 		infoLogger:  log.New(os.Stdout , "INFO" , log.LstdFlags),
// 		warnLogger:  log.New(os.Stdout , "WARN" , log.LstdFlags),
// 		errorLogger: log.New(os.Stdout , "ERROR" , log.LstdFlags | log.Lshortfile),
// 	 }
// }

// // set log level 

// func SelLevel(level int) {
// 	logger.Level = level 
// }

// // Methods to add log (at different level)


// func Info(message string) {
// 	if logger.Level <= InfoLevel {
// 		logger.infoLogger.Println(message)
// 	}
// }

// func Error(message string) {
// 	if logger.Level <= ErrorLevel {
// 		logger.errorLogger.Println(message)
// 	}
// }

// func Warn(message string) {
// 	if logger.Level <= WarningLevel {
// 		logger.warnLogger.Println(message)
// 	}
// }