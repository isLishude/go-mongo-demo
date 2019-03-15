package main

import (
	"context"
	"log"
	"time"

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

	col := cli.Database(dbName).Collection(colName)

	// updated doc should create by "$set" key
	newdoc := bson.M{"$set": bson.M{"filed": "new"}}

	// Update one
	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		col.UpdateOne(ctx, bson.M{"field": "update one"}, newdoc)
	}

	// upsert one, update one but if not exists insert
	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		col.UpdateOne(ctx, bson.M{"field": "upsert one"}, newdoc, options.Update().SetUpsert(true))
	}

	// update many
	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		col.UpdateMany(ctx, bson.M{"field": "update many"}, newdoc)
	}
}
