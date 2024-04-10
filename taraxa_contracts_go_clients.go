package taraxa_contracts_go_clients

import (
	"errors"
	"log"
	"math/big"

	"github.com/Taraxa-project/taraxa-contracts-go-clients/tara/dpos_contract_client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Network uint8

const (
	Mainnet = iota
	Testnet
	Devnet
)

type CommunicationProtocol uint8

const (
	Http = iota
	WebSocket
)

type Config struct {
	TaraxaNetworkHttpUrl      string
	TaraxaNetworkWsUrl        string
	TaraxaNetworkChainID      *big.Int
	TaraxaDposContractAddress common.Address
	EhereumNetworkUrl         string
}

// Creates new GenContractsClientsConfig based on default values set for different taraxa networks
func GenContractsClientsConfig(network Network) (*Config, error) {
	config := new(Config)

	switch network {
	case Mainnet:
		config.TaraxaNetworkHttpUrl = "https://rpc.mainnet.taraxa.io"
		config.TaraxaNetworkWsUrl = "" // TODO
		config.TaraxaNetworkChainID = big.NewInt(841)
		config.TaraxaDposContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		config.EhereumNetworkUrl = "" // TODO
		break
	case Testnet:
		config.TaraxaNetworkHttpUrl = "https://rpc.testnet.taraxa.io"
		config.TaraxaNetworkWsUrl = "" // TODO
		config.TaraxaNetworkChainID = big.NewInt(842)
		config.TaraxaDposContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		config.EhereumNetworkUrl = "" // TODO
		break
	case Devnet:
		config.TaraxaNetworkHttpUrl = "https://rpc.devnet.taraxa.io"
		config.TaraxaNetworkWsUrl = "" // TODO
		config.TaraxaNetworkChainID = big.NewInt(843)
		config.TaraxaDposContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		config.EhereumNetworkUrl = "" // TODO
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	return config, nil
}

type TaraxaContractsClients struct {
	config Config

	ethClient  *ethclient.Client
	taraClient *ethclient.Client

	//DposContractClient      *dpos_contract_client.DposContractClient
}

func NewTaraxaContractsClients(config Config) (*TaraxaContractsClients, error) {
	taraxaContractsClients := new(TaraxaContractsClients)
	taraxaContractsClients.config = config

	return taraxaContractsClients, nil
}

func (taraxaContractsClients *TaraxaContractsClients) NewDposContractClient(communicationProtocol CommunicationProtocol) (*dpos_contract_client.DposContractClient, error) {
	var err error
	var networkUrl string

	switch communicationProtocol {
	case Http:
		if taraxaContractsClients.config.TaraxaNetworkHttpUrl == "" {
			return nil, errors.New("config.TaraxaNetworkHttpUrl not configured")
		}
		networkUrl = taraxaContractsClients.config.TaraxaNetworkHttpUrl
		break
	case WebSocket:
		if taraxaContractsClients.config.TaraxaNetworkWsUrl == "" {
			return nil, errors.New("config.TaraxaNetworkWsUrl not configured")
		}
		networkUrl = taraxaContractsClients.config.TaraxaNetworkWsUrl
		break
	default:
		return nil, errors.New("invalid communicationProtocol argument")
	}

	// Create tara client in case it was not created yet
	if taraxaContractsClients.taraClient == nil {
		taraxaContractsClients.taraClient, err = ethclient.Dial(networkUrl)
		if err != nil {
			return nil, err
		}
	}
	log.Println("tu sme ", networkUrl, taraxaContractsClients.config.TaraxaDposContractAddress, taraxaContractsClients.config.TaraxaNetworkChainID)

	dposContractClient, err := dpos_contract_client.NewDposContractClient(taraxaContractsClients.taraClient, taraxaContractsClients.config.TaraxaDposContractAddress, taraxaContractsClients.config.TaraxaNetworkChainID)
	if err != nil {
		return nil, err
	}

	return dposContractClient, nil
}
