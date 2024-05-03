package main

import (
	"log"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client"
	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/eth_net_config"
)

func main() {
	log.Print("Bridge contract client demo")

	config, err := eth_net_config.GenNetConfig(clients_common.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	bridgeContractClient, err := bridge_contract_client.NewBridgeContractClient(*config, clients_common.Http)
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
