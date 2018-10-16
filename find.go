package main

import (
	"context"
	"fmt"
)

// FindDemo is
func FindDemo() {
	cur, err := MongoTrxCollection.Find(MongoCtx, map[string]string{"from": "a"})
	CheckError(err)
	defer cur.Close(context.Background())

	for cur.Next(MongoCtx) {
		tmp := new(Trx)
		CheckError(cur.Decode(tmp))
		fmt.Printf("%+v\n", tmp)
	}

	CheckError(cur.Err())
}

// FindOneDemo is
func FindOneDemo() {
	result := MongoTrxCollection.FindOne(MongoCtx, map[string]string{"from": "a"})

	tmp := new(Trx)
	if err := result.Decode(tmp); err != nil {
		fmt.Println("[Mongo::FindOne] nothing found")
		return
	}
	fmt.Printf("%+v\n", tmp)
}
