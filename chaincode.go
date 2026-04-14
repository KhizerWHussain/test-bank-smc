package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *PRChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke triggered")

	function, _ := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is running for function::::: " + function)

	txID := stub.GetTxID()

	connection := hypConnect{}
	connection.Connection = stub

	args, errTrans := getArguments(stub)
	if errTrans != nil {
		return shim.Error(errTrans.Error())
	}

	switch function {

	case "CreateTestDoc":
		return t.CreateTestDoc(connection, args, txID)
	case "CreateBulkTestDoc":
		return t.CreateBulkTestDoc(connection, args, txID)
	case "GetTestDoc":
		return t.GetTestDoc(connection, args, txID)
	default:
		return defaultInvoke(function)
	}
}

func defaultInvoke(function string) pb.Response {
	return shim.Error("Received unknown function invocation ---> " + function)
}

func getArguments(stub shim.ChaincodeStubInterface) ([]string, error) {
	transMap, err := stub.GetTransient()
	if err != nil {
		return nil, err
	}
	if _, ok := transMap["PrivateArgs"]; !ok {
		return nil, errors.New("PrivateArgs must be a key in the transient map")
	}
	fmt.Println("Arguments:", transMap)
	generalInput := string(transMap["PrivateArgs"])
	retVal := strings.Split(generalInput, "|")
	return retVal, nil
}
