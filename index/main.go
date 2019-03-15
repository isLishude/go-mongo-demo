package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://127.0.0.1:27017/?replSet=test"
	dbName := "test"
	colName := "test"

	cli, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(nil)

	col := cli.Database(dbName).Collection(colName)

	// Index build
	{
		index := mongo.IndexModel{
			Keys:    bson.M{"txid": 1},
			Options: options.Index().SetBackground(true),
		}
		if _, err := col.Indexes().CreateOne(context.Background(), index); err != nil {
			log.Fatal(err)
		}
	}

	// Get index list
	{
		cur, err := col.Indexes().List(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(context.Background())

		res := make(map[string]interface{})
		for cur.Next(context.TODO()) {
			if err = cur.Decode(&res); err != nil {
				log.Fatal(err)
			}
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", res)
	}
}
