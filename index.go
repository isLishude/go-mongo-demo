package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// IndexBuildDemo is
func IndexBuildDemo() {
	index := mongo.IndexModel{
		Keys:    bson.M{"txid": 1},
		Options: options.Index().SetBackground(true),
	}

	res, err := MongoTrxCollection.Indexes().CreateOne(MongoCtx, index)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

// GetIndexesDemo is
func GetIndexesDemo() {
	cur, err := MongoTrxCollection.Indexes().List(MongoCtx)
	CheckError(err)
	defer cur.Close(MongoCtx)

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
