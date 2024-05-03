package rpc_client

import (
	"context"

	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"
	tara_rpc_types "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client/types"

	"github.com/ethereum/go-ethereum"
)

type RpcClient struct {
	*clients_common.ContractClient
}

func NewRpcClient(config clients_common.NetConfig, communicationProtocol clients_common.CommunicationProtocol) (*RpcClient, error) {
	var err error

	rpcClient := new(RpcClient)
	rpcClient.ContractClient, err = clients_common.NewContractClient(config, communicationProtocol)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}

func NewSharedRpcClient(sharedClient *clients_common.ContractClient) *RpcClient {
	rpcClient := new(RpcClient)
	rpcClient.ContractClient = sharedClient

	return rpcClient
}

func (RpcClient *RpcClient) GetPillarBlockData(period uint64, includeBinaryData bool) (*tara_rpc_types.PillarBlockData, error) {
	var pillarBlockData *tara_rpc_types.PillarBlockData
	err := RpcClient.EthClient.Client().CallContext(context.Background(), &pillarBlockData, "taraxa_getPillarBlockData", period, includeBinaryData)
	if err == nil && pillarBlockData == nil {
		err = ethereum.NotFound
	}

	return pillarBlockData, err
}

func (RpcClient *RpcClient) GetTaraConfig() (*tara_rpc_types.TaraConfig, error) {
	var taraConfig *tara_rpc_types.TaraConfig
	err := RpcClient.EthClient.Client().CallContext(context.Background(), &taraConfig, "taraxa_getConfig")
	if err == nil && taraConfig == nil {
		err = ethereum.NotFound
	}

	return taraConfig, err
}
