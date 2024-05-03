package bridge_contract_client

import (
	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BridgeContractClient struct {
	*clients_common.ContractClient
	bridgeInterface *bridge_contract_interface.BridgeContractInterface
}

func NewBridgeContractClient(config clients_common.NetConfig, communicationProtocol clients_common.CommunicationProtocol) (*BridgeContractClient, error) {
	var err error

	bridgeContractClient := new(BridgeContractClient)
	bridgeContractClient.ContractClient, err = clients_common.NewContractClient(config, communicationProtocol)
	if err != nil {
		return nil, err
	}

	// TODO: fix config
	bridgeContractClient.bridgeInterface, err = bridge_contract_interface.NewBridgeContractInterface(bridgeContractClient.Config.ContractAddress, bridgeContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return bridgeContractClient, nil
}

func NewSharedBridgeContractClient(sharedClient *clients_common.ContractClient, contractAddress common.Address) (*BridgeContractClient, error) {
	var err error

	bridgeContractClient := new(BridgeContractClient)
	bridgeContractClient.ContractClient = sharedClient

	bridgeContractClient.bridgeInterface, err = bridge_contract_interface.NewBridgeContractInterface(contractAddress, bridgeContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return bridgeContractClient, nil
}

func (BridgeContractClient *BridgeContractClient) GetStateWithProof() (bridge_contract_interface.SharedStructsStateWithProof, error) {
	return BridgeContractClient.bridgeInterface.GetStateWithProof(&bind.CallOpts{})
}

func (BridgeContractClient *BridgeContractClient) ApplyState(transactor *clients_common.Transactor, stateWithProof bridge_contract_interface.SharedStructsStateWithProof) (*types.Transaction, error) {
	transactOpts, err := BridgeContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return BridgeContractClient.bridgeInterface.ApplyState(transactOpts, stateWithProof)
}
