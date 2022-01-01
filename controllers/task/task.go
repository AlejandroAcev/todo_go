package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alejandroacev/todo_go/connector"
	"github.com/alejandroacev/todo_go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	err             error
	ctx, collection = connector.Collection()
)

func Show(w http.ResponseWriter, r *http.Request) {
	// var tasks []models.Task

}

func Create(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	createdTask, err := models.CreateTask(&task)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)

}
