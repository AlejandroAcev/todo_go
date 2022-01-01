package controllers

import (
	"testing"
	"time"

	"github.com/alejandroacev/todo_go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTaskController(t *testing.T) {

	taskModel := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Title",
		Description: "Test Description",
		Completed:   true,
		Tags:        []string{"Tag1", "Tag2", "Tag3"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task, err := models.CreateTask(&taskModel)

	if err != nil {
		t.Fatal("CreateTask(task) requires", err)
	}

	if task != task {
		t.Fatal("CreateTask(task) does not match with the object")
	}
}
