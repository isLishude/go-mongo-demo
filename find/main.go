package main

import (
	"context"
	"fmt"
	"log"

	"github.com/islishude/demo/mongo/instance"
	"github.com/islishude/demo/mongo/schema"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var tableName = "trx"
var dbName = "test"
var url = "mongodb://127.0.0.1:27017"

func find() {
	defer demoMongo.MongoCancel()
	cur, err := demoMongo.MongoTrxCollection.Find(demoMongo.MongoCtx, map[string]string{"from": "a"})
	checkError(err)
	defer cur.Close(context.Background())

	for cur.Next(demoMongo.MongoCtx) {
		tmp := new(demoTest.Trx)
		checkError(cur.Decode(tmp))
		fmt.Printf("%+v\n", tmp)
	}

	checkError(cur.Err())
}

func findOne() {
	defer demoMongo.MongoCancel()
	result := demoMongo.MongoTrxCollection.FindOne(demoMongo.MongoCtx, map[string]string{"from": "a"})

	tmp := new(demoTest.Trx)
	checkError(result.Decode(tmp))
	fmt.Printf("%+v\n", tmp)
}

func main() {
	///
}
