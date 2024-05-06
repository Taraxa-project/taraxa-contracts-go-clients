package tara_client_contract_client

import (
	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients/client_base"
	tara_client_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TaraClientContractClient contains variables needed for communication with taraxa client smart contract on eth
type TaraClientContractClient struct {
	*client_base.ClientBase
	taraClientInterface *tara_client_contract_interface.TaraClientContractInterface
}

func NewTaraClientContractClient(config client_base.NetConfig, communicationProtocol client_base.CommunicationProtocol) (*TaraClientContractClient, error) {
	var err error

	taraClientContractClient := new(TaraClientContractClient)
	taraClientContractClient.ClientBase, err = client_base.NewClientBase(config, communicationProtocol)
	if err != nil {
		return nil, err
	}

	taraClientContractClient.taraClientInterface, err = tara_client_contract_interface.NewTaraClientContractInterface(taraClientContractClient.Config.ContractAddress, taraClientContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient, nil
}

func NewSharedTaraClientContractClient(clientBase *client_base.ClientBase, contractAddress common.Address) (*TaraClientContractClient, error) {
	var err error

	taraClientContractClient := new(TaraClientContractClient)
	taraClientContractClient.ClientBase = clientBase

	taraClientContractClient.taraClientInterface, err = tara_client_contract_interface.NewTaraClientContractInterface(contractAddress, taraClientContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient, nil
}

func (taraClientContractClient *TaraClientContractClient) GetFinalizedPillarBlock() (tara_client_contract_interface.PillarBlockFinalizedBlock, error) {
	return taraClientContractClient.taraClientInterface.GetFinalized(&bind.CallOpts{})
}

func (taraClientContractClient *TaraClientContractClient) FinalizeBlocks(transactor *client_base.Transactor, blocks []tara_client_contract_interface.PillarBlockWithChanges, lastBlockSigs []tara_client_contract_interface.CompactSignature) (*types.Transaction, error) {
	transactOpts, err := taraClientContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient.taraClientInterface.FinalizeBlocks(transactOpts, blocks, lastBlockSigs)
}
