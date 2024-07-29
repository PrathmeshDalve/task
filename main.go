package main

import (
	"log"
)

func main() {
	err := DbConnection()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	handler()
}
