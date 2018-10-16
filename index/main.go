package demoIndex

import (
	"context"
	"fmt"
	"log"

	"github.com/islishude/demo/mongo/helper"
	"github.com/islishude/demo/mongo/instance"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func init() {
	fmt.Println("demo index running")
}

var tableName = "trx"
var dbName = "test"
var url = "mongodb://127.0.0.1:27017"

// IndexBuildDemo is
func IndexBuildDemo() {
	defer demoMongo.MongoCancel()

	index := mongo.IndexModel{
		Keys:    bson.NewDocument(bson.EC.Int32("txid", int32(1))),
		Options: mongo.NewIndexOptionsBuilder().Background(true).Build(),
	}

	res, err := demoMongo.MongoTrxCollection.Indexes().CreateOne(demoMongo.MongoCtx, index)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

// GetIndexes is
func GetIndexes() {
	defer demoMongo.MongoCancel()

	cur, err := demoMongo.MongoTrxCollection.Indexes().List(demoMongo.MongoCtx)
	demoHelper.CheckError(err)

	defer cur.Close(demoMongo.MongoCtx)

	res := make(map[string]interface{})
	for cur.Next(context.TODO()) {
		cur.Decode(&res)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}
