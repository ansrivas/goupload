package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// SetupLogger sets up logger with a given field and returns a logrus entry.
func SetupLogger(appName string) *logrus.Entry {

	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the Info severity or above.
	logrus.SetLevel(logrus.InfoLevel)

	return logrus.WithField("appName", appName)
}
