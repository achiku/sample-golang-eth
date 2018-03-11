package main

import (
	"log"

	web3 "github.com/regcostajr/go-web3"

	"github.com/regcostajr/go-web3/providers"
)

// BlockParameter
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
const (
	DefaultBlockParamLatest string = "latest"
)

func main() {
	c := web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, true))
	b, err := c.Eth.GetBalance("0x24fc5c1c97f838e57c944240fa2ffc18256bc415", DefaultBlockParamLatest)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", b)
}
