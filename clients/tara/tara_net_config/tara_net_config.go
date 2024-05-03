package tara_net_config

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
		config.HttpUrl = "https://rpc.mainnet.taraxa.io"
		config.WsUrl = "" // TODO
		config.ChainID = big.NewInt(841)
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	case clients_common.Testnet:
		config.HttpUrl = "https://rpc.testnet.taraxa.io"
		config.WsUrl = "" // TODO
		config.ChainID = big.NewInt(842)
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	case clients_common.Devnet:
		config.HttpUrl = "https://rpc.devnet.taraxa.io"
		config.WsUrl = "" // TODO
		config.ChainID = big.NewInt(843)
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	return config, nil
}
