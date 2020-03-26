package utils

import "log"

func LogError(err error, context string) {
	if err != nil {
		log.Printf("%s[ERROR] %v", err)
	}
}