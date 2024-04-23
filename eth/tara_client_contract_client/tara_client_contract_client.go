package tara_client_contract_client

import (
	"errors"
	"math/big"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/clients_common"
	tara_client_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/eth/tara_client_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TaraClientContractClient contains variables needed for communication with taraxa client smart contract on eth
type TaraClientContractClient struct {
	*clients_common.ContractClient
	taraClientInterface *tara_client_contract_interface.TaraClientContractInterface
}

func GenNetConfig(network clients_common.Network) (*clients_common.NetConfig, error) {
	config := new(clients_common.NetConfig)

	switch network {
	case clients_common.Mainnet:
		// TODO:
		return nil, errors.New("Mainnet not supported")
		break
	case clients_common.Testnet:
		config.HttpUrl = "https://holesky.drpc.org"
		config.ChainID = big.NewInt(17000)
		config.ContractAddress = common.HexToAddress("0x52a7c8db4a32016e4b8b6b4b44590c52079f32a9")
		break
	case clients_common.Devnet:
		// TODO
		return nil, errors.New("Devnet not supported")
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	return config, nil
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
