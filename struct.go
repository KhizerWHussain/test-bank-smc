package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type Testdoc struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	Document  string `json:"document"`
	TxID      string `json:"txId"`
	Timestamp string `json:"timestamp"`
}

type generalEventStruct struct {
	EventName      string            `json:"eventName"`
	EventList      []eventDataFormat `json:"events"`
	AdditionalData interface{}       `json:"additionalData"`
}

type eventDataFormat struct {
	Key        string `json:"Key"`
	Collection string `json:"Collection"`
}

type hypConnect struct {
	Connection shim.ChaincodeStubInterface
	EventList  []eventDataFormat
}
