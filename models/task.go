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
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Tags        []string           `bson:"tags" json:"tags"`
	Completed   bool               `bson:"completed" json:"completed"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
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

	return tasks, nil
}

func GetAllTask() ([]*Task, error) {
	filter := bson.D{{}}
	return FilterTask(filter)
}

func GetTaskByID(id string) (*Task, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: oid}}

	tasks, err := FilterTask(filter)

	if len(tasks) <= 0 {
		return nil, err
	}

	return tasks[0], err
}

func CreateTask(task *Task) (*Task, error) {

	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func UpdateTask(id string, updatedTask *Task) (*Task, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: oid}}

	update := bson.M{
		"$set": primitive.D{
			primitive.E{Key: "title", Value: updatedTask.Title},
			primitive.E{Key: "description", Value: updatedTask.Description},
			primitive.E{Key: "completed", Value: updatedTask.Completed},
			primitive.E{Key: "tags", Value: updatedTask.Tags},
			primitive.E{Key: "updatedAt", Value: time.Now()},
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	if result.ModifiedCount != 1 {
		return nil, mongo.ErrNilDocument
	}

	// var task Task
	tasks, err := FilterTask(filter)

	if err != nil || len(tasks) <= 0 {
		return nil, err
	}

	return tasks[0], nil
}

func DeleteTaskByID(id string) (bool, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return false, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: oid}}

	result, err := collection.DeleteOne(ctx, filter)

	if result.DeletedCount <= 0 {
		return false, err
	}

	return true, err
}
