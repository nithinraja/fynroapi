package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// Init initializes the global logger
func Init() {
	Log = logrus.New()

	// Output to stdout instead of the default stderr
	Log.Out = os.Stdout

	// Set log format (JSONFormatter or TextFormatter)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set log level (can also be made configurable)
	Log.SetLevel(logrus.InfoLevel)

	Log.Info("📋 Logger initialized")
}
