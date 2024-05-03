package eth_net_config

import (
	"errors"
	"math/big"

	clients_common "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/common"
	"github.com/ethereum/go-ethereum/common"
)

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
