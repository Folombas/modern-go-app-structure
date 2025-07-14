// pkg/logger/logger.go
package logger

import "log"

func Info(message string) {
	log.Printf("INFO: %s", message)
}


