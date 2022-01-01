package connector

import (
	"context"
	"log"

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

	clientOptions := options.Client().ApplyURI("mongodb+srv://alejacevgonz:Calamard0.1.m@mongogeneric.ujgee.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
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
