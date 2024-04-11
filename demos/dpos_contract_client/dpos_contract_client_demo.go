package dpos_contract_client_demo

import (
	"log"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients_common"
	dpos_contract_client "github.com/Taraxa-project/taraxa-contracts-go-clients/tara/dpos_contract_client"
)

func Demo() {
	log.Print("Dpos client demo")

	// var config clients_common.NetConfig
	// config.HttpUrl = "http://localhost:7017"
	// config.ChainID = big.NewInt(649)
	// config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")

	config, err := dpos_contract_client.GenNetConfig(clients_common.Mainnet)
	if err != nil {
		log.Fatal(err)
	}

	dposContractClient, err := dpos_contract_client.NewDposContractClient(*config, clients_common.Http)
	if err != nil {
		log.Fatal(err)
	}

	totalEligibleVotesCount, err := dposContractClient.GetTotalEligibleVotesCount()
	if err != nil {
		log.Print("GetTotalEligibleVotesCount err: ", err)
	} else {
		log.Printf("GetTotalEligibleVotesCount: %d\n\n", totalEligibleVotesCount)
	}
}
