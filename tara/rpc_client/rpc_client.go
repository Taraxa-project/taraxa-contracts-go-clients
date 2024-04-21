package rpc_client

import (
	"context"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients_common"
	tara_rpc_types "github.com/Taraxa-project/taraxa-contracts-go-clients/tara/rpc_client/types"

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

func (RpcClient *RpcClient) GetPillarBlockData(period uint64, includeBinaryData bool) (*tara_rpc_types.PillarBlockData, error) {
	var pillarBlockData *tara_rpc_types.PillarBlockData
	err := RpcClient.EthClient.Client().CallContext(context.Background(), &pillarBlockData, "taraxa_getPillarBlockData", period, includeBinaryData)
	if err == nil && pillarBlockData == nil {
		err = ethereum.NotFound
	}

	return pillarBlockData, err
}
