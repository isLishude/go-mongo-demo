package main

import (
	"fmt"
)

// InsertManyDemo is for `insertMany`
func InsertManyDemo() {
	defer MongoCancel()
	var bulk []interface{}

	for i := 0; i < 10; i++ {
		tmp := Trx{
			TxID:      fmt.Sprintf("%d%d", i, i+100),
			Height:    uint32(i),
			IsSuccess: false,
			Timestamp: 0,
			From:      "a",
			To:        "b",
		}
		bulk = append(bulk, tmp)
	}

	_, err := MongoTrxCollection.InsertMany(MongoCtx, bulk)
	CheckError(err)
}

// InsertOneDemo is for `insertOne`
func InsertOneDemo() {
	defer MongoCancel()
	tmp := Trx{
		TxID:      "0",
		Height:    0,
		IsSuccess: false,
		Timestamp: 0,
		From:      "a",
		To:        "b",
	}
	_, err := MongoTrxCollection.InsertOne(MongoCtx, tmp)
	CheckError(err)
}
