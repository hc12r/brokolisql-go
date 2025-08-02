package utils

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
