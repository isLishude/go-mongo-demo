package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
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

// TODO: wait for driver fix IndexView export bug
func indexBuild() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(url)
	checkError(err)

	err = client.Connect(ctx)
	checkError(err)

	// collection := client.Database(dbName).Collection(tableName)

	index := mongo.IndexModel{
		Keys:    bson.NewDocument(bson.EC.Int32("txid", int32(1))),
		Options: mongo.NewIndexOptionsBuilder().Background(true).Build(),
	}

	// @TODO it's driver bug
	// mongo.IndexView{Coll: collection}
	res, err := new(mongo.IndexView).CreateOne(ctx, index)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func main() {
	//
}
