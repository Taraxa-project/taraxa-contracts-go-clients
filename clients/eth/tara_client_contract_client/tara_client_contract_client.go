package tara_client_contract_client

import (
	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"
	tara_client_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TaraClientContractClient contains variables needed for communication with taraxa client smart contract on eth
type TaraClientContractClient struct {
	*clients_common.ContractClient
	taraClientInterface *tara_client_contract_interface.TaraClientContractInterface
}

func NewTaraClientContractClient(config clients_common.NetConfig, communicationProtocol clients_common.CommunicationProtocol) (*TaraClientContractClient, error) {
	var err error

	taraClientContractClient := new(TaraClientContractClient)
	taraClientContractClient.ContractClient, err = clients_common.NewContractClient(config, communicationProtocol)
	if err != nil {
		return nil, err
	}

	taraClientContractClient.taraClientInterface, err = tara_client_contract_interface.NewTaraClientContractInterface(taraClientContractClient.Config.ContractAddress, taraClientContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient, nil
}

func NewSharedTaraClientContractClient(sharedClient *clients_common.ContractClient, contractAddress common.Address) (*TaraClientContractClient, error) {
	var err error

	taraClientContractClient := new(TaraClientContractClient)
	taraClientContractClient.ContractClient = sharedClient

	taraClientContractClient.taraClientInterface, err = tara_client_contract_interface.NewTaraClientContractInterface(contractAddress, taraClientContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient, nil
}

func (taraClientContractClient *TaraClientContractClient) GetFinalizedPillarBlock() (tara_client_contract_interface.PillarBlockFinalizedBlock, error) {
	return taraClientContractClient.taraClientInterface.GetFinalized(&bind.CallOpts{})
}

func (taraClientContractClient *TaraClientContractClient) FinalizeBlocks(transactor *clients_common.Transactor, blocks []tara_client_contract_interface.PillarBlockWithChanges, lastBlockSigs []tara_client_contract_interface.CompactSignature) (*types.Transaction, error) {
	transactOpts, err := taraClientContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient.taraClientInterface.FinalizeBlocks(transactOpts, blocks, lastBlockSigs)
}
