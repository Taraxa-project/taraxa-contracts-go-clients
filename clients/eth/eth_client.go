package eth_client

import (
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client"
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client"
	"github.com/ethereum/go-ethereum/common"

	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"
)

type EthClient struct {
	TaraClientContractClient *tara_client_contract_client.TaraClientContractClient
	BridgeContractClient     *bridge_contract_client.BridgeContractClient
}

type EthClientConfig struct {
	clients_common.NetConfig
	BridgeContractAddress     common.Address `json:"bridge_contract_address"`
	TaraClientContractAddress common.Address `json:"tara_client_contract_address"`
}

func NewEthClient(config EthClientConfig, communicationProtocol clients_common.CommunicationProtocol) (*EthClient, error) {
	sharedClient, err := clients_common.NewContractClient(config.NetConfig, communicationProtocol)
	if err != nil {
		return nil, err
	}

	ethClient := new(EthClient)
	ethClient.BridgeContractClient, err = bridge_contract_client.NewSharedBridgeContractClient(sharedClient, config.BridgeContractAddress)
	if err != nil {
		return nil, err
	}

	ethClient.TaraClientContractClient, err = tara_client_contract_client.NewSharedTaraClientContractClient(sharedClient, config.TaraClientContractAddress)
	if err != nil {
		return nil, err
	}

	return ethClient, nil
}
