package main

import (
	"log"

	"github.com/globalsign/mgo"
)

func createIndexDemo() {
	index := mgo.Index{
		// prefix name with dash (-) for descending order
		Key:        []string{"txid"},
		Unique:     true,
		Background: true,
	}

	if err := TrxColl.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}
}
