package main

import (
	"fmt"
	"net/http"

	"github.com/alejandroacev/todo_go/routes"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Starting server")

	port := routes.InitializeRoutes(router)

	fmt.Println("Server running at", port)
	http.ListenAndServe(":"+port, router)
}
