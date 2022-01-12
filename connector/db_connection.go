package connector

import (
	"context"
	"log"
	"os"

	"github.com/alejandroacev/todo_go/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "todo_db"
	COLLECTION = "todo"
)

var collection *mongo.Collection
var ctx = context.TODO()

func Collection() (context.Context, *mongo.Collection) {
	// Loads the mongoURI
	MONGO_URI, err := util.GetMongoURI()

	if err != nil {
		os.Exit(1)
	}

	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("-> DB connection error:", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("-> DB ping error:", err)
	}

	collection = client.Database(DATABASE).Collection(COLLECTION)

	return ctx, collection
}
