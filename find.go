package main

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
)

// FindDemo is
func FindDemo() {
	cur, err := MongoTrxCollection.Find(MongoCtx, bson.M{"from": "a"})
	CheckError(err)
	defer cur.Close(context.Background())

	for cur.Next(MongoCtx) {
		var tmp Trx
		CheckError(cur.Decode(&tmp))
		fmt.Printf("%+v\n", tmp)
	}

	CheckError(cur.Err())
}

// FindOneDemo is
func FindOneDemo() {
	result := MongoTrxCollection.FindOne(MongoCtx, bson.M{"from": "a"})

	var tmp Trx
	if err := result.Decode(&tmp); err != nil {
		fmt.Println("[Mongo::FindOne] nothing found")
		return
	}
	fmt.Printf("%+v\n", tmp)
}
