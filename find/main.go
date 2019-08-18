package main

import (
	"context"
	"fmt"
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
	defer cli.Disconnect(context.Background())

	col := cli.Database(dbName).Collection(colName)

	// Find
	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
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
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		result := col.FindOne(ctx, bson.M{"field": "value0"})
		tmp := make(map[string]interface{})
		if err := result.Decode(&tmp); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", tmp)
	}
}
