package tara_client_contract_client

import (
	tara_client_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/eth/tara_client_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TaraClientContractClient contains variables needed for communication with taraxa client smart contract on eth
type TaraClientContractClient struct {
	taraClientInterface *tara_client_contract_interface.TaraClientContractInterface
	ethClient           *ethclient.Client
	// chainID       *big.Int
}

type Transactor struct {
	TransactOpts *bind.TransactOpts
	Address      common.Address
	Nonce        uint64
}

func NewTaraClientContractClient(ethClient *ethclient.Client, taraClientContractAddress common.Address) (*TaraClientContractClient, error) {
	taraClientContractClient := new(TaraClientContractClient)
	var err error

	taraClientContractClient.taraClientInterface, err = tara_client_contract_interface.NewTaraClientContractInterface(taraClientContractAddress, ethClient)
	taraClientContractClient.ethClient = ethClient
	// taraClient.chainID = chainID
	if err != nil {
		return nil, err
	}

	return taraClientContractClient, nil
}

func (taraClientContractClient *TaraClientContractClient) GetPendingPillarBlock() (tara_client_contract_interface.PillarBlockWithChanges, error) {
	return taraClientContractClient.taraClientInterface.GetPending(&bind.CallOpts{})
}
