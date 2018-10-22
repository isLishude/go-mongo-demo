package main

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

var dbName = "test"
var collectionName = "trx"
var connectString = "mongodb://127.0.0.1:27017"

var (
	// TrxColl is transaction collection
	TrxColl *mgo.Collection
)

// Trx is transaction schema
type Trx struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	TxID  string        `bson:"txid" json:"txid"`
	From  string        `bson:"from" json:"from"`
	To    string        `bson:"to" json:"to"`
	Value string        `bson:"value"`
}

// CheckError is to check it error and exit
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	session, err := mgo.Dial(connectString)
	CheckError(err)
	TrxColl = session.DB(dbName).C(collectionName)

	insertDemo()
}
