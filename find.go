package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo/bson"
)

func findDemo() {
	iter := TrxColl.Find(bson.M{"from": "b"}).Iter()
	defer iter.Close()

	tmp := Trx{}
	for iter.Next(&tmp) {
		fmt.Printf("%#v\n", tmp)
	}
}

func findOneDemo() {
	tmp := Trx{}
	err := TrxColl.Find(bson.M{"from": "a"}).One(&tmp)

	// if found nothing then reports error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", tmp)
}
