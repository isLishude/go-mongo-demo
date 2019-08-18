package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// make sure that mongod version is ^4.0

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
	defer cli.Disconnect(context.Background())

	testDb := cli.Database(dbName)

	// make sure that colltion is created
	testDb.RunCommand(context.Background(), bson.M{"create": colName})
	col := testDb.Collection(colName)

	log.Println("Insert 1,2 and commit")

	// runTransactionWithRetry
	for {
		err := cli.UseSessionWithOptions(context.Background(),
			options.Session().SetDefaultReadPreference(readpref.Primary()),
			func(mctx mongo.SessionContext) error {
				if err := mctx.StartTransaction(options.Transaction().
					SetReadConcern(readconcern.Snapshot()).
					SetWriteConcern(writeconcern.New(writeconcern.WMajority()))); err != nil {
					mctx.AbortTransaction(mctx)
					return err
				}

				if _, err := col.InsertOne(mctx, bson.M{"field": "ACID-01"}); err != nil {
					mctx.AbortTransaction(mctx)
					return err
				}

				if _, err := col.InsertOne(mctx, bson.M{"field": "ACID-01"}); err != nil {
					mctx.AbortTransaction(mctx)
					return err
				}

				// commitWithRetry
				for {
					err = mctx.CommitTransaction(mctx)
					switch e := err.(type) {
					case nil:
						return nil
					case mongo.CommandError:
						if e.HasErrorLabel("UnknownTransactionCommitResult") {
							log.Println("UnknownTransactionCommitResult, retrying commit operation...")
							continue
						}
						log.Println("Error during commit...")
						return e
					default:
						log.Println("Error during commit...")
						return e
					}
				}
			})
		if err == nil {
			return
		}

		log.Println("Transaction aborted. Caught exception during transaction.")

		// If transient error, retry the whole transaction
		if cmdErr, ok := err.(mongo.CommandError); ok && cmdErr.HasErrorLabel("TransientTransactionError") {
			log.Println("TransientTransactionError, retrying transaction...")
			continue
		}
		log.Fatal(err)
	}
}
