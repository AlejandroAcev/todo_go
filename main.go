package main

import (
	"fmt"
	"net/http"

	"github.com/alejandroacev/todo_go/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Starting server")

	// router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/show", controllers.Show)
	router.HandleFunc("/create", controllers.Create)
	// http.Handle("/", router)

	http.ListenAndServe(":8080", router)
	fmt.Println("Server running at 8080")
}
