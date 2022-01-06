package main

import (
	"fmt"
	"net/http"

	controllers "github.com/alejandroacev/todo_go/controllers/task"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Starting server")

	router.PathPrefix("/api")

	// General Actions
	router.HandleFunc("/create", controllers.Create).Methods("POST")

	// TASK Actions
	router.HandleFunc("/tasks", controllers.ShowAll).Methods("GET")
	router.HandleFunc("/task/{id}", controllers.Show).Methods("GET")
	router.HandleFunc("/task/{id}", controllers.Update).Methods("POST", "PUT")

	fmt.Println("Server running at 8080")
	http.ListenAndServe(":8080", router)
}
