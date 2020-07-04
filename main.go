package main

import (
	"log"
	"nitesh/books-api/api/controllers"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment variables")
	}

	app := controllers.App{}
	log.Println(os.Getenv("DBPORT"))
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	app.Runserver()
}
