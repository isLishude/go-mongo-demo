package demoTest

// Trx is transaction schema
type Trx struct {
	TxID      string `bson:"txid" json:"txid"`
	Height    uint32 `bson:"height" json:"height"`
	Timestamp uint32 `bson:"timestamp" json:"timestamp"`
	IsSuccess bool   `bson:"status" json:"status"`
	From      string `bson:"from" json:"from"`
	To        string `bson:"to" json:"to"`
}
