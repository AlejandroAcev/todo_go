package main

import (
	"fmt"
	"net/http"

	"github.com/alejandroacev/todo_go/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Starting server")

	port := routes.InitializeRoutes(router)

	fmt.Println("Server running at", port)
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(port, handler)

	if err != nil {
		fmt.Println("Error running server", err)
	}
}
