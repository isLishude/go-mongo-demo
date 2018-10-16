package main

import (
	"context"
	"fmt"
	"log"

	"github.com/islishude/demo/mongo/schema"
	"github.com/mongodb/mongo-go-driver/mongo"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(url)
	checkError(err)

	err = client.Connect(ctx)
	checkError(err)

	collection := client.Database(dbName).Collection(tableName)

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

	_, err = collection.InsertMany(ctx, bulk)

	checkError(err)
}

func insertOne() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(url)
	checkError(err)

	err = client.Connect(ctx)
	checkError(err)

	collection := client.Database(dbName).Collection(tableName)

	tmp := demoTest.Trx{
		TxID:      "0",
		Height:    0,
		IsSuccess: false,
		Timestamp: 0,
		From:      "a",
		To:        "b",
	}

	_, err = collection.InsertOne(ctx, tmp)

	checkError(err)
}

func main() {
	//
}
