package routes

import (
	"github.com/gorilla/mux"
	"golang-crud/src/controllers"
	"golang-crud/src/middleware"
)

func RegisterRoutes(r *mux.Router){
	
	r.HandleFunc("/register", controllers.Register).Methods("Post")
	r.HandleFunc("login",controllers.Login).Methods("Post")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
}