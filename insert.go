package main

import (
	"log"
)

func insertDemo() {
	tmp := Trx{
		TxID:  "0",
		From:  "a",
		To:    "b",
		Value: "1",
	}

	tmp2 := Trx{
		TxID:  "1",
		From:  "a",
		To:    "b",
		Value: "1",
	}

	err := TrxColl.Insert(&tmp, &tmp2)

	if err != nil {
		log.Fatal(err)
	}
}
