package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoClient is project's client instance
var MongoClient mongo.Client

// MongoCtx is MongoDB top-level context
var MongoCtx context.Context

// MongoCancel is Mongodb instance top-level context.CancelFuc
var MongoCancel context.CancelFunc

// MongoDatabase is Mongodb database instance
var MongoDatabase *mongo.Database

// mongoInstanceName is
const mongoInstanceName = "test"

// MongoHistoryCollection is
var MongoHistoryCollection *mongo.Collection

// MongoTrxCollection is
var MongoTrxCollection *mongo.Collection

var collectionName = "trx"

var connectionString = "mongodb://127.0.0.1:27017"

// Trx is transaction schema
type Trx struct {
	TxID      string `bson:"txid" json:"txid"`
	Height    uint32 `bson:"height" json:"height"`
	Timestamp uint32 `bson:"timestamp" json:"timestamp"`
	IsSuccess bool   `bson:"status" json:"status"`
	From      string `bson:"from" json:"from"`
	To        string `bson:"to" json:"to"`
}

// CheckError is to check it error and exit
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	MongoCtx, MongoCancel = context.WithCancel(context.Background())
	MongoClient, err := mongo.Connect(MongoCtx, connectionString)
	defer MongoCancel()

	if err != nil {
		log.Fatal(err)
	}

	MongoDatabase = MongoClient.Database(mongoInstanceName)
	MongoTrxCollection = MongoDatabase.Collection(collectionName)
}
