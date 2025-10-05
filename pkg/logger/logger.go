package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yourusername/go-skeleton/pkg/config"
)

var Log *logrus.Logger

func InitLogger(cfg *config.LogConfig) {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)

	// Set log level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)

	// Set log format
	if cfg.Format == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}

func GetLogger() *logrus.Logger {
	if Log == nil {
		Log = logrus.New()
	}
	return Log
}
