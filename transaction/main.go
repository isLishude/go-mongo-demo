package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sessionContext struct {
	context.Context
	mongo.Session
}

func contextWithSession(ctx context.Context, sess mongo.Session) mongo.SessionContext {
	return &sessionContext{
		Context: ctx,
		Session: sess,
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	uri := "mongodb://127.0.0.1:27017/?replSet=test"
	dbName := "test"
	colName := "test"

	cli, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(nil)

	db := cli.Database(dbName)
	col := db.Collection(colName)
	err = db.Drop(context.Background())
	checkError(err)

	// create collection
	if _, err = col.InsertOne(context.Background(), bson.M{"field": "value0"}); err != nil {
		log.Fatal(err)
	}

	{
		log.Println("Insert 1,2 and commit")
		ses, err := cli.StartSession()
		checkError(err)
		defer ses.EndSession(context.Background())

		mctx := contextWithSession(context.Background(), ses)

		err = ses.StartTransaction()
		checkError(err)

		_, err = col.InsertOne(mctx, bson.M{"field": "1"})
		checkError(err)
		_, err = col.InsertOne(mctx, bson.M{"field": "2"})
		checkError(err)

		err = ses.CommitTransaction(mctx)
		checkError(err)
	}
}
