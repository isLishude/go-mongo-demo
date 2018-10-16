package main

import (
	"fmt"
)

// InsertManyDemo is for `insertMany`
func InsertManyDemo() {
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

	res, err := MongoTrxCollection.InsertMany(MongoCtx, bulk)
	CheckError(err)
	fmt.Printf("%#v", res)
}

// InsertOneDemo is for `insertOne`
func InsertOneDemo() {
	tmp := Trx{
		TxID:      "0",
		Height:    0,
		IsSuccess: false,
		Timestamp: 0,
		From:      "a",
		To:        "b",
	}
	res, err := MongoTrxCollection.InsertOne(MongoCtx, tmp)
	CheckError(err)
	fmt.Printf("%#v", res)
}
