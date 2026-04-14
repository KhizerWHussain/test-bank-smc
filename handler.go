package main

import (
	"encoding/json"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// CREATE SINGLE DOC
func (t *PRChainCode) CreateTestDoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 3 {
		return shim.Error("Expected id, key, document")
	}

	txID := stub.GetTxID()

	doc := Testdoc{
		ID:        args[0],
		Key:       args[1],
		Document:  "Testdoc",
		TxID:      txID,
		Timestamp: time.Now().String(),
	}

	bytes, _ := json.Marshal(doc)

	err := stub.PutState(doc.ID, bytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(bytes)
}

// BULK INSERT (STRINGIFIED JSON ARRAY)
func (t *PRChainCode) CreateBulkTestDoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Expecting JSON array string")
	}

	var docs []Testdoc

	err := json.Unmarshal([]byte(args[0]), &docs)
	if err != nil {
		return shim.Error(err.Error())
	}

	txID := stub.GetTxID()

	for _, d := range docs {

		d.TxID = txID
		d.Timestamp = time.Now().String()
		d.TxData = args[0]

		bytes, _ := json.Marshal(d)

		err := stub.PutState(d.ID, bytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success([]byte("Bulk Insert Success"))
}

// GET SINGLE DOC
func (t *PRChainCode) GetTestDoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("ID required")
	}

	data, err := stub.GetState(args[0])
	if err != nil || data == nil {
		return shim.Error("Document not found")
	}

	return shim.Success(data)
}
