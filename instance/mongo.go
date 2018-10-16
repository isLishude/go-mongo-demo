package demoMongo

import (
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
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

func init() {

	MongoCtx, MongoCancel = context.WithCancel(context.Background())
	MongoClient, err := mongo.NewClient("mongodb://127.0.0.1:27017")

	if err != nil {
		log.Fatal(err)
	}

	if err := MongoClient.Connect(MongoCtx); err != nil {
		log.Fatal(err)
	}

	MongoDatabase = MongoClient.Database(mongoInstanceName)
	MongoTrxCollection = MongoDatabase.Collection(collectionName)
}
