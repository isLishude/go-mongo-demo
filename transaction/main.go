package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
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
	uri := "mongodb://127.0.0.1:27017"
	dbName := "test"
	colName := "test"
	ctx := context.Background()

	cli, err := mongo.Connect(ctx, uri)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)

	db := cli.Database(dbName)
	err = db.Drop(ctx)
	checkError(err)

	col := cli.Database(dbName).Collection(colName)
	// create collection
	if _, err = col.InsertOne(ctx, bson.M{"field": "value0"}); err != nil {
		log.Fatal(err)
	}

	// TODO FIX
	// Multi-document transactions are available for replica sets only.
	{
		ctx := context.Background()
		ses, err := cli.StartSession()
		checkError(err)
		defer ses.EndSession(ctx)

		mctx := contextWithSession(ctx, ses)

		err = ses.StartTransaction()
		checkError(err)
		_, err = col.InsertOne(mctx, bson.M{"field": "value1"})
		checkError(err)
		_, err = col.InsertOne(mctx, bson.M{"field": "value2"})
		checkError(err)
		_, err = col.InsertOne(mctx, bson.M{"field": "value3"})
		checkError(err)
		_, err = col.InsertOne(mctx, bson.M{"field": "value4"})
		checkError(err)
		err = ses.AbortTransaction(mctx)
		checkError(err)

		mctx2 := contextWithSession(ctx, ses)

		ses2, err := cli.StartSession()
		checkError(err)
		defer ses2.EndSession(ctx)

		_, err = col.InsertOne(mctx2, bson.M{"field": "value4"})
		checkError(err)
		err = ses.CommitTransaction(mctx)
		checkError(err)
	}
}
