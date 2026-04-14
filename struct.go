package main

type Testdoc struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	Document  string `json:"document"`
	TxData    string `json:"txData"` // full stringified payload
	TxID      string `json:"txId"`
	Timestamp string `json:"timestamp"`
}
