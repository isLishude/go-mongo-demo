package demoFind

import (
	"context"
	"fmt"

	"github.com/islishude/demo/mongo/helper"
	"github.com/islishude/demo/mongo/instance"
	"github.com/islishude/demo/mongo/schema"
)

func init() {
	fmt.Println("demo find running")
}

var tableName = "trx"
var dbName = "test"
var url = "mongodb://127.0.0.1:27017"

// FindDemo is
func FindDemo() {
	defer demoMongo.MongoCancel()
	cur, err := demoMongo.MongoTrxCollection.Find(demoMongo.MongoCtx, map[string]string{"from": "a"})
	demoHelper.CheckError(err)
	defer cur.Close(context.Background())

	for cur.Next(demoMongo.MongoCtx) {
		tmp := new(demoTest.Trx)
		demoHelper.CheckError(cur.Decode(tmp))
		fmt.Printf("%+v\n", tmp)
	}

	demoHelper.CheckError(cur.Err())
}

// FindOneDemo is
func FindOneDemo() {
	defer demoMongo.MongoCancel()
	result := demoMongo.MongoTrxCollection.FindOne(demoMongo.MongoCtx, map[string]string{"from": "a"})

	tmp := new(demoTest.Trx)
	if err := result.Decode(tmp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", tmp)
}
