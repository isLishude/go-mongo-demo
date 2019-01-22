package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	uri := "mongodb://127.0.0.1:27017/?replSet=test"
	dbName := "test"
	colName := "test"
	ctx := context.Background()

	cli, err := mongo.Connect(ctx, uri)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)

	col := cli.Database(dbName).Collection(colName)

	// InsertMany
	{
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
		tmp := bson.M{"field": 11}
		if _, err := col.InsertOne(ctx, tmp); err != nil {
			log.Fatal(err)
		}
	}
}
