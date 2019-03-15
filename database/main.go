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

	// list databses
	cli.ListDatabases(context.Background(), nil, nil)

	// get database instance
	db := cli.Database(dbName)

	// list collection in database
	db.ListCollections(context.Background(), nil, nil)

	// run commands @see https://docs.mongodb.com/manual/reference/method/db.runCommand/
	// e.g create collection
	db.RunCommand(context.Background(), bson.M{"create": colName})

	// drop database
	db.Drop(context.Background())
}
