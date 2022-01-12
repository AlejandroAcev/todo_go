package routes

import (
	"log"
	"os"

	controllers "github.com/alejandroacev/todo_go/controllers/task"
	"github.com/alejandroacev/todo_go/util"
	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) string {

	subRouter := router.PathPrefix("/api").Subrouter()

	// TASK Actions
	subRouter.HandleFunc("/tasks", controllers.ShowAll).Methods("GET")
	subRouter.HandleFunc("/task", controllers.Create).Methods("POST")
	subRouter.HandleFunc("/task/{id}", controllers.Show).Methods("GET")
	subRouter.HandleFunc("/task/{id}", controllers.Update).Methods("POST", "PUT")
	subRouter.HandleFunc("/task/{id}", controllers.Delete).Methods("DELETE")

	port, err := util.GetEnvProp("PORT")

	if err != nil {
		log.Fatal("Could not load the port: ", err)
		os.Exit(1)
	}

	return ":" + port
}
