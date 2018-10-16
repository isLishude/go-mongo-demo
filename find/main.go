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

func find() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(url)
	checkError(err)

	err = client.Connect(ctx)
	checkError(err)

	collection := client.Database(dbName).Collection(tableName)

	cur, err := collection.Find(ctx, map[string]string{"from": "a"})
	checkError(err)
	defer cur.Close(context.Background())

	for cur.Next(ctx) {
		tmp := new(demoTest.Trx)
		checkError(cur.Decode(tmp))
		fmt.Printf("%+v\n", tmp)
	}

	checkError(cur.Err())
}

func findOne() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(url)
	checkError(err)

	err = client.Connect(ctx)
	checkError(err)

	collection := client.Database(dbName).Collection(tableName)

	result := collection.FindOne(ctx, map[string]string{"from": "a"})

	tmp := new(demoTest.Trx)
	checkError(result.Decode(tmp))
	fmt.Printf("%+v\n", tmp)
}

func main() {
	///
}
