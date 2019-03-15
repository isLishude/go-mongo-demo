package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	const (
		uri     = "mongodb://127.0.0.1:27017,127.0.0.1:27018/admin?replicaSet=test"
		dbName  = "test"
		colName = "test"
	)

	cli, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(nil)

	db := cli.Database(dbName)
	col := db.Collection(colName)
	err = db.Drop(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// create collection
	if _, err = col.InsertOne(context.Background(), bson.M{"field": "value0"}); err != nil {
		log.Fatal(err)
	}

	{
		log.Println("Insert 1,2 and commit")
		if err := cli.UseSession(context.Background(), func(mctx mongo.SessionContext) error {
			mctx.StartTransaction()

			if _, err := col.InsertOne(mctx, bson.M{"field": "ACID-01"}); err != nil {
				return err
			}

			if _, err := col.InsertOne(mctx, bson.M{"field": "ACID-01"}); err != nil {
				return err
			}
			return mctx.CommitTransaction(context.Background())
		}); err != nil {
			log.Panic(err)
		}
	}

	{
		log.Println("Insert 3,4 and abort")
		if err := cli.UseSession(context.Background(), func(mctx mongo.SessionContext) error {
			mctx.StartTransaction()

			if _, err := col.InsertOne(mctx, bson.M{"field": "ACID-03"}); err != nil {
				return err
			}

			if _, err := col.InsertOne(mctx, bson.M{"field": "ACID-04"}); err != nil {
				return err
			}
			return mctx.AbortTransaction(context.Background())
		}); err != nil {
			log.Panic(err)
		}
	}
}
