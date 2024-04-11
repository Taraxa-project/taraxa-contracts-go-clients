package main

import (
	"log"

	tara_client_contract_client "github.com/Taraxa-project/taraxa-contracts-go-clients/demos/tara_client_contract_client"
)

func main() {
	log.Print("Taraxa contracts clients demos")

	tara_client_contract_client.Demo()

	//dpos_contract_client.Demo()

}
