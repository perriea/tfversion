package terraform

import (
	"fmt"
	"os"
)

// Quiet : Show or not message
func Quiet(message string, quiet bool) {
	if quiet == true {
		fmt.Println(message)
	}
}

// Debug : Show debug message
func Debug(message string) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(message)
	}
}
