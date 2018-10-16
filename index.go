package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// IndexBuildDemo is
func IndexBuildDemo() {
	defer MongoCancel()

	index := mongo.IndexModel{
		Keys:    bson.NewDocument(bson.EC.Int32("txid", int32(1))),
		Options: mongo.NewIndexOptionsBuilder().Background(true).Build(),
	}

	res, err := MongoTrxCollection.Indexes().CreateOne(MongoCtx, index)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

// GetIndexes is
func GetIndexes() {
	defer MongoCancel()

	cur, err := MongoTrxCollection.Indexes().List(MongoCtx)
	CheckError(err)

	defer cur.Close(MongoCtx)

	res := make(map[string]interface{})
	for cur.Next(context.TODO()) {
		cur.Decode(&res)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}
