package utilities

import "log"

func ErrorHandler(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
