package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type serverConfig struct {
	CCID    string
	Address string
}

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
