// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tara_client_contract_interface

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CompactSignature is an auto generated low-level Go binding around an user-defined struct.
type CompactSignature struct {
	R  [32]byte
	Vs [32]byte
}

// PillarBlockFinalizationData is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockFinalizationData struct {
	Period     *big.Int
	StateRoot  [32]byte
	BridgeRoot [32]byte
	PrevHash   [32]byte
}

// PillarBlockFinalizedBlock is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockFinalizedBlock struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}

// PillarBlockVoteCountChange is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockVoteCountChange struct {
	Validator common.Address
	Change    int32
}

// PillarBlockWithChanges is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockWithChanges struct {
	Block            PillarBlockFinalizationData
	ValidatorChanges []PillarBlockVoteCountChange
}

// TaraClientContractInterfaceMetaData contains all meta data concerning the TaraClientContractInterface contract.
var TaraClientContractInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"change\",\"type\":\"int32\"}],\"internalType\":\"structPillarBlock.VoteCountChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"internalType\":\"structPillarBlock.WithChanges\",\"name\":\"_genesisBlock\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pillarBlockInterval\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"HashesNotMatching\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InvalidBlockInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"change\",\"type\":\"int32\"}],\"internalType\":\"structPillarBlock.VoteCountChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"internalType\":\"structPillarBlock.WithChanges[]\",\"name\":\"blocks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature[]\",\"name\":\"lastBlockSigs\",\"type\":\"tuple[]\"}],\"name\":\"finalizeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"finalizedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalized\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"finalizedAt\",\"type\":\"uint256\"}],\"internalType\":\"structPillarBlock.FinalizedBlock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedBridgeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"getSignaturesWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pillarBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"change\",\"type\":\"int32\"}],\"internalType\":\"structPillarBlock.VoteCountChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"name\":\"processValidatorChanges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorVoteCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TaraClientContractInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use TaraClientContractInterfaceMetaData.ABI instead.
var TaraClientContractInterfaceABI = TaraClientContractInterfaceMetaData.ABI

// TaraClientContractInterface is an auto generated Go binding around an Ethereum contract.
type TaraClientContractInterface struct {
	TaraClientContractInterfaceCaller     // Read-only binding to the contract
	TaraClientContractInterfaceTransactor // Write-only binding to the contract
	TaraClientContractInterfaceFilterer   // Log filterer for contract events
}

// TaraClientContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientContractInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaraClientContractInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaraClientContractInterfaceSession struct {
	Contract     *TaraClientContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TaraClientContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaraClientContractInterfaceCallerSession struct {
	Contract *TaraClientContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TaraClientContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaraClientContractInterfaceTransactorSession struct {
	Contract     *TaraClientContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TaraClientContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaraClientContractInterfaceRaw struct {
	Contract *TaraClientContractInterface // Generic contract binding to access the raw methods on
}

// TaraClientContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceCallerRaw struct {
	Contract *TaraClientContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// TaraClientContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceTransactorRaw struct {
	Contract *TaraClientContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaraClientContractInterface creates a new instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterface(address common.Address, backend bind.ContractBackend) (*TaraClientContractInterface, error) {
	contract, err := bindTaraClientContractInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterface{TaraClientContractInterfaceCaller: TaraClientContractInterfaceCaller{contract: contract}, TaraClientContractInterfaceTransactor: TaraClientContractInterfaceTransactor{contract: contract}, TaraClientContractInterfaceFilterer: TaraClientContractInterfaceFilterer{contract: contract}}, nil
}

// NewTaraClientContractInterfaceCaller creates a new read-only instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*TaraClientContractInterfaceCaller, error) {
	contract, err := bindTaraClientContractInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceCaller{contract: contract}, nil
}

// NewTaraClientContractInterfaceTransactor creates a new write-only instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TaraClientContractInterfaceTransactor, error) {
	contract, err := bindTaraClientContractInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceTransactor{contract: contract}, nil
}

// NewTaraClientContractInterfaceFilterer creates a new log filterer instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TaraClientContractInterfaceFilterer, error) {
	contract, err := bindTaraClientContractInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceFilterer{contract: contract}, nil
}

// bindTaraClientContractInterface binds a generic wrapper to an already deployed contract.
func bindTaraClientContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaraClientContractInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClientContractInterface *TaraClientContractInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClientContractInterface.Contract.TaraClientContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClientContractInterface *TaraClientContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.TaraClientContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClientContractInterface *TaraClientContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.TaraClientContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClientContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.contract.Transact(opts, method, params...)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32) block, uint256 finalizedAt)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) Finalized(opts *bind.CallOpts) (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "finalized")

	outstruct := new(struct {
		BlockHash   [32]byte
		Block       PillarBlockFinalizationData
		FinalizedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Block = *abi.ConvertType(out[1], new(PillarBlockFinalizationData)).(*PillarBlockFinalizationData)
	outstruct.FinalizedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32) block, uint256 finalizedAt)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClientContractInterface.Contract.Finalized(&_TaraClientContractInterface.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32) block, uint256 finalizedAt)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClientContractInterface.Contract.Finalized(&_TaraClientContractInterface.CallOpts)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32),uint256))
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) GetFinalized(opts *bind.CallOpts) (PillarBlockFinalizedBlock, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "getFinalized")

	if err != nil {
		return *new(PillarBlockFinalizedBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(PillarBlockFinalizedBlock)).(*PillarBlockFinalizedBlock)

	return out0, err

}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32),uint256))
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClientContractInterface.Contract.GetFinalized(&_TaraClientContractInterface.CallOpts)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32),uint256))
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClientContractInterface.Contract.GetFinalized(&_TaraClientContractInterface.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) GetFinalizedBridgeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "getFinalizedBridgeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetFinalizedBridgeRoot() ([32]byte, error) {
	return _TaraClientContractInterface.Contract.GetFinalizedBridgeRoot(&_TaraClientContractInterface.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetFinalizedBridgeRoot() ([32]byte, error) {
	return _TaraClientContractInterface.Contract.GetFinalizedBridgeRoot(&_TaraClientContractInterface.CallOpts)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) GetSignaturesWeight(opts *bind.CallOpts, h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "getSignaturesWeight", h, signatures)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.GetSignaturesWeight(&_TaraClientContractInterface.CallOpts, h, signatures)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.GetSignaturesWeight(&_TaraClientContractInterface.CallOpts, h, signatures)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) PillarBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "pillarBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.PillarBlockInterval(&_TaraClientContractInterface.CallOpts)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.PillarBlockInterval(&_TaraClientContractInterface.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) Threshold() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.Threshold(&_TaraClientContractInterface.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) Threshold() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.Threshold(&_TaraClientContractInterface.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) TotalWeight() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.TotalWeight(&_TaraClientContractInterface.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) TotalWeight() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.TotalWeight(&_TaraClientContractInterface.CallOpts)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) ValidatorVoteCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "validatorVoteCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.ValidatorVoteCounts(&_TaraClientContractInterface.CallOpts, arg0)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.ValidatorVoteCounts(&_TaraClientContractInterface.CallOpts, arg0)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x508a0785.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) FinalizeBlocks(opts *bind.TransactOpts, blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "finalizeBlocks", blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x508a0785.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizeBlocks(&_TaraClientContractInterface.TransactOpts, blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x508a0785.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizeBlocks(&_TaraClientContractInterface.TransactOpts, blocks, lastBlockSigs)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0x8e3aa66d.
//
// Solidity: function processValidatorChanges((address,int32)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) ProcessValidatorChanges(opts *bind.TransactOpts, validatorChanges []PillarBlockVoteCountChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "processValidatorChanges", validatorChanges)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0x8e3aa66d.
//
// Solidity: function processValidatorChanges((address,int32)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) ProcessValidatorChanges(validatorChanges []PillarBlockVoteCountChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.ProcessValidatorChanges(&_TaraClientContractInterface.TransactOpts, validatorChanges)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0x8e3aa66d.
//
// Solidity: function processValidatorChanges((address,int32)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) ProcessValidatorChanges(validatorChanges []PillarBlockVoteCountChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.ProcessValidatorChanges(&_TaraClientContractInterface.TransactOpts, validatorChanges)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) SetThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.SetThreshold(&_TaraClientContractInterface.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.SetThreshold(&_TaraClientContractInterface.TransactOpts, _threshold)
}
