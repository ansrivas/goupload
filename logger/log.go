package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Logger should be used everywhere throughout the application
var (
	Logger = log.New().WithField("appName", "goupload")
)

func init() {

	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

}
