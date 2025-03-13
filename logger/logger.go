package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Log is the global logger instance
var Log = logrus.New()

// InitializeLogger sets up the logger to log in JSON format
func InitializeLogger() {
	// Set log format to JSON
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05Z", // You can adjust timestamp format here
	})

	// Output logs to standard output (console)
	Log.SetOutput(os.Stdout)

	// Set log level (adjust based on your needs)
	Log.SetLevel(logrus.InfoLevel)
}

// LogRequest logs an HTTP request with relevant information
func LogRequest(method, endpoint string) {
	Log.WithFields(logrus.Fields{
		"time":     time.Now().Format("2006-01-02T15:04:05Z"),
		"endpoint": endpoint,
		"method":   method,
		"status":   "success",
	}).Info("SUCCESS")
}

// LogError logs error-related requests with relevant information
func LogError(method, endpoint, errorMessage string) {
	Log.WithFields(logrus.Fields{
		"time":     time.Now().Format("2006-01-02T15:04:05Z"),
		"endpoint": endpoint,
		"method":   method,
		"status":   "error",
		"error":    errorMessage,
	}).Error("FAILED")
}

func LogServerStart(port string, endpoint string) {
	logrus.WithFields(logrus.Fields{
		"time":     time.Now().Format("2006-01-02T15:04:05Z"),
		"status":   "info",
		"message":  "Server listening on port",
		"port":     port,
		"endpoint": endpoint,
	}).Info("Server started")
}

func LogInfo(message string) {
	logrus.WithFields(logrus.Fields{
		"time":    time.Now().Format("2006-01-02T15:04:05Z"),
		"status":  "info",
		"message": message,
	}).Info("INFO")
}
