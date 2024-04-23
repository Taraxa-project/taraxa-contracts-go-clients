package main

import (
	"log"
	"math/big"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients_common"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/tara/rpc_client"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Print("Tara rpc client demo")

	// config, err := net_config.GenNetConfig(clients_common.Mainnet)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var config clients_common.NetConfig
	config.HttpUrl = "http://localhost:7017"
	config.ChainID = big.NewInt(649)
	config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")

	rpcClient, err := rpc_client.NewRpcClient(config, clients_common.Http)
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
