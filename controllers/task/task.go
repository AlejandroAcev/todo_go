package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alejandroacev/todo_go/models"
	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var task *models.Task
	task, err := models.GetTaskByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if task == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func ShowAll(w http.ResponseWriter, r *http.Request) {
	var tasks []*models.Task

	tasks, err := models.GetAllTask()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdTask, err := models.CreateTask(&task)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)

}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var currentTask *models.Task
	currentTask, err := models.GetTaskByID(id)

	if id == "" || currentTask == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newTaskFields *models.Task
	err = json.NewDecoder(r.Body).Decode(&newTaskFields)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if newTaskFields.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty title"))
		return
	}

	updatedTask, err := models.UpdateTask(id, newTaskFields)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)

}
