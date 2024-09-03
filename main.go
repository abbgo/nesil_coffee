package main

import (
	"log"
	"nesil_coffe/config"
	"nesil_coffe/routes"
	"os"

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

	r := routes.Routes()

	// static file
	os.Mkdir("./uploads", os.ModePerm)
	r.Static("/uploads", "./uploads")

	// run routes
	if err := r.Run(":" + os.Getenv("PROJECT_RUN_PORT")); err != nil {
		log.Fatal(err)
	}
}
