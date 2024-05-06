package main

import (
	"log"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/client_base"
	tara_client "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/tara_net_config"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Print("Tara client demo")

	netConfig, err := tara_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	var config tara_client.TaraClientConfig
	config.NetConfig = *netConfig
	config.BridgeContractAddress = common.HexToAddress("0x0")
	config.DposContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")

	taraClient, err := tara_client.NewTaraClient(config, client_base.WebSocket)
	if err != nil {
		log.Fatal(err)
	}

	totalEligibleVotesCount, err := taraClient.DposContractClient.GetTotalEligibleVotesCount()
	if err != nil {
		log.Print("GetTotalEligibleVotesCount err: ", err)
	} else {
		log.Printf("GetTotalEligibleVotesCount: %d\n\n", totalEligibleVotesCount)
	}

	stateWithProof, err := taraClient.BridgeContractClient.GetStateWithProof()
	if err != nil {
		log.Print("GetStateWithProof err: ", err)
	} else {
		log.Printf("GetStateWithProof: %d\n\n", stateWithProof)
	}
}
