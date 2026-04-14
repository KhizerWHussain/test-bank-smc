package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	return nil
}

func (s *SmartContract) Ping(ctx contractapi.TransactionContextInterface) string {
	return "OK"
}

func main() {
	cc, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		log.Panicf("error creating chaincode: %v", err)
	}

	// IMPORTANT: this keeps process running properly
	if err := cc.Start(); err != nil {
		log.Panicf("error starting chaincode: %v", err)
	}
}
