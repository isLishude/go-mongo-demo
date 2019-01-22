package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

func main() {
	uri := "mongodb://127.0.0.1:27017"
	dbName := "test"
	colName := "test"
	ctx := context.Background()

	cli, err := mongo.Connect(ctx, uri)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)

	col := cli.Database(dbName).Collection(colName)

	// Index build
	{
		index := mongo.IndexModel{
			Keys:    bson.M{"txid": 1},
			Options: options.Index().SetBackground(true),
		}
		res, err := col.Indexes().CreateOne(ctx, index)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	// Get index list
	{
		cur, err := col.Indexes().List(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)

		res := make(map[string]interface{})
		for cur.Next(context.TODO()) {
			if err = cur.Decode(res); err != nil {
				log.Fatal(err)
			}
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", res)
	}
}
