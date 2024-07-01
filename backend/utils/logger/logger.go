package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var logFile *os.File

// InitLogger initializes the logger by creating the directory if it doesn't exist,
// opening the log file, and setting the log output to this file.
func InitLogger() {
	var err error

	// Create directory if it doesn't exist
	logDir := "./data/logs"
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the log file exists and rename if it does
	logPath := filepath.Join(logDir, "logs.log")
	if _, err = os.Stat(logPath); err == nil {
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		err = os.Rename(logPath, filepath.Join(logDir, "logs_"+timestamp+".log"))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open or create the log file
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Set log output to the log file
	log.SetOutput(logFile)
}

// Info logs an info message.
func Info(msg string) {
	infoLogger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
	infoLogger.Println(msg)
}

// Error logs an error message.
func Error(msg string) {
	errorLogger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime)
	errorLogger.Println(msg)
}

// Warning logs a warning message.
func Warning(msg string) {
	warningLogger := log.New(logFile, "WARNING: ", log.Ldate|log.Ltime)
	warningLogger.Println(msg)
}
