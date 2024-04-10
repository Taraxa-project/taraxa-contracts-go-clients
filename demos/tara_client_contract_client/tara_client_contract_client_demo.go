package tara_client_contract_client_demo

import (
	"log"

	taraxa_contracts_go_clients "github.com/Taraxa-project/taraxa-contracts-go-clients"
)

func Demo() {
	log.Print("Tara client demo")

	config, err := taraxa_contracts_go_clients.GenContractsClientsConfig(taraxa_contracts_go_clients.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	contractsClients, err := taraxa_contracts_go_clients.NewTaraxaContractsClients(*config)
	if err != nil {
		log.Fatal(err)
	}

	tarClientContractClient, err := contractsClients.NewTaraClientContractClient(taraxa_contracts_go_clients.Http)
	if err != nil {
		log.Fatal(err)
	}

	pendingPillarBlock, err := tarClientContractClient.GetPendingPillarBlock()
	if err != nil {
		log.Print("GetPendingPillarBlock err: ", err)
	} else {
		log.Printf("GetPendingPillarBlock: %d\n\n", pendingPillarBlock)
	}
}
