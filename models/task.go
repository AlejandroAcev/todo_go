package models

import (
	"time"

	"github.com/alejandroacev/todo_go/connector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx, collection = connector.Collection()
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Tags        []string           `bson:"tags"`
	Completed   bool               `bson:"completed"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func CreateTask(task *Task) (*Task, error) {

	_, err := collection.InsertOne(ctx, task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func FilterTask(filter interface{}) ([]*Task, error) {
	var tasks []*Task
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return tasks, err
	}

	// While has next document
	for cursor.Next(ctx) {

		var t Task

		// Store the document in the var
		err := cursor.Decode(&t)

		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)

	}

	if err := cursor.Err(); err != nil {
		return tasks, err
	}

	cursor.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func GetAllTask() ([]*Task, error) {
	filter := bson.D{{}}
	return FilterTask(filter)
}
