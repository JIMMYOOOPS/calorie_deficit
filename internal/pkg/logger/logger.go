package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	// Output to stdout instead of the default stderr
	Logger.SetOutput(os.Stdout)

	// Set the log level (can be made configurable)
	Logger.SetLevel(logrus.InfoLevel)

	// Use JSON formatter (or TextFormatter)
	Logger.SetFormatter(&logrus.JSONFormatter{})
}

func EnsureLoggerInitialized() {
	if Logger == nil {
		InitLogger()
	}
}
