package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	const (
		uri     = "mongodb://127.0.0.1:27017,127.0.0.1:27018/admin?replicaSet=test"
		dbName  = "test"
		colName = "test"
	)

	cli, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(nil)

	col := cli.Database(dbName).Collection(colName)

	// delete one
	{
		col.DeleteOne(context.Background(), bson.M{"field": "delete one"})
	}

	// delete many
	{
		col.DeleteMany(context.Background(), bson.M{"field": "delete many"})
	}
}
