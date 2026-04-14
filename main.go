package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type serverConfig struct {
	CCID    string
	Address string
}

type PRChainCode struct{}

func main() {
	fmt.Println("PR Chaincode Starting...")

	if os.Getenv("EXTERNAL_CHAINCODE") == "YES" {

		config := serverConfig{
			CCID:    os.Getenv("CHAINCODE_ID"),
			Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
		}

		if config.CCID == "" || config.Address == "" {
			log.Fatal("CHAINCODE_ID and CHAINCODE_SERVER_ADDRESS must be set")
		}

		server := &shim.ChaincodeServer{
			CCID:    config.CCID,
			Address: config.Address,
			CC:      &PRChainCode{},
			TLSProps: shim.TLSProperties{
				Disabled: true,
			},
		}

		if err := server.Start(); err != nil {
			log.Panicf("Error starting chaincode: %s", err)
		}

	} else {
		err := shim.Start(new(PRChainCode))
		if err != nil {
			log.Panicf("Error starting chaincode: %s", err)
		}
	}
}

func (t *PRChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("*****************************")
	fmt.Println("PRChainCode Init Started")
	fmt.Println("*****************************")

	_, args := stub.GetFunctionAndParameters()

	// If no args passed during instantiation
	if len(args) == 0 {
		return shim.Error("Init requires initial data (e.g. MSP mapping or bootstrap config)")
	}

	// Store init payload on ledger
	key := "INIT_DATA"

	err := stub.PutState(key, []byte(args[0]))
	if err != nil {
		return shim.Error("Failed to store init data: " + err.Error())
	}

	fmt.Println("Init data stored successfully")

	return shim.Success(nil)
}
