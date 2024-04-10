// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pillar_block_interface

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

// PillarBlockSignedVote is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockSignedVote struct {
	Vote      PillarBlockVote
	Signature CompactSignature
}

// PillarBlockVote is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockVote struct {
	Period    *big.Int
	BlockHash [32]byte
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

// PillarBlockInterfaceMetaData contains all meta data concerning the PillarBlockInterface contract.
var PillarBlockInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"fromBytes\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int96\",\"name\":\"change\",\"type\":\"int96\"}],\"internalType\":\"structPillarBlock.WeightChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"internalType\":\"structPillarBlock.WithChanges\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.Vote\",\"name\":\"b\",\"type\":\"tuple\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int96\",\"name\":\"change\",\"type\":\"int96\"}],\"internalType\":\"structPillarBlock.WeightChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"internalType\":\"structPillarBlock.WithChanges\",\"name\":\"b\",\"type\":\"tuple\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.Vote\",\"name\":\"vote\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structPillarBlock.SignedVote\",\"name\":\"b\",\"type\":\"tuple\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bh\",\"type\":\"bytes32\"}],\"name\":\"getVoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int96\",\"name\":\"change\",\"type\":\"int96\"}],\"internalType\":\"structPillarBlock.WeightChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"internalType\":\"structPillarBlock.WithChanges\",\"name\":\"b\",\"type\":\"tuple\"}],\"name\":\"getVoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PillarBlockInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use PillarBlockInterfaceMetaData.ABI instead.
var PillarBlockInterfaceABI = PillarBlockInterfaceMetaData.ABI

// PillarBlockInterface is an auto generated Go binding around an Ethereum contract.
type PillarBlockInterface struct {
	PillarBlockInterfaceCaller     // Read-only binding to the contract
	PillarBlockInterfaceTransactor // Write-only binding to the contract
	PillarBlockInterfaceFilterer   // Log filterer for contract events
}

// PillarBlockInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PillarBlockInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PillarBlockInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PillarBlockInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PillarBlockInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PillarBlockInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PillarBlockInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PillarBlockInterfaceSession struct {
	Contract     *PillarBlockInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PillarBlockInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PillarBlockInterfaceCallerSession struct {
	Contract *PillarBlockInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PillarBlockInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PillarBlockInterfaceTransactorSession struct {
	Contract     *PillarBlockInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PillarBlockInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PillarBlockInterfaceRaw struct {
	Contract *PillarBlockInterface // Generic contract binding to access the raw methods on
}

// PillarBlockInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PillarBlockInterfaceCallerRaw struct {
	Contract *PillarBlockInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PillarBlockInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PillarBlockInterfaceTransactorRaw struct {
	Contract *PillarBlockInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPillarBlockInterface creates a new instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterface(address common.Address, backend bind.ContractBackend) (*PillarBlockInterface, error) {
	contract, err := bindPillarBlockInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterface{PillarBlockInterfaceCaller: PillarBlockInterfaceCaller{contract: contract}, PillarBlockInterfaceTransactor: PillarBlockInterfaceTransactor{contract: contract}, PillarBlockInterfaceFilterer: PillarBlockInterfaceFilterer{contract: contract}}, nil
}

// NewPillarBlockInterfaceCaller creates a new read-only instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PillarBlockInterfaceCaller, error) {
	contract, err := bindPillarBlockInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterfaceCaller{contract: contract}, nil
}

// NewPillarBlockInterfaceTransactor creates a new write-only instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PillarBlockInterfaceTransactor, error) {
	contract, err := bindPillarBlockInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterfaceTransactor{contract: contract}, nil
}

// NewPillarBlockInterfaceFilterer creates a new log filterer instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PillarBlockInterfaceFilterer, error) {
	contract, err := bindPillarBlockInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterfaceFilterer{contract: contract}, nil
}

// bindPillarBlockInterface binds a generic wrapper to an already deployed contract.
func bindPillarBlockInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PillarBlockInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PillarBlockInterface *PillarBlockInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PillarBlockInterface.Contract.PillarBlockInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PillarBlockInterface *PillarBlockInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.PillarBlockInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PillarBlockInterface *PillarBlockInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.PillarBlockInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PillarBlockInterface *PillarBlockInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PillarBlockInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PillarBlockInterface *PillarBlockInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PillarBlockInterface *PillarBlockInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.contract.Transact(opts, method, params...)
}

// FromBytes is a free data retrieval call binding the contract method 0x215b3e32.
//
// Solidity: function fromBytes(bytes b) pure returns(((uint256,bytes32,bytes32,bytes32),(address,int96)[]))
func (_PillarBlockInterface *PillarBlockInterfaceCaller) FromBytes(opts *bind.CallOpts, b []byte) (PillarBlockWithChanges, error) {
	var out []interface{}
	err := _PillarBlockInterface.contract.Call(opts, &out, "fromBytes", b)

	if err != nil {
		return *new(PillarBlockWithChanges), err
	}

	out0 := *abi.ConvertType(out[0], new(PillarBlockWithChanges)).(*PillarBlockWithChanges)

	return out0, err

}

// FromBytes is a free data retrieval call binding the contract method 0x215b3e32.
//
// Solidity: function fromBytes(bytes b) pure returns(((uint256,bytes32,bytes32,bytes32),(address,int96)[]))
func (_PillarBlockInterface *PillarBlockInterfaceSession) FromBytes(b []byte) (PillarBlockWithChanges, error) {
	return _PillarBlockInterface.Contract.FromBytes(&_PillarBlockInterface.CallOpts, b)
}

// FromBytes is a free data retrieval call binding the contract method 0x215b3e32.
//
// Solidity: function fromBytes(bytes b) pure returns(((uint256,bytes32,bytes32,bytes32),(address,int96)[]))
func (_PillarBlockInterface *PillarBlockInterfaceCallerSession) FromBytes(b []byte) (PillarBlockWithChanges, error) {
	return _PillarBlockInterface.Contract.FromBytes(&_PillarBlockInterface.CallOpts, b)
}

// GetHash is a free data retrieval call binding the contract method 0x7e885040.
//
// Solidity: function getHash((uint256,bytes32) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCaller) GetHash(opts *bind.CallOpts, b PillarBlockVote) ([32]byte, error) {
	var out []interface{}
	err := _PillarBlockInterface.contract.Call(opts, &out, "getHash", b)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x7e885040.
//
// Solidity: function getHash((uint256,bytes32) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceSession) GetHash(b PillarBlockVote) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash(&_PillarBlockInterface.CallOpts, b)
}

// GetHash is a free data retrieval call binding the contract method 0x7e885040.
//
// Solidity: function getHash((uint256,bytes32) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCallerSession) GetHash(b PillarBlockVote) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash(&_PillarBlockInterface.CallOpts, b)
}

// GetHash0 is a free data retrieval call binding the contract method 0xa621d8f9.
//
// Solidity: function getHash(((uint256,bytes32,bytes32,bytes32),(address,int96)[]) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCaller) GetHash0(opts *bind.CallOpts, b PillarBlockWithChanges) ([32]byte, error) {
	var out []interface{}
	err := _PillarBlockInterface.contract.Call(opts, &out, "getHash0", b)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash0 is a free data retrieval call binding the contract method 0xa621d8f9.
//
// Solidity: function getHash(((uint256,bytes32,bytes32,bytes32),(address,int96)[]) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceSession) GetHash0(b PillarBlockWithChanges) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash0(&_PillarBlockInterface.CallOpts, b)
}

// GetHash0 is a free data retrieval call binding the contract method 0xa621d8f9.
//
// Solidity: function getHash(((uint256,bytes32,bytes32,bytes32),(address,int96)[]) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCallerSession) GetHash0(b PillarBlockWithChanges) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash0(&_PillarBlockInterface.CallOpts, b)
}

// GetHash1 is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCaller) GetHash1(opts *bind.CallOpts, b []byte) ([32]byte, error) {
	var out []interface{}
	err := _PillarBlockInterface.contract.Call(opts, &out, "getHash1", b)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash1 is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceSession) GetHash1(b []byte) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash1(&_PillarBlockInterface.CallOpts, b)
}

// GetHash1 is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCallerSession) GetHash1(b []byte) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash1(&_PillarBlockInterface.CallOpts, b)
}

// GetHash2 is a free data retrieval call binding the contract method 0x1e26837b.
//
// Solidity: function getHash(((uint256,bytes32),(bytes32,bytes32)) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCaller) GetHash2(opts *bind.CallOpts, b PillarBlockSignedVote) ([32]byte, error) {
	var out []interface{}
	err := _PillarBlockInterface.contract.Call(opts, &out, "getHash2", b)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash2 is a free data retrieval call binding the contract method 0x1e26837b.
//
// Solidity: function getHash(((uint256,bytes32),(bytes32,bytes32)) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceSession) GetHash2(b PillarBlockSignedVote) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash2(&_PillarBlockInterface.CallOpts, b)
}

// GetHash2 is a free data retrieval call binding the contract method 0x1e26837b.
//
// Solidity: function getHash(((uint256,bytes32),(bytes32,bytes32)) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCallerSession) GetHash2(b PillarBlockSignedVote) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetHash2(&_PillarBlockInterface.CallOpts, b)
}

// GetVoteHash is a free data retrieval call binding the contract method 0x7734cd4a.
//
// Solidity: function getVoteHash(uint256 period, bytes32 bh) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCaller) GetVoteHash(opts *bind.CallOpts, period *big.Int, bh [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PillarBlockInterface.contract.Call(opts, &out, "getVoteHash", period, bh)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVoteHash is a free data retrieval call binding the contract method 0x7734cd4a.
//
// Solidity: function getVoteHash(uint256 period, bytes32 bh) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceSession) GetVoteHash(period *big.Int, bh [32]byte) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetVoteHash(&_PillarBlockInterface.CallOpts, period, bh)
}

// GetVoteHash is a free data retrieval call binding the contract method 0x7734cd4a.
//
// Solidity: function getVoteHash(uint256 period, bytes32 bh) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCallerSession) GetVoteHash(period *big.Int, bh [32]byte) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetVoteHash(&_PillarBlockInterface.CallOpts, period, bh)
}

// GetVoteHash0 is a free data retrieval call binding the contract method 0xfe2638c0.
//
// Solidity: function getVoteHash(((uint256,bytes32,bytes32,bytes32),(address,int96)[]) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCaller) GetVoteHash0(opts *bind.CallOpts, b PillarBlockWithChanges) ([32]byte, error) {
	var out []interface{}
	err := _PillarBlockInterface.contract.Call(opts, &out, "getVoteHash0", b)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVoteHash0 is a free data retrieval call binding the contract method 0xfe2638c0.
//
// Solidity: function getVoteHash(((uint256,bytes32,bytes32,bytes32),(address,int96)[]) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceSession) GetVoteHash0(b PillarBlockWithChanges) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetVoteHash0(&_PillarBlockInterface.CallOpts, b)
}

// GetVoteHash0 is a free data retrieval call binding the contract method 0xfe2638c0.
//
// Solidity: function getVoteHash(((uint256,bytes32,bytes32,bytes32),(address,int96)[]) b) pure returns(bytes32)
func (_PillarBlockInterface *PillarBlockInterfaceCallerSession) GetVoteHash0(b PillarBlockWithChanges) ([32]byte, error) {
	return _PillarBlockInterface.Contract.GetVoteHash0(&_PillarBlockInterface.CallOpts, b)
}
