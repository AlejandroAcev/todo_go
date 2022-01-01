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

	// General Actions
	router.HandleFunc("/create", controllers.Create)

	// TASK Actions
	router.HandleFunc("/tasks", controllers.ShowAll)
	router.HandleFunc("/task/{id}", controllers.Show)

	fmt.Println("Server running at 8080")
	http.ListenAndServe(":8080", router)
}
