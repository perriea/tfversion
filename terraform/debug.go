package terraform

import (
	"fmt"
	"os"
)

// ShowMessage : Show or not message
func ShowMessage(message string, quiet bool) {
	if !quiet {
		fmt.Println(message)
	}
}

// Debug : Show debug message
func Debug(message string) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(message)
	}
}
