package main

import (
	"log"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/client_base"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/eth_net_config"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Print("Bridge contract client demo")

	config, err := eth_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}
	config.ContractAddress = common.HexToAddress("0x762dA247D9F269b1689d4baaD587243eccF7910c")

	bridgeContractClient, err := bridge_contract_client.NewBridgeContractClient(*config, client_base.Http)
	if err != nil {
		log.Fatal(err)
	}

	stateWithProof, err := bridgeContractClient.GetStateWithProof()
	if err != nil {
		log.Print("GetStateWithProof err: ", err)
	} else {
		log.Printf("GetStateWithProof: %d\n\n", stateWithProof)
	}
}
