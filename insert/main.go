package main

import (
	"context"
	"log"
	"time"

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

	// InsertMany
	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		var bulk []interface{}
		for i := 0; i < 10; i++ {
			tmp := bson.M{"field": i}
			bulk = append(bulk, tmp)
		}
		if _, err := col.InsertMany(ctx, bulk); err != nil {
			log.Fatal(err)
		}
	}

	// InsertOne
	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		tmp := bson.M{"field": 11}
		if _, err := col.InsertOne(ctx, tmp); err != nil {
			log.Fatal(err)
		}
	}
}
