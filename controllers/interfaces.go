package controllers

import (
	// "fmt"
	// "sync"
	// "net/http"
	// topics "github.com/jaydeep87/poc-go-pg/topics"
	// "github.com/gin-gonic/gin"
)

type Logger interface {
    Debug(message string)
    Info(message string)
    // Error(message string, err error)
}

type FileLogger struct {
    FilePath string
}

func (l FileLogger) Debug(message string) {
    // Implement debug logging to file
}

func (l FileLogger) Info(message string) {
    // Implement info logging to file
}

// func (l FileLogger) Error(message string, err error) {
//     // Implement error logging to file
// }

func LoggerInterFace() {
    logger := FileLogger{FilePath: "app.log"}
    logger.Info("Starting application")
    logger.Debug("Processing request")
    // logger.Error("Failed to process request", err)
}
