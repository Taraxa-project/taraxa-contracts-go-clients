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

// PillarBlockWeightChange is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockWeightChange struct {
	Validator common.Address
	Change    *big.Int
}

// PillarBlockWithChanges is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockWithChanges struct {
	Block            PillarBlockFinalizationData
	ValidatorChanges []PillarBlockWeightChange
}

// TaraClientContractInterfaceMetaData contains all meta data concerning the TaraClientContractInterface contract.
var TaraClientContractInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int96\",\"name\":\"change\",\"type\":\"int96\"}],\"internalType\":\"structPillarBlock.WeightChange[]\",\"name\":\"_validatorChanges\",\"type\":\"tuple[]\"},{\"internalType\":\"int256\",\"name\":\"_threshold\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"addPendingBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"finalizeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"finalizePendingBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"finalizedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedBridgeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPending\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int96\",\"name\":\"change\",\"type\":\"int96\"}],\"internalType\":\"structPillarBlock.WeightChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"internalType\":\"structPillarBlock.WithChanges\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"getSignaturesWeight\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"weight\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int96\",\"name\":\"change\",\"type\":\"int96\"}],\"internalType\":\"structPillarBlock.WeightChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"name\":\"processValidatorChanges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_threshold\",\"type\":\"int256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"int96\",\"name\":\"\",\"type\":\"int96\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(((uint256,bytes32,bytes32,bytes32),(address,int96)[]))
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) GetPending(opts *bind.CallOpts) (PillarBlockWithChanges, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "getPending")

	if err != nil {
		return *new(PillarBlockWithChanges), err
	}

	out0 := *abi.ConvertType(out[0], new(PillarBlockWithChanges)).(*PillarBlockWithChanges)

	return out0, err

}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(((uint256,bytes32,bytes32,bytes32),(address,int96)[]))
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetPending() (PillarBlockWithChanges, error) {
	return _TaraClientContractInterface.Contract.GetPending(&_TaraClientContractInterface.CallOpts)
}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(((uint256,bytes32,bytes32,bytes32),(address,int96)[]))
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetPending() (PillarBlockWithChanges, error) {
	return _TaraClientContractInterface.Contract.GetPending(&_TaraClientContractInterface.CallOpts)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(int256 weight)
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
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(int256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.GetSignaturesWeight(&_TaraClientContractInterface.CallOpts, h, signatures)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(int256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.GetSignaturesWeight(&_TaraClientContractInterface.CallOpts, h, signatures)
}

// PendingHash is a free data retrieval call binding the contract method 0x628449fd.
//
// Solidity: function pendingHash() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) PendingHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "pendingHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingHash is a free data retrieval call binding the contract method 0x628449fd.
//
// Solidity: function pendingHash() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) PendingHash() ([32]byte, error) {
	return _TaraClientContractInterface.Contract.PendingHash(&_TaraClientContractInterface.CallOpts)
}

// PendingHash is a free data retrieval call binding the contract method 0x628449fd.
//
// Solidity: function pendingHash() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) PendingHash() ([32]byte, error) {
	return _TaraClientContractInterface.Contract.PendingHash(&_TaraClientContractInterface.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(int96)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(int96)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) Validators(arg0 common.Address) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.Validators(&_TaraClientContractInterface.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(int96)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) Validators(arg0 common.Address) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.Validators(&_TaraClientContractInterface.CallOpts, arg0)
}

// AddPendingBlock is a paid mutator transaction binding the contract method 0xdc8ec114.
//
// Solidity: function addPendingBlock(bytes b) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) AddPendingBlock(opts *bind.TransactOpts, b []byte) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "addPendingBlock", b)
}

// AddPendingBlock is a paid mutator transaction binding the contract method 0xdc8ec114.
//
// Solidity: function addPendingBlock(bytes b) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) AddPendingBlock(b []byte) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.AddPendingBlock(&_TaraClientContractInterface.TransactOpts, b)
}

// AddPendingBlock is a paid mutator transaction binding the contract method 0xdc8ec114.
//
// Solidity: function addPendingBlock(bytes b) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) AddPendingBlock(b []byte) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.AddPendingBlock(&_TaraClientContractInterface.TransactOpts, b)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x0d159634.
//
// Solidity: function finalizeBlock(bytes b, (bytes32,bytes32)[] signatures) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) FinalizeBlock(opts *bind.TransactOpts, b []byte, signatures []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "finalizeBlock", b, signatures)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x0d159634.
//
// Solidity: function finalizeBlock(bytes b, (bytes32,bytes32)[] signatures) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) FinalizeBlock(b []byte, signatures []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizeBlock(&_TaraClientContractInterface.TransactOpts, b, signatures)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x0d159634.
//
// Solidity: function finalizeBlock(bytes b, (bytes32,bytes32)[] signatures) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) FinalizeBlock(b []byte, signatures []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizeBlock(&_TaraClientContractInterface.TransactOpts, b, signatures)
}

// FinalizePendingBlock is a paid mutator transaction binding the contract method 0x1f2b13c0.
//
// Solidity: function finalizePendingBlock((bytes32,bytes32)[] signatures) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) FinalizePendingBlock(opts *bind.TransactOpts, signatures []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "finalizePendingBlock", signatures)
}

// FinalizePendingBlock is a paid mutator transaction binding the contract method 0x1f2b13c0.
//
// Solidity: function finalizePendingBlock((bytes32,bytes32)[] signatures) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) FinalizePendingBlock(signatures []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizePendingBlock(&_TaraClientContractInterface.TransactOpts, signatures)
}

// FinalizePendingBlock is a paid mutator transaction binding the contract method 0x1f2b13c0.
//
// Solidity: function finalizePendingBlock((bytes32,bytes32)[] signatures) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) FinalizePendingBlock(signatures []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizePendingBlock(&_TaraClientContractInterface.TransactOpts, signatures)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0xf7ee9808.
//
// Solidity: function processValidatorChanges((address,int96)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) ProcessValidatorChanges(opts *bind.TransactOpts, validatorChanges []PillarBlockWeightChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "processValidatorChanges", validatorChanges)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0xf7ee9808.
//
// Solidity: function processValidatorChanges((address,int96)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) ProcessValidatorChanges(validatorChanges []PillarBlockWeightChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.ProcessValidatorChanges(&_TaraClientContractInterface.TransactOpts, validatorChanges)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0xf7ee9808.
//
// Solidity: function processValidatorChanges((address,int96)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) ProcessValidatorChanges(validatorChanges []PillarBlockWeightChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.ProcessValidatorChanges(&_TaraClientContractInterface.TransactOpts, validatorChanges)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xef0fb558.
//
// Solidity: function setThreshold(int256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) SetThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xef0fb558.
//
// Solidity: function setThreshold(int256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.SetThreshold(&_TaraClientContractInterface.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xef0fb558.
//
// Solidity: function setThreshold(int256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.SetThreshold(&_TaraClientContractInterface.TransactOpts, _threshold)
}
