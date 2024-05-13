// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package drawableNft721

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

// ERC721ATokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type ERC721ATokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
}

// DrawableNft721MetaData contains all meta data concerning the DrawableNft721 contract.
var DrawableNft721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subscriptionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBatchSize_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"drawtypes_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"collectionSize_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountForDevs_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"OnlyOwnerOrCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"}],\"name\":\"CoordinatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"drawId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"probability\",\"type\":\"uint256[]\"}],\"name\":\"DrawPropertySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"drawId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"drawIndex\",\"type\":\"uint256\"}],\"name\":\"Drawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"drawId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"probability\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"SpecialPropertySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountForDevs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"drawId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"probability\",\"type\":\"uint256[]\"}],\"name\":\"createSpecialEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"devMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_drawId\",\"type\":\"uint256\"}],\"name\":\"drawByVRF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drawProbability\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drawStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drawTypes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwnershipData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structERC721A.TokenOwnership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"publicPriceWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicSaleStartTime\",\"type\":\"uint256\"}],\"name\":\"isPublicSaleOn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPerAddressDuringMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOwnerToExplicitlySet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"numberMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"publicSaleMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestToDrawId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"drawId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_vrfCoordinator\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorV2Plus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"publicSaleStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicPriceWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"drawId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"probability\",\"type\":\"uint256[]\"}],\"name\":\"setDrawProperty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"setOwnersExplicit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"publicPriceWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicSaleStartTime\",\"type\":\"uint256\"}],\"name\":\"setupPublicSaleInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"specialProbability\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTypes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrfCoordinator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DrawableNft721ABI is the input ABI used to generate the binding from.
// Deprecated: Use DrawableNft721MetaData.ABI instead.
var DrawableNft721ABI = DrawableNft721MetaData.ABI

// DrawableNft721 is an auto generated Go binding around an Ethereum contract.
type DrawableNft721 struct {
	DrawableNft721Caller     // Read-only binding to the contract
	DrawableNft721Transactor // Write-only binding to the contract
	DrawableNft721Filterer   // Log filterer for contract events
}

// DrawableNft721Caller is an auto generated read-only Go binding around an Ethereum contract.
type DrawableNft721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrawableNft721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DrawableNft721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrawableNft721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DrawableNft721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrawableNft721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DrawableNft721Session struct {
	Contract     *DrawableNft721   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DrawableNft721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DrawableNft721CallerSession struct {
	Contract *DrawableNft721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DrawableNft721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DrawableNft721TransactorSession struct {
	Contract     *DrawableNft721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DrawableNft721Raw is an auto generated low-level Go binding around an Ethereum contract.
type DrawableNft721Raw struct {
	Contract *DrawableNft721 // Generic contract binding to access the raw methods on
}

// DrawableNft721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DrawableNft721CallerRaw struct {
	Contract *DrawableNft721Caller // Generic read-only contract binding to access the raw methods on
}

// DrawableNft721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DrawableNft721TransactorRaw struct {
	Contract *DrawableNft721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDrawableNft721 creates a new instance of DrawableNft721, bound to a specific deployed contract.
func NewDrawableNft721(address common.Address, backend bind.ContractBackend) (*DrawableNft721, error) {
	contract, err := bindDrawableNft721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721{DrawableNft721Caller: DrawableNft721Caller{contract: contract}, DrawableNft721Transactor: DrawableNft721Transactor{contract: contract}, DrawableNft721Filterer: DrawableNft721Filterer{contract: contract}}, nil
}

// NewDrawableNft721Caller creates a new read-only instance of DrawableNft721, bound to a specific deployed contract.
func NewDrawableNft721Caller(address common.Address, caller bind.ContractCaller) (*DrawableNft721Caller, error) {
	contract, err := bindDrawableNft721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721Caller{contract: contract}, nil
}

// NewDrawableNft721Transactor creates a new write-only instance of DrawableNft721, bound to a specific deployed contract.
func NewDrawableNft721Transactor(address common.Address, transactor bind.ContractTransactor) (*DrawableNft721Transactor, error) {
	contract, err := bindDrawableNft721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721Transactor{contract: contract}, nil
}

// NewDrawableNft721Filterer creates a new log filterer instance of DrawableNft721, bound to a specific deployed contract.
func NewDrawableNft721Filterer(address common.Address, filterer bind.ContractFilterer) (*DrawableNft721Filterer, error) {
	contract, err := bindDrawableNft721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721Filterer{contract: contract}, nil
}

// bindDrawableNft721 binds a generic wrapper to an already deployed contract.
func bindDrawableNft721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DrawableNft721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrawableNft721 *DrawableNft721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DrawableNft721.Contract.DrawableNft721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrawableNft721 *DrawableNft721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrawableNft721.Contract.DrawableNft721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrawableNft721 *DrawableNft721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DrawableNft721.Contract.DrawableNft721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrawableNft721 *DrawableNft721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DrawableNft721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrawableNft721 *DrawableNft721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrawableNft721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrawableNft721 *DrawableNft721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DrawableNft721.Contract.contract.Transact(opts, method, params...)
}

// AmountForDevs is a free data retrieval call binding the contract method 0xfbe1aa51.
//
// Solidity: function amountForDevs() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) AmountForDevs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "amountForDevs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountForDevs is a free data retrieval call binding the contract method 0xfbe1aa51.
//
// Solidity: function amountForDevs() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) AmountForDevs() (*big.Int, error) {
	return _DrawableNft721.Contract.AmountForDevs(&_DrawableNft721.CallOpts)
}

// AmountForDevs is a free data retrieval call binding the contract method 0xfbe1aa51.
//
// Solidity: function amountForDevs() view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) AmountForDevs() (*big.Int, error) {
	return _DrawableNft721.Contract.AmountForDevs(&_DrawableNft721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DrawableNft721.Contract.BalanceOf(&_DrawableNft721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DrawableNft721.Contract.BalanceOf(&_DrawableNft721.CallOpts, owner)
}

// DrawProbability is a free data retrieval call binding the contract method 0x820318a6.
//
// Solidity: function drawProbability(uint256 , uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) DrawProbability(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "drawProbability", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DrawProbability is a free data retrieval call binding the contract method 0x820318a6.
//
// Solidity: function drawProbability(uint256 , uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) DrawProbability(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.DrawProbability(&_DrawableNft721.CallOpts, arg0, arg1)
}

// DrawProbability is a free data retrieval call binding the contract method 0x820318a6.
//
// Solidity: function drawProbability(uint256 , uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) DrawProbability(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.DrawProbability(&_DrawableNft721.CallOpts, arg0, arg1)
}

// DrawStatus is a free data retrieval call binding the contract method 0xaa869589.
//
// Solidity: function drawStatus(uint256 , uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) DrawStatus(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "drawStatus", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DrawStatus is a free data retrieval call binding the contract method 0xaa869589.
//
// Solidity: function drawStatus(uint256 , uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) DrawStatus(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.DrawStatus(&_DrawableNft721.CallOpts, arg0, arg1)
}

// DrawStatus is a free data retrieval call binding the contract method 0xaa869589.
//
// Solidity: function drawStatus(uint256 , uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) DrawStatus(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.DrawStatus(&_DrawableNft721.CallOpts, arg0, arg1)
}

// DrawTypes is a free data retrieval call binding the contract method 0x4e34ecfe.
//
// Solidity: function drawTypes(uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) DrawTypes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "drawTypes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DrawTypes is a free data retrieval call binding the contract method 0x4e34ecfe.
//
// Solidity: function drawTypes(uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) DrawTypes(arg0 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.DrawTypes(&_DrawableNft721.CallOpts, arg0)
}

// DrawTypes is a free data retrieval call binding the contract method 0x4e34ecfe.
//
// Solidity: function drawTypes(uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) DrawTypes(arg0 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.DrawTypes(&_DrawableNft721.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DrawableNft721 *DrawableNft721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DrawableNft721 *DrawableNft721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DrawableNft721.Contract.GetApproved(&_DrawableNft721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DrawableNft721 *DrawableNft721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DrawableNft721.Contract.GetApproved(&_DrawableNft721.CallOpts, tokenId)
}

// GetOwnershipData is a free data retrieval call binding the contract method 0x9231ab2a.
//
// Solidity: function getOwnershipData(uint256 tokenId) view returns((address,uint64))
func (_DrawableNft721 *DrawableNft721Caller) GetOwnershipData(opts *bind.CallOpts, tokenId *big.Int) (ERC721ATokenOwnership, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "getOwnershipData", tokenId)

	if err != nil {
		return *new(ERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC721ATokenOwnership)).(*ERC721ATokenOwnership)

	return out0, err

}

// GetOwnershipData is a free data retrieval call binding the contract method 0x9231ab2a.
//
// Solidity: function getOwnershipData(uint256 tokenId) view returns((address,uint64))
func (_DrawableNft721 *DrawableNft721Session) GetOwnershipData(tokenId *big.Int) (ERC721ATokenOwnership, error) {
	return _DrawableNft721.Contract.GetOwnershipData(&_DrawableNft721.CallOpts, tokenId)
}

// GetOwnershipData is a free data retrieval call binding the contract method 0x9231ab2a.
//
// Solidity: function getOwnershipData(uint256 tokenId) view returns((address,uint64))
func (_DrawableNft721 *DrawableNft721CallerSession) GetOwnershipData(tokenId *big.Int) (ERC721ATokenOwnership, error) {
	return _DrawableNft721.Contract.GetOwnershipData(&_DrawableNft721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DrawableNft721 *DrawableNft721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DrawableNft721 *DrawableNft721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DrawableNft721.Contract.IsApprovedForAll(&_DrawableNft721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DrawableNft721 *DrawableNft721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DrawableNft721.Contract.IsApprovedForAll(&_DrawableNft721.CallOpts, owner, operator)
}

// IsPublicSaleOn is a free data retrieval call binding the contract method 0xdbcad76f.
//
// Solidity: function isPublicSaleOn(uint256 publicPriceWei, uint256 publicSaleStartTime) view returns(bool)
func (_DrawableNft721 *DrawableNft721Caller) IsPublicSaleOn(opts *bind.CallOpts, publicPriceWei *big.Int, publicSaleStartTime *big.Int) (bool, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "isPublicSaleOn", publicPriceWei, publicSaleStartTime)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSaleOn is a free data retrieval call binding the contract method 0xdbcad76f.
//
// Solidity: function isPublicSaleOn(uint256 publicPriceWei, uint256 publicSaleStartTime) view returns(bool)
func (_DrawableNft721 *DrawableNft721Session) IsPublicSaleOn(publicPriceWei *big.Int, publicSaleStartTime *big.Int) (bool, error) {
	return _DrawableNft721.Contract.IsPublicSaleOn(&_DrawableNft721.CallOpts, publicPriceWei, publicSaleStartTime)
}

// IsPublicSaleOn is a free data retrieval call binding the contract method 0xdbcad76f.
//
// Solidity: function isPublicSaleOn(uint256 publicPriceWei, uint256 publicSaleStartTime) view returns(bool)
func (_DrawableNft721 *DrawableNft721CallerSession) IsPublicSaleOn(publicPriceWei *big.Int, publicSaleStartTime *big.Int) (bool, error) {
	return _DrawableNft721.Contract.IsPublicSaleOn(&_DrawableNft721.CallOpts, publicPriceWei, publicSaleStartTime)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_DrawableNft721 *DrawableNft721Caller) KeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "keyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_DrawableNft721 *DrawableNft721Session) KeyHash() ([32]byte, error) {
	return _DrawableNft721.Contract.KeyHash(&_DrawableNft721.CallOpts)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_DrawableNft721 *DrawableNft721CallerSession) KeyHash() ([32]byte, error) {
	return _DrawableNft721.Contract.KeyHash(&_DrawableNft721.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) LastRequestId() (*big.Int, error) {
	return _DrawableNft721.Contract.LastRequestId(&_DrawableNft721.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) LastRequestId() (*big.Int, error) {
	return _DrawableNft721.Contract.LastRequestId(&_DrawableNft721.CallOpts)
}

// MaxPerAddressDuringMint is a free data retrieval call binding the contract method 0x8bc35c2f.
//
// Solidity: function maxPerAddressDuringMint() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) MaxPerAddressDuringMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "maxPerAddressDuringMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerAddressDuringMint is a free data retrieval call binding the contract method 0x8bc35c2f.
//
// Solidity: function maxPerAddressDuringMint() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) MaxPerAddressDuringMint() (*big.Int, error) {
	return _DrawableNft721.Contract.MaxPerAddressDuringMint(&_DrawableNft721.CallOpts)
}

// MaxPerAddressDuringMint is a free data retrieval call binding the contract method 0x8bc35c2f.
//
// Solidity: function maxPerAddressDuringMint() view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) MaxPerAddressDuringMint() (*big.Int, error) {
	return _DrawableNft721.Contract.MaxPerAddressDuringMint(&_DrawableNft721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DrawableNft721 *DrawableNft721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DrawableNft721 *DrawableNft721Session) Name() (string, error) {
	return _DrawableNft721.Contract.Name(&_DrawableNft721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DrawableNft721 *DrawableNft721CallerSession) Name() (string, error) {
	return _DrawableNft721.Contract.Name(&_DrawableNft721.CallOpts)
}

// NextOwnerToExplicitlySet is a free data retrieval call binding the contract method 0xd7224ba0.
//
// Solidity: function nextOwnerToExplicitlySet() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) NextOwnerToExplicitlySet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "nextOwnerToExplicitlySet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOwnerToExplicitlySet is a free data retrieval call binding the contract method 0xd7224ba0.
//
// Solidity: function nextOwnerToExplicitlySet() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) NextOwnerToExplicitlySet() (*big.Int, error) {
	return _DrawableNft721.Contract.NextOwnerToExplicitlySet(&_DrawableNft721.CallOpts)
}

// NextOwnerToExplicitlySet is a free data retrieval call binding the contract method 0xd7224ba0.
//
// Solidity: function nextOwnerToExplicitlySet() view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) NextOwnerToExplicitlySet() (*big.Int, error) {
	return _DrawableNft721.Contract.NextOwnerToExplicitlySet(&_DrawableNft721.CallOpts)
}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address owner) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) NumberMinted(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "numberMinted", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address owner) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) NumberMinted(owner common.Address) (*big.Int, error) {
	return _DrawableNft721.Contract.NumberMinted(&_DrawableNft721.CallOpts, owner)
}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address owner) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) NumberMinted(owner common.Address) (*big.Int, error) {
	return _DrawableNft721.Contract.NumberMinted(&_DrawableNft721.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DrawableNft721 *DrawableNft721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DrawableNft721 *DrawableNft721Session) Owner() (common.Address, error) {
	return _DrawableNft721.Contract.Owner(&_DrawableNft721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DrawableNft721 *DrawableNft721CallerSession) Owner() (common.Address, error) {
	return _DrawableNft721.Contract.Owner(&_DrawableNft721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DrawableNft721 *DrawableNft721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DrawableNft721 *DrawableNft721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DrawableNft721.Contract.OwnerOf(&_DrawableNft721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DrawableNft721 *DrawableNft721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DrawableNft721.Contract.OwnerOf(&_DrawableNft721.CallOpts, tokenId)
}

// RequestToDrawId is a free data retrieval call binding the contract method 0xe2b24f56.
//
// Solidity: function requestToDrawId(uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) RequestToDrawId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "requestToDrawId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestToDrawId is a free data retrieval call binding the contract method 0xe2b24f56.
//
// Solidity: function requestToDrawId(uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) RequestToDrawId(arg0 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.RequestToDrawId(&_DrawableNft721.CallOpts, arg0)
}

// RequestToDrawId is a free data retrieval call binding the contract method 0xe2b24f56.
//
// Solidity: function requestToDrawId(uint256 ) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) RequestToDrawId(arg0 *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.RequestToDrawId(&_DrawableNft721.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool exists, bool fulfilled, uint256 drawId, uint256 rank)
func (_DrawableNft721 *DrawableNft721Caller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Exists    bool
	Fulfilled bool
	DrawId    *big.Int
	Rank      *big.Int
}, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "s_requests", arg0)

	outstruct := new(struct {
		Exists    bool
		Fulfilled bool
		DrawId    *big.Int
		Rank      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.DrawId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rank = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool exists, bool fulfilled, uint256 drawId, uint256 rank)
func (_DrawableNft721 *DrawableNft721Session) SRequests(arg0 *big.Int) (struct {
	Exists    bool
	Fulfilled bool
	DrawId    *big.Int
	Rank      *big.Int
}, error) {
	return _DrawableNft721.Contract.SRequests(&_DrawableNft721.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool exists, bool fulfilled, uint256 drawId, uint256 rank)
func (_DrawableNft721 *DrawableNft721CallerSession) SRequests(arg0 *big.Int) (struct {
	Exists    bool
	Fulfilled bool
	DrawId    *big.Int
	Rank      *big.Int
}, error) {
	return _DrawableNft721.Contract.SRequests(&_DrawableNft721.CallOpts, arg0)
}

// SVrfCoordinator is a free data retrieval call binding the contract method 0x9eccacf6.
//
// Solidity: function s_vrfCoordinator() view returns(address)
func (_DrawableNft721 *DrawableNft721Caller) SVrfCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "s_vrfCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SVrfCoordinator is a free data retrieval call binding the contract method 0x9eccacf6.
//
// Solidity: function s_vrfCoordinator() view returns(address)
func (_DrawableNft721 *DrawableNft721Session) SVrfCoordinator() (common.Address, error) {
	return _DrawableNft721.Contract.SVrfCoordinator(&_DrawableNft721.CallOpts)
}

// SVrfCoordinator is a free data retrieval call binding the contract method 0x9eccacf6.
//
// Solidity: function s_vrfCoordinator() view returns(address)
func (_DrawableNft721 *DrawableNft721CallerSession) SVrfCoordinator() (common.Address, error) {
	return _DrawableNft721.Contract.SVrfCoordinator(&_DrawableNft721.CallOpts)
}

// SaleConfig is a free data retrieval call binding the contract method 0x90aa0b0f.
//
// Solidity: function saleConfig() view returns(uint256 publicSaleStartTime, uint256 publicPriceWei)
func (_DrawableNft721 *DrawableNft721Caller) SaleConfig(opts *bind.CallOpts) (struct {
	PublicSaleStartTime *big.Int
	PublicPriceWei      *big.Int
}, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "saleConfig")

	outstruct := new(struct {
		PublicSaleStartTime *big.Int
		PublicPriceWei      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicSaleStartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PublicPriceWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SaleConfig is a free data retrieval call binding the contract method 0x90aa0b0f.
//
// Solidity: function saleConfig() view returns(uint256 publicSaleStartTime, uint256 publicPriceWei)
func (_DrawableNft721 *DrawableNft721Session) SaleConfig() (struct {
	PublicSaleStartTime *big.Int
	PublicPriceWei      *big.Int
}, error) {
	return _DrawableNft721.Contract.SaleConfig(&_DrawableNft721.CallOpts)
}

// SaleConfig is a free data retrieval call binding the contract method 0x90aa0b0f.
//
// Solidity: function saleConfig() view returns(uint256 publicSaleStartTime, uint256 publicPriceWei)
func (_DrawableNft721 *DrawableNft721CallerSession) SaleConfig() (struct {
	PublicSaleStartTime *big.Int
	PublicPriceWei      *big.Int
}, error) {
	return _DrawableNft721.Contract.SaleConfig(&_DrawableNft721.CallOpts)
}

// SpecialProbability is a free data retrieval call binding the contract method 0xf9586c62.
//
// Solidity: function specialProbability(uint256 ) view returns(bool exists, uint256 startTimestamp, uint256 endTimestamp)
func (_DrawableNft721 *DrawableNft721Caller) SpecialProbability(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Exists         bool
	StartTimestamp *big.Int
	EndTimestamp   *big.Int
}, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "specialProbability", arg0)

	outstruct := new(struct {
		Exists         bool
		StartTimestamp *big.Int
		EndTimestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.StartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SpecialProbability is a free data retrieval call binding the contract method 0xf9586c62.
//
// Solidity: function specialProbability(uint256 ) view returns(bool exists, uint256 startTimestamp, uint256 endTimestamp)
func (_DrawableNft721 *DrawableNft721Session) SpecialProbability(arg0 *big.Int) (struct {
	Exists         bool
	StartTimestamp *big.Int
	EndTimestamp   *big.Int
}, error) {
	return _DrawableNft721.Contract.SpecialProbability(&_DrawableNft721.CallOpts, arg0)
}

// SpecialProbability is a free data retrieval call binding the contract method 0xf9586c62.
//
// Solidity: function specialProbability(uint256 ) view returns(bool exists, uint256 startTimestamp, uint256 endTimestamp)
func (_DrawableNft721 *DrawableNft721CallerSession) SpecialProbability(arg0 *big.Int) (struct {
	Exists         bool
	StartTimestamp *big.Int
	EndTimestamp   *big.Int
}, error) {
	return _DrawableNft721.Contract.SpecialProbability(&_DrawableNft721.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DrawableNft721 *DrawableNft721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DrawableNft721 *DrawableNft721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DrawableNft721.Contract.SupportsInterface(&_DrawableNft721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DrawableNft721 *DrawableNft721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DrawableNft721.Contract.SupportsInterface(&_DrawableNft721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DrawableNft721 *DrawableNft721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DrawableNft721 *DrawableNft721Session) Symbol() (string, error) {
	return _DrawableNft721.Contract.Symbol(&_DrawableNft721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DrawableNft721 *DrawableNft721CallerSession) Symbol() (string, error) {
	return _DrawableNft721.Contract.Symbol(&_DrawableNft721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.TokenByIndex(&_DrawableNft721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.TokenByIndex(&_DrawableNft721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.TokenOfOwnerByIndex(&_DrawableNft721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _DrawableNft721.Contract.TokenOfOwnerByIndex(&_DrawableNft721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DrawableNft721 *DrawableNft721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DrawableNft721 *DrawableNft721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _DrawableNft721.Contract.TokenURI(&_DrawableNft721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DrawableNft721 *DrawableNft721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DrawableNft721.Contract.TokenURI(&_DrawableNft721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) TotalSupply() (*big.Int, error) {
	return _DrawableNft721.Contract.TotalSupply(&_DrawableNft721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) TotalSupply() (*big.Int, error) {
	return _DrawableNft721.Contract.TotalSupply(&_DrawableNft721.CallOpts)
}

// TotalTypes is a free data retrieval call binding the contract method 0xf15df2e5.
//
// Solidity: function totalTypes() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Caller) TotalTypes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "totalTypes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTypes is a free data retrieval call binding the contract method 0xf15df2e5.
//
// Solidity: function totalTypes() view returns(uint256)
func (_DrawableNft721 *DrawableNft721Session) TotalTypes() (*big.Int, error) {
	return _DrawableNft721.Contract.TotalTypes(&_DrawableNft721.CallOpts)
}

// TotalTypes is a free data retrieval call binding the contract method 0xf15df2e5.
//
// Solidity: function totalTypes() view returns(uint256)
func (_DrawableNft721 *DrawableNft721CallerSession) TotalTypes() (*big.Int, error) {
	return _DrawableNft721.Contract.TotalTypes(&_DrawableNft721.CallOpts)
}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_DrawableNft721 *DrawableNft721Caller) VrfCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DrawableNft721.contract.Call(opts, &out, "vrfCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_DrawableNft721 *DrawableNft721Session) VrfCoordinator() (common.Address, error) {
	return _DrawableNft721.Contract.VrfCoordinator(&_DrawableNft721.CallOpts)
}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_DrawableNft721 *DrawableNft721CallerSession) VrfCoordinator() (common.Address, error) {
	return _DrawableNft721.Contract.VrfCoordinator(&_DrawableNft721.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DrawableNft721 *DrawableNft721Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DrawableNft721 *DrawableNft721Session) AcceptOwnership() (*types.Transaction, error) {
	return _DrawableNft721.Contract.AcceptOwnership(&_DrawableNft721.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DrawableNft721.Contract.AcceptOwnership(&_DrawableNft721.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.Approve(&_DrawableNft721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.Approve(&_DrawableNft721.TransactOpts, to, tokenId)
}

// CreateSpecialEvent is a paid mutator transaction binding the contract method 0x9f1070e4.
//
// Solidity: function createSpecialEvent(uint256 startTimestamp, uint256 endTimestamp, uint256 drawId, uint256[] probability) returns()
func (_DrawableNft721 *DrawableNft721Transactor) CreateSpecialEvent(opts *bind.TransactOpts, startTimestamp *big.Int, endTimestamp *big.Int, drawId *big.Int, probability []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "createSpecialEvent", startTimestamp, endTimestamp, drawId, probability)
}

// CreateSpecialEvent is a paid mutator transaction binding the contract method 0x9f1070e4.
//
// Solidity: function createSpecialEvent(uint256 startTimestamp, uint256 endTimestamp, uint256 drawId, uint256[] probability) returns()
func (_DrawableNft721 *DrawableNft721Session) CreateSpecialEvent(startTimestamp *big.Int, endTimestamp *big.Int, drawId *big.Int, probability []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.CreateSpecialEvent(&_DrawableNft721.TransactOpts, startTimestamp, endTimestamp, drawId, probability)
}

// CreateSpecialEvent is a paid mutator transaction binding the contract method 0x9f1070e4.
//
// Solidity: function createSpecialEvent(uint256 startTimestamp, uint256 endTimestamp, uint256 drawId, uint256[] probability) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) CreateSpecialEvent(startTimestamp *big.Int, endTimestamp *big.Int, drawId *big.Int, probability []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.CreateSpecialEvent(&_DrawableNft721.TransactOpts, startTimestamp, endTimestamp, drawId, probability)
}

// DevMint is a paid mutator transaction binding the contract method 0x375a069a.
//
// Solidity: function devMint(uint256 quantity) returns()
func (_DrawableNft721 *DrawableNft721Transactor) DevMint(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "devMint", quantity)
}

// DevMint is a paid mutator transaction binding the contract method 0x375a069a.
//
// Solidity: function devMint(uint256 quantity) returns()
func (_DrawableNft721 *DrawableNft721Session) DevMint(quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.DevMint(&_DrawableNft721.TransactOpts, quantity)
}

// DevMint is a paid mutator transaction binding the contract method 0x375a069a.
//
// Solidity: function devMint(uint256 quantity) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) DevMint(quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.DevMint(&_DrawableNft721.TransactOpts, quantity)
}

// DrawByVRF is a paid mutator transaction binding the contract method 0x48a36147.
//
// Solidity: function drawByVRF(uint256 _tokenId, uint256 _drawId) returns()
func (_DrawableNft721 *DrawableNft721Transactor) DrawByVRF(opts *bind.TransactOpts, _tokenId *big.Int, _drawId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "drawByVRF", _tokenId, _drawId)
}

// DrawByVRF is a paid mutator transaction binding the contract method 0x48a36147.
//
// Solidity: function drawByVRF(uint256 _tokenId, uint256 _drawId) returns()
func (_DrawableNft721 *DrawableNft721Session) DrawByVRF(_tokenId *big.Int, _drawId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.DrawByVRF(&_DrawableNft721.TransactOpts, _tokenId, _drawId)
}

// DrawByVRF is a paid mutator transaction binding the contract method 0x48a36147.
//
// Solidity: function drawByVRF(uint256 _tokenId, uint256 _drawId) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) DrawByVRF(_tokenId *big.Int, _drawId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.DrawByVRF(&_DrawableNft721.TransactOpts, _tokenId, _drawId)
}

// PublicSaleMint is a paid mutator transaction binding the contract method 0xb3ab66b0.
//
// Solidity: function publicSaleMint(uint256 quantity) payable returns()
func (_DrawableNft721 *DrawableNft721Transactor) PublicSaleMint(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "publicSaleMint", quantity)
}

// PublicSaleMint is a paid mutator transaction binding the contract method 0xb3ab66b0.
//
// Solidity: function publicSaleMint(uint256 quantity) payable returns()
func (_DrawableNft721 *DrawableNft721Session) PublicSaleMint(quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.PublicSaleMint(&_DrawableNft721.TransactOpts, quantity)
}

// PublicSaleMint is a paid mutator transaction binding the contract method 0xb3ab66b0.
//
// Solidity: function publicSaleMint(uint256 quantity) payable returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) PublicSaleMint(quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.PublicSaleMint(&_DrawableNft721.TransactOpts, quantity)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_DrawableNft721 *DrawableNft721Transactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_DrawableNft721 *DrawableNft721Session) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.RawFulfillRandomWords(&_DrawableNft721.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.RawFulfillRandomWords(&_DrawableNft721.TransactOpts, requestId, randomWords)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SafeTransferFrom(&_DrawableNft721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SafeTransferFrom(&_DrawableNft721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_DrawableNft721 *DrawableNft721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SafeTransferFrom0(&_DrawableNft721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SafeTransferFrom0(&_DrawableNft721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DrawableNft721 *DrawableNft721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetApprovalForAll(&_DrawableNft721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetApprovalForAll(&_DrawableNft721.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_DrawableNft721 *DrawableNft721Session) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetBaseURI(&_DrawableNft721.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetBaseURI(&_DrawableNft721.TransactOpts, baseURI)
}

// SetCoordinator is a paid mutator transaction binding the contract method 0x8ea98117.
//
// Solidity: function setCoordinator(address _vrfCoordinator) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "setCoordinator", _vrfCoordinator)
}

// SetCoordinator is a paid mutator transaction binding the contract method 0x8ea98117.
//
// Solidity: function setCoordinator(address _vrfCoordinator) returns()
func (_DrawableNft721 *DrawableNft721Session) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetCoordinator(&_DrawableNft721.TransactOpts, _vrfCoordinator)
}

// SetCoordinator is a paid mutator transaction binding the contract method 0x8ea98117.
//
// Solidity: function setCoordinator(address _vrfCoordinator) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetCoordinator(&_DrawableNft721.TransactOpts, _vrfCoordinator)
}

// SetDrawProperty is a paid mutator transaction binding the contract method 0x9b4f254c.
//
// Solidity: function setDrawProperty(uint256 drawId, uint256[] probability) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SetDrawProperty(opts *bind.TransactOpts, drawId *big.Int, probability []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "setDrawProperty", drawId, probability)
}

// SetDrawProperty is a paid mutator transaction binding the contract method 0x9b4f254c.
//
// Solidity: function setDrawProperty(uint256 drawId, uint256[] probability) returns()
func (_DrawableNft721 *DrawableNft721Session) SetDrawProperty(drawId *big.Int, probability []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetDrawProperty(&_DrawableNft721.TransactOpts, drawId, probability)
}

// SetDrawProperty is a paid mutator transaction binding the contract method 0x9b4f254c.
//
// Solidity: function setDrawProperty(uint256 drawId, uint256[] probability) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SetDrawProperty(drawId *big.Int, probability []*big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetDrawProperty(&_DrawableNft721.TransactOpts, drawId, probability)
}

// SetOwnersExplicit is a paid mutator transaction binding the contract method 0x2d20fb60.
//
// Solidity: function setOwnersExplicit(uint256 quantity) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SetOwnersExplicit(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "setOwnersExplicit", quantity)
}

// SetOwnersExplicit is a paid mutator transaction binding the contract method 0x2d20fb60.
//
// Solidity: function setOwnersExplicit(uint256 quantity) returns()
func (_DrawableNft721 *DrawableNft721Session) SetOwnersExplicit(quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetOwnersExplicit(&_DrawableNft721.TransactOpts, quantity)
}

// SetOwnersExplicit is a paid mutator transaction binding the contract method 0x2d20fb60.
//
// Solidity: function setOwnersExplicit(uint256 quantity) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SetOwnersExplicit(quantity *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetOwnersExplicit(&_DrawableNft721.TransactOpts, quantity)
}

// SetupPublicSaleInfo is a paid mutator transaction binding the contract method 0xe95877c7.
//
// Solidity: function setupPublicSaleInfo(uint256 publicPriceWei, uint256 publicSaleStartTime) returns()
func (_DrawableNft721 *DrawableNft721Transactor) SetupPublicSaleInfo(opts *bind.TransactOpts, publicPriceWei *big.Int, publicSaleStartTime *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "setupPublicSaleInfo", publicPriceWei, publicSaleStartTime)
}

// SetupPublicSaleInfo is a paid mutator transaction binding the contract method 0xe95877c7.
//
// Solidity: function setupPublicSaleInfo(uint256 publicPriceWei, uint256 publicSaleStartTime) returns()
func (_DrawableNft721 *DrawableNft721Session) SetupPublicSaleInfo(publicPriceWei *big.Int, publicSaleStartTime *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetupPublicSaleInfo(&_DrawableNft721.TransactOpts, publicPriceWei, publicSaleStartTime)
}

// SetupPublicSaleInfo is a paid mutator transaction binding the contract method 0xe95877c7.
//
// Solidity: function setupPublicSaleInfo(uint256 publicPriceWei, uint256 publicSaleStartTime) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) SetupPublicSaleInfo(publicPriceWei *big.Int, publicSaleStartTime *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.SetupPublicSaleInfo(&_DrawableNft721.TransactOpts, publicPriceWei, publicSaleStartTime)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.TransferFrom(&_DrawableNft721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DrawableNft721.Contract.TransferFrom(&_DrawableNft721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_DrawableNft721 *DrawableNft721Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_DrawableNft721 *DrawableNft721Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DrawableNft721.Contract.TransferOwnership(&_DrawableNft721.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DrawableNft721.Contract.TransferOwnership(&_DrawableNft721.TransactOpts, to)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_DrawableNft721 *DrawableNft721Transactor) WithdrawMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrawableNft721.contract.Transact(opts, "withdrawMoney")
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_DrawableNft721 *DrawableNft721Session) WithdrawMoney() (*types.Transaction, error) {
	return _DrawableNft721.Contract.WithdrawMoney(&_DrawableNft721.TransactOpts)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_DrawableNft721 *DrawableNft721TransactorSession) WithdrawMoney() (*types.Transaction, error) {
	return _DrawableNft721.Contract.WithdrawMoney(&_DrawableNft721.TransactOpts)
}

// DrawableNft721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DrawableNft721 contract.
type DrawableNft721ApprovalIterator struct {
	Event *DrawableNft721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721Approval represents a Approval event raised by the DrawableNft721 contract.
type DrawableNft721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DrawableNft721 *DrawableNft721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DrawableNft721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721ApprovalIterator{contract: _DrawableNft721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DrawableNft721 *DrawableNft721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DrawableNft721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721Approval)
				if err := _DrawableNft721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DrawableNft721 *DrawableNft721Filterer) ParseApproval(log types.Log) (*DrawableNft721Approval, error) {
	event := new(DrawableNft721Approval)
	if err := _DrawableNft721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DrawableNft721 contract.
type DrawableNft721ApprovalForAllIterator struct {
	Event *DrawableNft721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721ApprovalForAll represents a ApprovalForAll event raised by the DrawableNft721 contract.
type DrawableNft721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DrawableNft721 *DrawableNft721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DrawableNft721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721ApprovalForAllIterator{contract: _DrawableNft721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DrawableNft721 *DrawableNft721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DrawableNft721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721ApprovalForAll)
				if err := _DrawableNft721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DrawableNft721 *DrawableNft721Filterer) ParseApprovalForAll(log types.Log) (*DrawableNft721ApprovalForAll, error) {
	event := new(DrawableNft721ApprovalForAll)
	if err := _DrawableNft721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721CoordinatorSetIterator is returned from FilterCoordinatorSet and is used to iterate over the raw logs and unpacked data for CoordinatorSet events raised by the DrawableNft721 contract.
type DrawableNft721CoordinatorSetIterator struct {
	Event *DrawableNft721CoordinatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721CoordinatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721CoordinatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721CoordinatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721CoordinatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721CoordinatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721CoordinatorSet represents a CoordinatorSet event raised by the DrawableNft721 contract.
type DrawableNft721CoordinatorSet struct {
	VrfCoordinator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCoordinatorSet is a free log retrieval operation binding the contract event 0xd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be6.
//
// Solidity: event CoordinatorSet(address vrfCoordinator)
func (_DrawableNft721 *DrawableNft721Filterer) FilterCoordinatorSet(opts *bind.FilterOpts) (*DrawableNft721CoordinatorSetIterator, error) {

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "CoordinatorSet")
	if err != nil {
		return nil, err
	}
	return &DrawableNft721CoordinatorSetIterator{contract: _DrawableNft721.contract, event: "CoordinatorSet", logs: logs, sub: sub}, nil
}

// WatchCoordinatorSet is a free log subscription operation binding the contract event 0xd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be6.
//
// Solidity: event CoordinatorSet(address vrfCoordinator)
func (_DrawableNft721 *DrawableNft721Filterer) WatchCoordinatorSet(opts *bind.WatchOpts, sink chan<- *DrawableNft721CoordinatorSet) (event.Subscription, error) {

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "CoordinatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721CoordinatorSet)
				if err := _DrawableNft721.contract.UnpackLog(event, "CoordinatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCoordinatorSet is a log parse operation binding the contract event 0xd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be6.
//
// Solidity: event CoordinatorSet(address vrfCoordinator)
func (_DrawableNft721 *DrawableNft721Filterer) ParseCoordinatorSet(log types.Log) (*DrawableNft721CoordinatorSet, error) {
	event := new(DrawableNft721CoordinatorSet)
	if err := _DrawableNft721.contract.UnpackLog(event, "CoordinatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721DrawPropertySetIterator is returned from FilterDrawPropertySet and is used to iterate over the raw logs and unpacked data for DrawPropertySet events raised by the DrawableNft721 contract.
type DrawableNft721DrawPropertySetIterator struct {
	Event *DrawableNft721DrawPropertySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721DrawPropertySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721DrawPropertySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721DrawPropertySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721DrawPropertySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721DrawPropertySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721DrawPropertySet represents a DrawPropertySet event raised by the DrawableNft721 contract.
type DrawableNft721DrawPropertySet struct {
	DrawId      *big.Int
	Probability []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDrawPropertySet is a free log retrieval operation binding the contract event 0xe5d2b4b0cc0c9f55ddd8d897dbacf9a1eef29fa5d5a08b5f1a8a7a074442c352.
//
// Solidity: event DrawPropertySet(uint256 drawId, uint256[] probability)
func (_DrawableNft721 *DrawableNft721Filterer) FilterDrawPropertySet(opts *bind.FilterOpts) (*DrawableNft721DrawPropertySetIterator, error) {

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "DrawPropertySet")
	if err != nil {
		return nil, err
	}
	return &DrawableNft721DrawPropertySetIterator{contract: _DrawableNft721.contract, event: "DrawPropertySet", logs: logs, sub: sub}, nil
}

// WatchDrawPropertySet is a free log subscription operation binding the contract event 0xe5d2b4b0cc0c9f55ddd8d897dbacf9a1eef29fa5d5a08b5f1a8a7a074442c352.
//
// Solidity: event DrawPropertySet(uint256 drawId, uint256[] probability)
func (_DrawableNft721 *DrawableNft721Filterer) WatchDrawPropertySet(opts *bind.WatchOpts, sink chan<- *DrawableNft721DrawPropertySet) (event.Subscription, error) {

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "DrawPropertySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721DrawPropertySet)
				if err := _DrawableNft721.contract.UnpackLog(event, "DrawPropertySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDrawPropertySet is a log parse operation binding the contract event 0xe5d2b4b0cc0c9f55ddd8d897dbacf9a1eef29fa5d5a08b5f1a8a7a074442c352.
//
// Solidity: event DrawPropertySet(uint256 drawId, uint256[] probability)
func (_DrawableNft721 *DrawableNft721Filterer) ParseDrawPropertySet(log types.Log) (*DrawableNft721DrawPropertySet, error) {
	event := new(DrawableNft721DrawPropertySet)
	if err := _DrawableNft721.contract.UnpackLog(event, "DrawPropertySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721DrawnIterator is returned from FilterDrawn and is used to iterate over the raw logs and unpacked data for Drawn events raised by the DrawableNft721 contract.
type DrawableNft721DrawnIterator struct {
	Event *DrawableNft721Drawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721DrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721Drawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721Drawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721DrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721DrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721Drawn represents a Drawn event raised by the DrawableNft721 contract.
type DrawableNft721Drawn struct {
	TokenId   *big.Int
	DrawId    *big.Int
	DrawIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDrawn is a free log retrieval operation binding the contract event 0x843520851468b1dfa3ba7caa0809ae4f500d8457b08e4303518775f999de8e51.
//
// Solidity: event Drawn(uint256 indexed tokenId, uint256 drawId, uint256 drawIndex)
func (_DrawableNft721 *DrawableNft721Filterer) FilterDrawn(opts *bind.FilterOpts, tokenId []*big.Int) (*DrawableNft721DrawnIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "Drawn", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721DrawnIterator{contract: _DrawableNft721.contract, event: "Drawn", logs: logs, sub: sub}, nil
}

// WatchDrawn is a free log subscription operation binding the contract event 0x843520851468b1dfa3ba7caa0809ae4f500d8457b08e4303518775f999de8e51.
//
// Solidity: event Drawn(uint256 indexed tokenId, uint256 drawId, uint256 drawIndex)
func (_DrawableNft721 *DrawableNft721Filterer) WatchDrawn(opts *bind.WatchOpts, sink chan<- *DrawableNft721Drawn, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "Drawn", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721Drawn)
				if err := _DrawableNft721.contract.UnpackLog(event, "Drawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDrawn is a log parse operation binding the contract event 0x843520851468b1dfa3ba7caa0809ae4f500d8457b08e4303518775f999de8e51.
//
// Solidity: event Drawn(uint256 indexed tokenId, uint256 drawId, uint256 drawIndex)
func (_DrawableNft721 *DrawableNft721Filterer) ParseDrawn(log types.Log) (*DrawableNft721Drawn, error) {
	event := new(DrawableNft721Drawn)
	if err := _DrawableNft721.contract.UnpackLog(event, "Drawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721OwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the DrawableNft721 contract.
type DrawableNft721OwnershipTransferRequestedIterator struct {
	Event *DrawableNft721OwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721OwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721OwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721OwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721OwnershipTransferRequested represents a OwnershipTransferRequested event raised by the DrawableNft721 contract.
type DrawableNft721OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_DrawableNft721 *DrawableNft721Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DrawableNft721OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721OwnershipTransferRequestedIterator{contract: _DrawableNft721.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_DrawableNft721 *DrawableNft721Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DrawableNft721OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721OwnershipTransferRequested)
				if err := _DrawableNft721.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_DrawableNft721 *DrawableNft721Filterer) ParseOwnershipTransferRequested(log types.Log) (*DrawableNft721OwnershipTransferRequested, error) {
	event := new(DrawableNft721OwnershipTransferRequested)
	if err := _DrawableNft721.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DrawableNft721 contract.
type DrawableNft721OwnershipTransferredIterator struct {
	Event *DrawableNft721OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721OwnershipTransferred represents a OwnershipTransferred event raised by the DrawableNft721 contract.
type DrawableNft721OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_DrawableNft721 *DrawableNft721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DrawableNft721OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721OwnershipTransferredIterator{contract: _DrawableNft721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_DrawableNft721 *DrawableNft721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DrawableNft721OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721OwnershipTransferred)
				if err := _DrawableNft721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_DrawableNft721 *DrawableNft721Filterer) ParseOwnershipTransferred(log types.Log) (*DrawableNft721OwnershipTransferred, error) {
	event := new(DrawableNft721OwnershipTransferred)
	if err := _DrawableNft721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721SpecialPropertySetIterator is returned from FilterSpecialPropertySet and is used to iterate over the raw logs and unpacked data for SpecialPropertySet events raised by the DrawableNft721 contract.
type DrawableNft721SpecialPropertySetIterator struct {
	Event *DrawableNft721SpecialPropertySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721SpecialPropertySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721SpecialPropertySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721SpecialPropertySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721SpecialPropertySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721SpecialPropertySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721SpecialPropertySet represents a SpecialPropertySet event raised by the DrawableNft721 contract.
type DrawableNft721SpecialPropertySet struct {
	DrawId      *big.Int
	Probability []*big.Int
	Start       *big.Int
	End         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSpecialPropertySet is a free log retrieval operation binding the contract event 0xc5d037ef0955a33827a41da60cb3236439fed82422922d5aa8d38242948b6534.
//
// Solidity: event SpecialPropertySet(uint256 drawId, uint256[] probability, uint256 start, uint256 end)
func (_DrawableNft721 *DrawableNft721Filterer) FilterSpecialPropertySet(opts *bind.FilterOpts) (*DrawableNft721SpecialPropertySetIterator, error) {

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "SpecialPropertySet")
	if err != nil {
		return nil, err
	}
	return &DrawableNft721SpecialPropertySetIterator{contract: _DrawableNft721.contract, event: "SpecialPropertySet", logs: logs, sub: sub}, nil
}

// WatchSpecialPropertySet is a free log subscription operation binding the contract event 0xc5d037ef0955a33827a41da60cb3236439fed82422922d5aa8d38242948b6534.
//
// Solidity: event SpecialPropertySet(uint256 drawId, uint256[] probability, uint256 start, uint256 end)
func (_DrawableNft721 *DrawableNft721Filterer) WatchSpecialPropertySet(opts *bind.WatchOpts, sink chan<- *DrawableNft721SpecialPropertySet) (event.Subscription, error) {

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "SpecialPropertySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721SpecialPropertySet)
				if err := _DrawableNft721.contract.UnpackLog(event, "SpecialPropertySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSpecialPropertySet is a log parse operation binding the contract event 0xc5d037ef0955a33827a41da60cb3236439fed82422922d5aa8d38242948b6534.
//
// Solidity: event SpecialPropertySet(uint256 drawId, uint256[] probability, uint256 start, uint256 end)
func (_DrawableNft721 *DrawableNft721Filterer) ParseSpecialPropertySet(log types.Log) (*DrawableNft721SpecialPropertySet, error) {
	event := new(DrawableNft721SpecialPropertySet)
	if err := _DrawableNft721.contract.UnpackLog(event, "SpecialPropertySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrawableNft721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DrawableNft721 contract.
type DrawableNft721TransferIterator struct {
	Event *DrawableNft721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DrawableNft721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrawableNft721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DrawableNft721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DrawableNft721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrawableNft721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrawableNft721Transfer represents a Transfer event raised by the DrawableNft721 contract.
type DrawableNft721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DrawableNft721 *DrawableNft721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DrawableNft721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DrawableNft721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DrawableNft721TransferIterator{contract: _DrawableNft721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DrawableNft721 *DrawableNft721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DrawableNft721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DrawableNft721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrawableNft721Transfer)
				if err := _DrawableNft721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DrawableNft721 *DrawableNft721Filterer) ParseTransfer(log types.Log) (*DrawableNft721Transfer, error) {
	event := new(DrawableNft721Transfer)
	if err := _DrawableNft721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
