package main

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

func updateOneDemo() {

	err := TrxColl.Update(bson.M{"txid": "a"}, bson.M{"$set": bson.M{"value": "2"}})

	if err != nil {
		log.Fatal(err)
	}
}

func upsertDemo() {
	_, err := TrxColl.Upsert(
		bson.M{"txid": "c"},
		bson.M{"from": "upsert", "to": "upsert", "value": "upsert"},
	)

	if err != nil {
		log.Fatal(err)
	}
}
