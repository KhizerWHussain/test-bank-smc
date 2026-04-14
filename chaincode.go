package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *PRChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Invoke triggered")

	function, args := stub.GetFunctionAndParameters()

	switch function {

	case "CreateTestDoc":
		return t.CreateTestDoc(stub, args)
	case "CreateBulkTestDoc":
		return t.CreateBulkTestDoc(stub, args)
	case "GetTestDoc":
		return t.GetTestDoc(stub, args)
	default:
		return defaultInvoke(function)
	}
}

func defaultInvoke(function string) pb.Response {
	return shim.Error("Received unknown function invocation ---> " + function)
}
