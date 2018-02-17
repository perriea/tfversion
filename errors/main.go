package errors

import (
	"fmt"
	"os"

	"github.com/apsdehal/go-logger"
)

var (
	err error
	log *logger.Logger
)

func init() {
	if os.Getenv("DEBUG") == "1" {
		// Create the log file if doesn't exist. And append to it if it already exists.
		f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		Panic(err)

		log, err = logger.New("tfversion-debug", 1, f)
		if err != nil {
			panic(err) // Check for error
		}
	}
}

// Panic : Show fatal errors
func Panic(err error) {
	if err != nil {
		Panic(err)
	}
}

// Debug :
func Debug(level int, title string, message string) {
	if os.Getenv("DEBUG") == "1" {
		switch level {
		case 1:
			log.Info(fmt.Sprintf("%s: %s", title, message))
		case 2:
			log.Warning(fmt.Sprintf("%s: %s", title, message))
		case 3:
			log.Fatal(fmt.Sprintf("%s: %s", title, message))
		default:
			log.Debug(fmt.Sprintf("%s: %s", title, message))
		}
	}
}
