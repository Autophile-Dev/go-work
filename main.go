package main

import (
	"log"
	"net/http"

	"example.com/crudapi/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to MongoDB
	controllers.ConnectDB()

	// Initialize Router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Start Server
	log.Fatal(http.ListenAndServe(":8000", router))
}
