package main

import (
	"context"
	"fmt"
)

func init() {
	fmt.Println("demo find running")
}

var tableName = "trx"
var dbName = "test"
var url = "mongodb://127.0.0.1:27017"

// FindDemo is
func FindDemo() {
	cur, err := MongoTrxCollection.Find(MongoCtx, map[string]string{"from": "a"})
	CheckError(err)
	defer cur.Close(context.Background())

	for cur.Next(MongoCtx) {
		tmp := new(Trx)
		CheckError(cur.Decode(tmp))
		fmt.Printf("%+v\n", tmp)
	}

	CheckError(cur.Err())
}

// FindOneDemo is
func FindOneDemo() {
	defer MongoCancel()
	result := MongoTrxCollection.FindOne(MongoCtx, map[string]string{"from": "a"})

	tmp := new(Trx)
	if err := result.Decode(tmp); err != nil {
		fmt.Println("[Mongo::FindOne] nothing found")
		return
	}
	fmt.Printf("%+v\n", tmp)
}
