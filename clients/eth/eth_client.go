package eth_client

import (
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/client_base"
)

type EthClient struct {
	*client_base.ClientBase
	TaraClientContractClient *tara_client_contract_client.TaraClientContractClient
	BridgeContractClient     *bridge_contract_client.BridgeContractClient
}

type EthClientConfig struct {
	client_base.NetConfig
	BridgeContractAddress     common.Address `json:"bridge_contract_address"`
	TaraClientContractAddress common.Address `json:"tara_client_contract_address"`
}

func NewEthClient(config EthClientConfig, communicationProtocol client_base.CommunicationProtocol) (*EthClient, error) {
	var err error

	ethClient := new(EthClient)
	ethClient.ClientBase, err = client_base.NewClientBase(config.NetConfig, communicationProtocol)
	if err != nil {
		return nil, err
	}

	ethClient.BridgeContractClient, err = bridge_contract_client.NewSharedBridgeContractClient(ethClient.ClientBase, config.BridgeContractAddress)
	if err != nil {
		return nil, err
	}

	ethClient.TaraClientContractClient, err = tara_client_contract_client.NewSharedTaraClientContractClient(ethClient.ClientBase, config.TaraClientContractAddress)
	if err != nil {
		return nil, err
	}

	return ethClient, nil
}
