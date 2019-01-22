package main

import (
	"context"
	"fmt"
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

	// Find
	{
		cur, err := col.Find(ctx, bson.M{"field": "value0"})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			tmp := make(map[string]interface{})
			if err := cur.Decode(&tmp); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%+v\n", tmp)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
	}
	// FindOne
	{
		result := col.FindOne(ctx, bson.M{"field": "value0"})
		tmp := make(map[string]interface{})
		if err := result.Decode(&tmp); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", tmp)
	}
}
