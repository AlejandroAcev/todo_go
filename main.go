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

	// router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/show", controllers.Show)
	router.HandleFunc("/create", controllers.Create)
	// http.Handle("/", router)

	fmt.Println("Server running at 8080")
	http.ListenAndServe(":8080", router)
}
