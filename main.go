package main

import (
	"github.com/HosseinSiw/book-store-api/users-crud-api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	appPort = "3000"
)

func main() {
	//	defining a new router
	router := mux.NewRouter()

	//	Defining the http routes
	router.HandleFunc("/users", users_crud_api.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", users_crud_api.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", users_crud_api.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", users_crud_api.DeleteUserHandler).Methods("DELETE")

	log.Printf("Starting the server on %v", appPort)
	err := http.ListenAndServe(appPort, router)
	if err != nil {
		log.Fatal("Couldn't listen and serve on port", appPort)
	}

}
