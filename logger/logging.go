package logger

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func InitLoggingConfig() {
	logPath := os.Getenv("AKAMAI_CLI_LOG_PATH")
	logLevel := strings.ToLower(os.Getenv("AKAMAI_CLI_LOG"))

	switch logLevel {
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.FatalLevel)
	}

	if logPath != "" {

		if _, err := os.Stat(logPath); os.IsNotExist(err) {
			dir := filepath.Dir(logPath)
			err := os.MkdirAll(dir, 0700)
			if err != nil {
				log.Warnln("Failed to create log file, using default stderr.", err)
			}
		}

		// Create your file
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Warnln("Failed to log to file, using default stderr.", err)
		}
	}
}
