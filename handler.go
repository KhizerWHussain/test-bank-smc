package main

import (
	"encoding/json"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *PRChainCode) CreateTestDoc(stub hypConnect, args []string, txID string) pb.Response {

	if len(args) < 3 {
		return shim.Error("Expected id, key, document")
	}

	ID := sanitize(args[0], "string").(string)
	Key := sanitize(args[1], "string").(string)

	doc := Testdoc{
		ID:        ID,
		Key:       Key,
		Document:  "Testdoc",
		TxID:      txID,
		Timestamp: time.Now().String(),
	}

	bytes, _ := json.Marshal(doc)

	err := stub.Connection.PutState(doc.ID, bytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	RaiseEvent(stub, "TestDocCreated", doc)

	return shim.Success(bytes)
}

func (t *PRChainCode) CreateBulkTestDoc(stub hypConnect, args []string, txID string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Expecting JSON array string")
	}

	var docs []Testdoc

	err := json.Unmarshal([]byte(args[0]), &docs)
	if err != nil {
		return shim.Error(err.Error())
	}

	for _, d := range docs {

		d.TxID = txID
		d.Timestamp = time.Now().String()

		bytes, _ := json.Marshal(d)

		err := stub.Connection.PutState(d.ID, bytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	RaiseEvent(stub, "BulkTestDocCreated", docs)

	return shim.Success([]byte("Bulk Insert Success"))
}

func (t *PRChainCode) GetTestDoc(stub hypConnect, args []string, txID string) pb.Response {
	if len(args) < 1 {
		return shim.Error("ID required")
	}

	ID := sanitize(args[0], "string").(string)

	data, err := stub.Connection.GetState(ID)
	if err != nil || data == nil {
		return shim.Error("Document not found")
	}

	RaiseEvent(stub, "TestDocRetrieved", data)

	return shim.Success(data)
}
