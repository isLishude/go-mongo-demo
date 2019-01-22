package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

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

	col := cli.Database(dbName).Collection(colName)

	// TODO: fix bugs
	{
		ctx := context.Background()
		ses, err := cli.StartSession()
		defer ses.EndSession(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if err := ses.StartTransaction(); err != nil {
			log.Fatal(err)
		}
		if err := mongo.WithSession(ctx, ses, func(mctx mongo.SessionContext) error {
			_, err := col.InsertOne(mctx, bson.M{"field": "value0"})
			return err
		}); err != nil {
			log.Println("abort", err)
			ses.AbortTransaction(ctx)
		}
	}
}
