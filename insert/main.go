package demoInsert

import (
	"fmt"

	"github.com/islishude/demo/mongo/helper"
	"github.com/islishude/demo/mongo/instance"
	"github.com/islishude/demo/mongo/schema"
)

func init() {
	fmt.Println("demo insert running")
}

var tableName = "trx"
var dbName = "test"
var url = "mongodb://127.0.0.1:27017"

// InsertManyDemo is for `insertMany`
func InsertManyDemo() {
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
	demoHelper.CheckError(err)
}

// InsertOneDemo is for `insertOne`
func InsertOneDemo() {
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
	demoHelper.CheckError(err)
}
