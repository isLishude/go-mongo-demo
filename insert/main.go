package main

import (
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

func insertMany() {
	defer demoMongo.MongoCancel()
	var bulk []interface{}

	for i := 0; i < 10; i++ {
		tmp := demoTest.Trx{
			TxID:      fmt.Sprintf("%d%d", i, i+100),
			Height:    uint32(i),
			IsSuccess: false,
			Timestamp: 0,
			From:      "a",
			To:        "b",
		}
		bulk = append(bulk, tmp)
	}

	_, err := demoMongo.MongoTrxCollection.InsertMany(demoMongo.MongoCtx, bulk)
	checkError(err)
}

func insertOne() {
	defer demoMongo.MongoCancel()
	tmp := demoTest.Trx{
		TxID:      "0",
		Height:    0,
		IsSuccess: false,
		Timestamp: 0,
		From:      "a",
		To:        "b",
	}
	_, err := demoMongo.MongoTrxCollection.InsertOne(demoMongo.MongoCtx, tmp)
	checkError(err)
}

func main() {
	//
}
