package demoIndex

import (
	"context"
	"fmt"
	"log"

	"github.com/islishude/demo/mongo/helper"
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
// TODO: wait for driver fix IndexView export bug
func IndexBuildDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(url)
	demoHelper.CheckError(err)

	err = client.Connect(ctx)
	demoHelper.CheckError(err)

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
