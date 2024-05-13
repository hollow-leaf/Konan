// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VerifierAndTable

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

// AttestationProof is an auto generated low-level Go binding around an user-defined struct.
type AttestationProof struct {
	TupleRootNonce *big.Int
	Tuple          DataRootTuple
	Proof          BinaryMerkleProof
}

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// Namespace is an auto generated low-level Go binding around an user-defined struct.
type Namespace struct {
	Version [1]byte
	Id      [28]byte
}

// NamespaceMerkleMultiproof is an auto generated low-level Go binding around an user-defined struct.
type NamespaceMerkleMultiproof struct {
	BeginKey  *big.Int
	EndKey    *big.Int
	SideNodes []NamespaceNode
}

// NamespaceNode is an auto generated low-level Go binding around an user-defined struct.
type NamespaceNode struct {
	Min    Namespace
	Max    Namespace
	Digest [32]byte
}

// SharesProof is an auto generated low-level Go binding around an user-defined struct.
type SharesProof struct {
	Data             [][]byte
	ShareProofs      []NamespaceMerkleMultiproof
	Namespace        Namespace
	RowRoots         []NamespaceNode
	RowProofs        []BinaryMerkleProof
	AttestationProof AttestationProof
}

// VerifierAndTableMetaData contains all meta data concerning the VerifierAndTable contract.
var VerifierAndTableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes28\",\"name\":\"namespace\",\"type\":\"bytes28\"}],\"name\":\"IsNamespaceVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"sharesProof\",\"type\":\"tuple\"}],\"name\":\"VerifySharesToDataRootTupleRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumDAVerifier.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes28\",\"name\":\"\",\"type\":\"bytes28\"}],\"name\":\"VerifyTable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToBytes\",\"outputs\":[{\"internalType\":\"bytes28\",\"name\":\"\",\"type\":\"bytes28\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VerifierAndTableABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierAndTableMetaData.ABI instead.
var VerifierAndTableABI = VerifierAndTableMetaData.ABI

// VerifierAndTable is an auto generated Go binding around an Ethereum contract.
type VerifierAndTable struct {
	VerifierAndTableCaller     // Read-only binding to the contract
	VerifierAndTableTransactor // Write-only binding to the contract
	VerifierAndTableFilterer   // Log filterer for contract events
}

// VerifierAndTableCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierAndTableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierAndTableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierAndTableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierAndTableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierAndTableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierAndTableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierAndTableSession struct {
	Contract     *VerifierAndTable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierAndTableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierAndTableCallerSession struct {
	Contract *VerifierAndTableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VerifierAndTableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierAndTableTransactorSession struct {
	Contract     *VerifierAndTableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VerifierAndTableRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierAndTableRaw struct {
	Contract *VerifierAndTable // Generic contract binding to access the raw methods on
}

// VerifierAndTableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierAndTableCallerRaw struct {
	Contract *VerifierAndTableCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierAndTableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierAndTableTransactorRaw struct {
	Contract *VerifierAndTableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierAndTable creates a new instance of VerifierAndTable, bound to a specific deployed contract.
func NewVerifierAndTable(address common.Address, backend bind.ContractBackend) (*VerifierAndTable, error) {
	contract, err := bindVerifierAndTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierAndTable{VerifierAndTableCaller: VerifierAndTableCaller{contract: contract}, VerifierAndTableTransactor: VerifierAndTableTransactor{contract: contract}, VerifierAndTableFilterer: VerifierAndTableFilterer{contract: contract}}, nil
}

// NewVerifierAndTableCaller creates a new read-only instance of VerifierAndTable, bound to a specific deployed contract.
func NewVerifierAndTableCaller(address common.Address, caller bind.ContractCaller) (*VerifierAndTableCaller, error) {
	contract, err := bindVerifierAndTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierAndTableCaller{contract: contract}, nil
}

// NewVerifierAndTableTransactor creates a new write-only instance of VerifierAndTable, bound to a specific deployed contract.
func NewVerifierAndTableTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierAndTableTransactor, error) {
	contract, err := bindVerifierAndTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierAndTableTransactor{contract: contract}, nil
}

// NewVerifierAndTableFilterer creates a new log filterer instance of VerifierAndTable, bound to a specific deployed contract.
func NewVerifierAndTableFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierAndTableFilterer, error) {
	contract, err := bindVerifierAndTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierAndTableFilterer{contract: contract}, nil
}

// bindVerifierAndTable binds a generic wrapper to an already deployed contract.
func bindVerifierAndTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierAndTableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierAndTable *VerifierAndTableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierAndTable.Contract.VerifierAndTableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierAndTable *VerifierAndTableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierAndTable.Contract.VerifierAndTableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierAndTable *VerifierAndTableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierAndTable.Contract.VerifierAndTableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierAndTable *VerifierAndTableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierAndTable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierAndTable *VerifierAndTableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierAndTable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierAndTable *VerifierAndTableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierAndTable.Contract.contract.Transact(opts, method, params...)
}

// IsNamespaceVerified is a free data retrieval call binding the contract method 0x17d0b41b.
//
// Solidity: function IsNamespaceVerified(bytes28 namespace) view returns(bool)
func (_VerifierAndTable *VerifierAndTableCaller) IsNamespaceVerified(opts *bind.CallOpts, namespace [28]byte) (bool, error) {
	var out []interface{}
	err := _VerifierAndTable.contract.Call(opts, &out, "IsNamespaceVerified", namespace)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNamespaceVerified is a free data retrieval call binding the contract method 0x17d0b41b.
//
// Solidity: function IsNamespaceVerified(bytes28 namespace) view returns(bool)
func (_VerifierAndTable *VerifierAndTableSession) IsNamespaceVerified(namespace [28]byte) (bool, error) {
	return _VerifierAndTable.Contract.IsNamespaceVerified(&_VerifierAndTable.CallOpts, namespace)
}

// IsNamespaceVerified is a free data retrieval call binding the contract method 0x17d0b41b.
//
// Solidity: function IsNamespaceVerified(bytes28 namespace) view returns(bool)
func (_VerifierAndTable *VerifierAndTableCallerSession) IsNamespaceVerified(namespace [28]byte) (bool, error) {
	return _VerifierAndTable.Contract.IsNamespaceVerified(&_VerifierAndTable.CallOpts, namespace)
}

// VerifyTable is a free data retrieval call binding the contract method 0x30c433f0.
//
// Solidity: function VerifyTable(bytes28 ) view returns(uint256 blockHeight, bool verified)
func (_VerifierAndTable *VerifierAndTableCaller) VerifyTable(opts *bind.CallOpts, arg0 [28]byte) (struct {
	BlockHeight *big.Int
	Verified    bool
}, error) {
	var out []interface{}
	err := _VerifierAndTable.contract.Call(opts, &out, "VerifyTable", arg0)

	outstruct := new(struct {
		BlockHeight *big.Int
		Verified    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Verified = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// VerifyTable is a free data retrieval call binding the contract method 0x30c433f0.
//
// Solidity: function VerifyTable(bytes28 ) view returns(uint256 blockHeight, bool verified)
func (_VerifierAndTable *VerifierAndTableSession) VerifyTable(arg0 [28]byte) (struct {
	BlockHeight *big.Int
	Verified    bool
}, error) {
	return _VerifierAndTable.Contract.VerifyTable(&_VerifierAndTable.CallOpts, arg0)
}

// VerifyTable is a free data retrieval call binding the contract method 0x30c433f0.
//
// Solidity: function VerifyTable(bytes28 ) view returns(uint256 blockHeight, bool verified)
func (_VerifierAndTable *VerifierAndTableCallerSession) VerifyTable(arg0 [28]byte) (struct {
	BlockHeight *big.Int
	Verified    bool
}, error) {
	return _VerifierAndTable.Contract.VerifyTable(&_VerifierAndTable.CallOpts, arg0)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 data) pure returns(bytes28)
func (_VerifierAndTable *VerifierAndTableCaller) Bytes32ToBytes(opts *bind.CallOpts, data [32]byte) ([28]byte, error) {
	var out []interface{}
	err := _VerifierAndTable.contract.Call(opts, &out, "bytes32ToBytes", data)

	if err != nil {
		return *new([28]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([28]byte)).(*[28]byte)

	return out0, err

}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 data) pure returns(bytes28)
func (_VerifierAndTable *VerifierAndTableSession) Bytes32ToBytes(data [32]byte) ([28]byte, error) {
	return _VerifierAndTable.Contract.Bytes32ToBytes(&_VerifierAndTable.CallOpts, data)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 data) pure returns(bytes28)
func (_VerifierAndTable *VerifierAndTableCallerSession) Bytes32ToBytes(data [32]byte) ([28]byte, error) {
	return _VerifierAndTable.Contract.Bytes32ToBytes(&_VerifierAndTable.CallOpts, data)
}

// VerifySharesToDataRootTupleRoot is a paid mutator transaction binding the contract method 0xb5e3cf5b.
//
// Solidity: function VerifySharesToDataRootTupleRoot((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) sharesProof) returns(bool, uint8)
func (_VerifierAndTable *VerifierAndTableTransactor) VerifySharesToDataRootTupleRoot(opts *bind.TransactOpts, sharesProof interface{}) (*types.Transaction, error) {
	return _VerifierAndTable.contract.Transact(opts, "VerifySharesToDataRootTupleRoot", sharesProof)
}

// VerifySharesToDataRootTupleRoot is a paid mutator transaction binding the contract method 0xb5e3cf5b.
//
// Solidity: function VerifySharesToDataRootTupleRoot((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) sharesProof) returns(bool, uint8)
func (_VerifierAndTable *VerifierAndTableSession) VerifySharesToDataRootTupleRoot(sharesProof SharesProof) (*types.Transaction, error) {
	return _VerifierAndTable.Contract.VerifySharesToDataRootTupleRoot(&_VerifierAndTable.TransactOpts, sharesProof)
}

// VerifySharesToDataRootTupleRoot is a paid mutator transaction binding the contract method 0xb5e3cf5b.
//
// Solidity: function VerifySharesToDataRootTupleRoot((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) sharesProof) returns(bool, uint8)
func (_VerifierAndTable *VerifierAndTableTransactorSession) VerifySharesToDataRootTupleRoot(sharesProof SharesProof) (*types.Transaction, error) {
	return _VerifierAndTable.Contract.VerifySharesToDataRootTupleRoot(&_VerifierAndTable.TransactOpts, sharesProof)
}
