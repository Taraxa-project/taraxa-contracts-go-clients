package main

import (
	"log"

	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/tara_net_config"
)

func main() {
	log.Print("Tara rpc client demo")

	config, err := tara_net_config.GenNetConfig(clients_common.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	// var config clients_common.NetConfig
	// config.HttpUrl = "http://localhost:7017"
	// config.ChainID = big.NewInt(649)

	rpcClient, err := rpc_client.NewRpcClient(*config, clients_common.Http)
	if err != nil {
		log.Fatal(err)
	}

	pillarBlockData, err := rpcClient.GetPillarBlockData(100, true)
	if err != nil {
		log.Fatal("GetPillarBlockData err: ", err)
	} else {
		log.Printf("GetPillarBlockData: %d\n\n", pillarBlockData)
		log.Printf("pillarBlockData.PillarBlock.PbftPeriod: %d\n\n", pillarBlockData.PillarBlock.PbftPeriod)
	}

	taraConfig, err := rpcClient.GetTaraConfig()
	if err != nil {
		log.Fatal("GetTaraConfig err: ", err)
	} else {
		log.Printf("GetTaraConfig: %d\n\n", taraConfig)
		log.Printf("taraConfig.Hardforks.FicusHf.PillarBlocksInterval: %d\n\n", uint64(taraConfig.Hardforks.FicusHf.PillarBlocksInterval))
	}
}
