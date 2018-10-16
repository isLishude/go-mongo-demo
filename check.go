package main

import (
	"log"
)

// CheckError is to check it error and exit
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
