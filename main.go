package main

import (
	"golang-crud/src/config"
	"golang-crud/src/models"
	"golang-crud/src/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Loading credentials from dotenv
	godotenv.Load()

	// Establishing DB connection
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	// Register routes
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	// Server listening on port
	log.Println("Server listening on PORT:", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
