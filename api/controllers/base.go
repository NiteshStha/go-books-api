package controllers

import (
	"fmt"
	"log"
	"net/http"
	"nitesh/books-api/api/middlewares"
	"nitesh/books-api/api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App the main App
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the server
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\nCannot connect to database: %s", DbName)
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Connection successfully established with database: %s", DbName)
	}

	a.DB.AutoMigrate(&models.Book{}) // database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	a.Router.HandleFunc("/api/books", a.GetBooks).Methods("GET")
	a.Router.HandleFunc("/api/books/{id}", a.GetBookByID).Methods("GET")
	a.Router.HandleFunc("/api/books", a.CreateBooks).Methods("POST")
	a.Router.HandleFunc("/api/books/{id}", a.UpdateBooks).Methods("PUT")
	a.Router.HandleFunc("/api/books/{id}", a.DeleteBooks).Methods("DELETE")
}

// Runserver starts the server
func (a *App) Runserver() {
	log.Printf("\nServer starting on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
