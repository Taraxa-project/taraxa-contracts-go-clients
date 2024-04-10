package dpos_contract_client_demo

import (
	"log"

	taraxa_contracts_go_clients "github.com/Taraxa-project/taraxa-contracts-go-clients"
)

func Demo() {
	log.Print("Dpos client demo")

	config, err := taraxa_contracts_go_clients.GenContractsClientsConfig(taraxa_contracts_go_clients.Mainnet)
	if err != nil {
		log.Fatal(err)
	}

	// var config taraxa_contracts_go_clients.Config
	// config.TaraxaNetworkHttpUrl = "http://localhost:7017"
	// config.TaraxaNetworkChainID = big.NewInt(649)
	// config.TaraxaDposContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")

	contractsClients, err := taraxa_contracts_go_clients.NewTaraxaContractsClients(*config)
	if err != nil {
		log.Fatal(err)
	}

	dposContractClient, err := contractsClients.NewDposContractClient(taraxa_contracts_go_clients.Http)
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
