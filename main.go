package main

import (
	"log"
	"nesil_coffe/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	// Database instance
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
