package tara_client

import (
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/dpos_contract_client"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client"
	"github.com/ethereum/go-ethereum/common"

	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"
)

type TaraClient struct {
	DposContractClient   *dpos_contract_client.DposContractClient
	BridgeContractClient *bridge_contract_client.BridgeContractClient
	RpcClient            *rpc_client.RpcClient
}

type TaraClientConfig struct {
	clients_common.NetConfig
	BridgeContractAddress common.Address `json:"bridge_contract_address"`
	DposContractAddress   common.Address `json:"dpos_contract_address"`
}

func NewTaraClient(config TaraClientConfig, communicationProtocol clients_common.CommunicationProtocol) (*TaraClient, error) {
	sharedClient, err := clients_common.NewContractClient(config.NetConfig, communicationProtocol)
	if err != nil {
		return nil, err
	}

	taraClient := new(TaraClient)
	taraClient.BridgeContractClient, err = bridge_contract_client.NewSharedBridgeContractClient(sharedClient, config.BridgeContractAddress)
	if err != nil {
		return nil, err
	}

	taraClient.DposContractClient, err = dpos_contract_client.NewSharedDposContractClient(sharedClient, config.DposContractAddress)
	if err != nil {
		return nil, err
	}

	taraClient.RpcClient = rpc_client.NewSharedRpcClient(sharedClient)

	return taraClient, nil
}
