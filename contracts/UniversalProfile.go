// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UniversalProfileABI is the input ABI used to generate the binding from.
const UniversalProfileABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operationType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dataValue\",\"type\":\"bytes\"}],\"name\":\"DataChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operationType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RenounceOwnershipStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receivedData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnedValue\",\"type\":\"bytes\"}],\"name\":\"UniversalReceiver\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"RENOUNCE_OWNERSHIP_CONFIRMATION_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RENOUNCE_OWNERSHIP_CONFIRMATION_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"batchCalls\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operationType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"operationsType\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataKey\",\"type\":\"bytes32\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"dataValue\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataKeys\",\"type\":\"bytes32[]\"}],\"name\":\"getDataBatch\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"dataValues\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"returnedStatus\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"dataValue\",\"type\":\"bytes\"}],\"name\":\"setData\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataKeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataValues\",\"type\":\"bytes[]\"}],\"name\":\"setDataBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingNewOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"typeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"receivedData\",\"type\":\"bytes\"}],\"name\":\"universalReceiver\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"returnedValues\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// UniversalProfile is an auto generated Go binding around an Ethereum contract.
type UniversalProfile struct {
	UniversalProfileCaller     // Read-only binding to the contract
	UniversalProfileTransactor // Write-only binding to the contract
	UniversalProfileFilterer   // Log filterer for contract events
}

// UniversalProfileCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalProfileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalProfileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalProfileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalProfileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalProfileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalProfileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalProfileSession struct {
	Contract     *UniversalProfile // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniversalProfileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalProfileCallerSession struct {
	Contract *UniversalProfileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UniversalProfileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalProfileTransactorSession struct {
	Contract     *UniversalProfileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UniversalProfileRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalProfileRaw struct {
	Contract *UniversalProfile // Generic contract binding to access the raw methods on
}

// UniversalProfileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalProfileCallerRaw struct {
	Contract *UniversalProfileCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalProfileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalProfileTransactorRaw struct {
	Contract *UniversalProfileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalProfile creates a new instance of UniversalProfile, bound to a specific deployed contract.
func NewUniversalProfile(address common.Address, backend bind.ContractBackend) (*UniversalProfile, error) {
	contract, err := bindUniversalProfile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalProfile{UniversalProfileCaller: UniversalProfileCaller{contract: contract}, UniversalProfileTransactor: UniversalProfileTransactor{contract: contract}, UniversalProfileFilterer: UniversalProfileFilterer{contract: contract}}, nil
}

// NewUniversalProfileCaller creates a new read-only instance of UniversalProfile, bound to a specific deployed contract.
func NewUniversalProfileCaller(address common.Address, caller bind.ContractCaller) (*UniversalProfileCaller, error) {
	contract, err := bindUniversalProfile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileCaller{contract: contract}, nil
}

// NewUniversalProfileTransactor creates a new write-only instance of UniversalProfile, bound to a specific deployed contract.
func NewUniversalProfileTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalProfileTransactor, error) {
	contract, err := bindUniversalProfile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileTransactor{contract: contract}, nil
}

// NewUniversalProfileFilterer creates a new log filterer instance of UniversalProfile, bound to a specific deployed contract.
func NewUniversalProfileFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalProfileFilterer, error) {
	contract, err := bindUniversalProfile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileFilterer{contract: contract}, nil
}

// bindUniversalProfile binds a generic wrapper to an already deployed contract.
func bindUniversalProfile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniversalProfileABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalProfile *UniversalProfileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalProfile.Contract.UniversalProfileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalProfile *UniversalProfileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalProfile.Contract.UniversalProfileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalProfile *UniversalProfileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalProfile.Contract.UniversalProfileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalProfile *UniversalProfileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalProfile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalProfile *UniversalProfileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalProfile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalProfile *UniversalProfileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalProfile.Contract.contract.Transact(opts, method, params...)
}

// RENOUNCEOWNERSHIPCONFIRMATIONDELAY is a free data retrieval call binding the contract method 0xead3fbdf.
//
// Solidity: function RENOUNCE_OWNERSHIP_CONFIRMATION_DELAY() view returns(uint256)
func (_UniversalProfile *UniversalProfileCaller) RENOUNCEOWNERSHIPCONFIRMATIONDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "RENOUNCE_OWNERSHIP_CONFIRMATION_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RENOUNCEOWNERSHIPCONFIRMATIONDELAY is a free data retrieval call binding the contract method 0xead3fbdf.
//
// Solidity: function RENOUNCE_OWNERSHIP_CONFIRMATION_DELAY() view returns(uint256)
func (_UniversalProfile *UniversalProfileSession) RENOUNCEOWNERSHIPCONFIRMATIONDELAY() (*big.Int, error) {
	return _UniversalProfile.Contract.RENOUNCEOWNERSHIPCONFIRMATIONDELAY(&_UniversalProfile.CallOpts)
}

// RENOUNCEOWNERSHIPCONFIRMATIONDELAY is a free data retrieval call binding the contract method 0xead3fbdf.
//
// Solidity: function RENOUNCE_OWNERSHIP_CONFIRMATION_DELAY() view returns(uint256)
func (_UniversalProfile *UniversalProfileCallerSession) RENOUNCEOWNERSHIPCONFIRMATIONDELAY() (*big.Int, error) {
	return _UniversalProfile.Contract.RENOUNCEOWNERSHIPCONFIRMATIONDELAY(&_UniversalProfile.CallOpts)
}

// RENOUNCEOWNERSHIPCONFIRMATIONPERIOD is a free data retrieval call binding the contract method 0x01bfba61.
//
// Solidity: function RENOUNCE_OWNERSHIP_CONFIRMATION_PERIOD() view returns(uint256)
func (_UniversalProfile *UniversalProfileCaller) RENOUNCEOWNERSHIPCONFIRMATIONPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "RENOUNCE_OWNERSHIP_CONFIRMATION_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RENOUNCEOWNERSHIPCONFIRMATIONPERIOD is a free data retrieval call binding the contract method 0x01bfba61.
//
// Solidity: function RENOUNCE_OWNERSHIP_CONFIRMATION_PERIOD() view returns(uint256)
func (_UniversalProfile *UniversalProfileSession) RENOUNCEOWNERSHIPCONFIRMATIONPERIOD() (*big.Int, error) {
	return _UniversalProfile.Contract.RENOUNCEOWNERSHIPCONFIRMATIONPERIOD(&_UniversalProfile.CallOpts)
}

// RENOUNCEOWNERSHIPCONFIRMATIONPERIOD is a free data retrieval call binding the contract method 0x01bfba61.
//
// Solidity: function RENOUNCE_OWNERSHIP_CONFIRMATION_PERIOD() view returns(uint256)
func (_UniversalProfile *UniversalProfileCallerSession) RENOUNCEOWNERSHIPCONFIRMATIONPERIOD() (*big.Int, error) {
	return _UniversalProfile.Contract.RENOUNCEOWNERSHIPCONFIRMATIONPERIOD(&_UniversalProfile.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalProfile *UniversalProfileCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalProfile *UniversalProfileSession) VERSION() (string, error) {
	return _UniversalProfile.Contract.VERSION(&_UniversalProfile.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalProfile *UniversalProfileCallerSession) VERSION() (string, error) {
	return _UniversalProfile.Contract.VERSION(&_UniversalProfile.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 dataKey) view returns(bytes dataValue)
func (_UniversalProfile *UniversalProfileCaller) GetData(opts *bind.CallOpts, dataKey [32]byte) ([]byte, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "getData", dataKey)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 dataKey) view returns(bytes dataValue)
func (_UniversalProfile *UniversalProfileSession) GetData(dataKey [32]byte) ([]byte, error) {
	return _UniversalProfile.Contract.GetData(&_UniversalProfile.CallOpts, dataKey)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 dataKey) view returns(bytes dataValue)
func (_UniversalProfile *UniversalProfileCallerSession) GetData(dataKey [32]byte) ([]byte, error) {
	return _UniversalProfile.Contract.GetData(&_UniversalProfile.CallOpts, dataKey)
}

// GetDataBatch is a free data retrieval call binding the contract method 0xdedff9c6.
//
// Solidity: function getDataBatch(bytes32[] dataKeys) view returns(bytes[] dataValues)
func (_UniversalProfile *UniversalProfileCaller) GetDataBatch(opts *bind.CallOpts, dataKeys [][32]byte) ([][]byte, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "getDataBatch", dataKeys)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetDataBatch is a free data retrieval call binding the contract method 0xdedff9c6.
//
// Solidity: function getDataBatch(bytes32[] dataKeys) view returns(bytes[] dataValues)
func (_UniversalProfile *UniversalProfileSession) GetDataBatch(dataKeys [][32]byte) ([][]byte, error) {
	return _UniversalProfile.Contract.GetDataBatch(&_UniversalProfile.CallOpts, dataKeys)
}

// GetDataBatch is a free data retrieval call binding the contract method 0xdedff9c6.
//
// Solidity: function getDataBatch(bytes32[] dataKeys) view returns(bytes[] dataValues)
func (_UniversalProfile *UniversalProfileCallerSession) GetDataBatch(dataKeys [][32]byte) ([][]byte, error) {
	return _UniversalProfile.Contract.GetDataBatch(&_UniversalProfile.CallOpts, dataKeys)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4 returnedStatus)
func (_UniversalProfile *UniversalProfileCaller) IsValidSignature(opts *bind.CallOpts, dataHash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "isValidSignature", dataHash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4 returnedStatus)
func (_UniversalProfile *UniversalProfileSession) IsValidSignature(dataHash [32]byte, signature []byte) ([4]byte, error) {
	return _UniversalProfile.Contract.IsValidSignature(&_UniversalProfile.CallOpts, dataHash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4 returnedStatus)
func (_UniversalProfile *UniversalProfileCallerSession) IsValidSignature(dataHash [32]byte, signature []byte) ([4]byte, error) {
	return _UniversalProfile.Contract.IsValidSignature(&_UniversalProfile.CallOpts, dataHash, signature)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalProfile *UniversalProfileCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalProfile *UniversalProfileSession) Owner() (common.Address, error) {
	return _UniversalProfile.Contract.Owner(&_UniversalProfile.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalProfile *UniversalProfileCallerSession) Owner() (common.Address, error) {
	return _UniversalProfile.Contract.Owner(&_UniversalProfile.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalProfile *UniversalProfileCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalProfile *UniversalProfileSession) PendingOwner() (common.Address, error) {
	return _UniversalProfile.Contract.PendingOwner(&_UniversalProfile.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalProfile *UniversalProfileCallerSession) PendingOwner() (common.Address, error) {
	return _UniversalProfile.Contract.PendingOwner(&_UniversalProfile.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniversalProfile *UniversalProfileCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UniversalProfile.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniversalProfile *UniversalProfileSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniversalProfile.Contract.SupportsInterface(&_UniversalProfile.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniversalProfile *UniversalProfileCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniversalProfile.Contract.SupportsInterface(&_UniversalProfile.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalProfile *UniversalProfileTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalProfile *UniversalProfileSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalProfile.Contract.AcceptOwnership(&_UniversalProfile.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalProfile *UniversalProfileTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalProfile.Contract.AcceptOwnership(&_UniversalProfile.TransactOpts)
}

// BatchCalls is a paid mutator transaction binding the contract method 0x6963d438.
//
// Solidity: function batchCalls(bytes[] data) returns(bytes[] results)
func (_UniversalProfile *UniversalProfileTransactor) BatchCalls(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "batchCalls", data)
}

// BatchCalls is a paid mutator transaction binding the contract method 0x6963d438.
//
// Solidity: function batchCalls(bytes[] data) returns(bytes[] results)
func (_UniversalProfile *UniversalProfileSession) BatchCalls(data [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.BatchCalls(&_UniversalProfile.TransactOpts, data)
}

// BatchCalls is a paid mutator transaction binding the contract method 0x6963d438.
//
// Solidity: function batchCalls(bytes[] data) returns(bytes[] results)
func (_UniversalProfile *UniversalProfileTransactorSession) BatchCalls(data [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.BatchCalls(&_UniversalProfile.TransactOpts, data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 operationType, address target, uint256 value, bytes data) payable returns(bytes)
func (_UniversalProfile *UniversalProfileTransactor) Execute(opts *bind.TransactOpts, operationType *big.Int, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "execute", operationType, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 operationType, address target, uint256 value, bytes data) payable returns(bytes)
func (_UniversalProfile *UniversalProfileSession) Execute(operationType *big.Int, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.Execute(&_UniversalProfile.TransactOpts, operationType, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 operationType, address target, uint256 value, bytes data) payable returns(bytes)
func (_UniversalProfile *UniversalProfileTransactorSession) Execute(operationType *big.Int, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.Execute(&_UniversalProfile.TransactOpts, operationType, target, value, data)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x31858452.
//
// Solidity: function executeBatch(uint256[] operationsType, address[] targets, uint256[] values, bytes[] datas) payable returns(bytes[])
func (_UniversalProfile *UniversalProfileTransactor) ExecuteBatch(opts *bind.TransactOpts, operationsType []*big.Int, targets []common.Address, values []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "executeBatch", operationsType, targets, values, datas)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x31858452.
//
// Solidity: function executeBatch(uint256[] operationsType, address[] targets, uint256[] values, bytes[] datas) payable returns(bytes[])
func (_UniversalProfile *UniversalProfileSession) ExecuteBatch(operationsType []*big.Int, targets []common.Address, values []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.ExecuteBatch(&_UniversalProfile.TransactOpts, operationsType, targets, values, datas)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x31858452.
//
// Solidity: function executeBatch(uint256[] operationsType, address[] targets, uint256[] values, bytes[] datas) payable returns(bytes[])
func (_UniversalProfile *UniversalProfileTransactorSession) ExecuteBatch(operationsType []*big.Int, targets []common.Address, values []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.ExecuteBatch(&_UniversalProfile.TransactOpts, operationsType, targets, values, datas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalProfile *UniversalProfileTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalProfile *UniversalProfileSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalProfile.Contract.RenounceOwnership(&_UniversalProfile.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalProfile *UniversalProfileTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalProfile.Contract.RenounceOwnership(&_UniversalProfile.TransactOpts)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 dataKey, bytes dataValue) payable returns()
func (_UniversalProfile *UniversalProfileTransactor) SetData(opts *bind.TransactOpts, dataKey [32]byte, dataValue []byte) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "setData", dataKey, dataValue)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 dataKey, bytes dataValue) payable returns()
func (_UniversalProfile *UniversalProfileSession) SetData(dataKey [32]byte, dataValue []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.SetData(&_UniversalProfile.TransactOpts, dataKey, dataValue)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 dataKey, bytes dataValue) payable returns()
func (_UniversalProfile *UniversalProfileTransactorSession) SetData(dataKey [32]byte, dataValue []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.SetData(&_UniversalProfile.TransactOpts, dataKey, dataValue)
}

// SetDataBatch is a paid mutator transaction binding the contract method 0x97902421.
//
// Solidity: function setDataBatch(bytes32[] dataKeys, bytes[] dataValues) payable returns()
func (_UniversalProfile *UniversalProfileTransactor) SetDataBatch(opts *bind.TransactOpts, dataKeys [][32]byte, dataValues [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "setDataBatch", dataKeys, dataValues)
}

// SetDataBatch is a paid mutator transaction binding the contract method 0x97902421.
//
// Solidity: function setDataBatch(bytes32[] dataKeys, bytes[] dataValues) payable returns()
func (_UniversalProfile *UniversalProfileSession) SetDataBatch(dataKeys [][32]byte, dataValues [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.SetDataBatch(&_UniversalProfile.TransactOpts, dataKeys, dataValues)
}

// SetDataBatch is a paid mutator transaction binding the contract method 0x97902421.
//
// Solidity: function setDataBatch(bytes32[] dataKeys, bytes[] dataValues) payable returns()
func (_UniversalProfile *UniversalProfileTransactorSession) SetDataBatch(dataKeys [][32]byte, dataValues [][]byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.SetDataBatch(&_UniversalProfile.TransactOpts, dataKeys, dataValues)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address pendingNewOwner) returns()
func (_UniversalProfile *UniversalProfileTransactor) TransferOwnership(opts *bind.TransactOpts, pendingNewOwner common.Address) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "transferOwnership", pendingNewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address pendingNewOwner) returns()
func (_UniversalProfile *UniversalProfileSession) TransferOwnership(pendingNewOwner common.Address) (*types.Transaction, error) {
	return _UniversalProfile.Contract.TransferOwnership(&_UniversalProfile.TransactOpts, pendingNewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address pendingNewOwner) returns()
func (_UniversalProfile *UniversalProfileTransactorSession) TransferOwnership(pendingNewOwner common.Address) (*types.Transaction, error) {
	return _UniversalProfile.Contract.TransferOwnership(&_UniversalProfile.TransactOpts, pendingNewOwner)
}

// UniversalReceiver is a paid mutator transaction binding the contract method 0x6bb56a14.
//
// Solidity: function universalReceiver(bytes32 typeId, bytes receivedData) payable returns(bytes returnedValues)
func (_UniversalProfile *UniversalProfileTransactor) UniversalReceiver(opts *bind.TransactOpts, typeId [32]byte, receivedData []byte) (*types.Transaction, error) {
	return _UniversalProfile.contract.Transact(opts, "universalReceiver", typeId, receivedData)
}

// UniversalReceiver is a paid mutator transaction binding the contract method 0x6bb56a14.
//
// Solidity: function universalReceiver(bytes32 typeId, bytes receivedData) payable returns(bytes returnedValues)
func (_UniversalProfile *UniversalProfileSession) UniversalReceiver(typeId [32]byte, receivedData []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.UniversalReceiver(&_UniversalProfile.TransactOpts, typeId, receivedData)
}

// UniversalReceiver is a paid mutator transaction binding the contract method 0x6bb56a14.
//
// Solidity: function universalReceiver(bytes32 typeId, bytes receivedData) payable returns(bytes returnedValues)
func (_UniversalProfile *UniversalProfileTransactorSession) UniversalReceiver(typeId [32]byte, receivedData []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.UniversalReceiver(&_UniversalProfile.TransactOpts, typeId, receivedData)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UniversalProfile *UniversalProfileTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UniversalProfile.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UniversalProfile *UniversalProfileSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.Fallback(&_UniversalProfile.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UniversalProfile *UniversalProfileTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UniversalProfile.Contract.Fallback(&_UniversalProfile.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalProfile *UniversalProfileTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalProfile.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalProfile *UniversalProfileSession) Receive() (*types.Transaction, error) {
	return _UniversalProfile.Contract.Receive(&_UniversalProfile.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalProfile *UniversalProfileTransactorSession) Receive() (*types.Transaction, error) {
	return _UniversalProfile.Contract.Receive(&_UniversalProfile.TransactOpts)
}

// UniversalProfileContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the UniversalProfile contract.
type UniversalProfileContractCreatedIterator struct {
	Event *UniversalProfileContractCreated // Event containing the contract specifics and raw log

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
func (it *UniversalProfileContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileContractCreated)
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
		it.Event = new(UniversalProfileContractCreated)
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
func (it *UniversalProfileContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileContractCreated represents a ContractCreated event raised by the UniversalProfile contract.
type UniversalProfileContractCreated struct {
	OperationType   *big.Int
	ContractAddress common.Address
	Value           *big.Int
	Salt            [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xa1fb700aaee2ae4a2ff6f91ce7eba292f89c2f5488b8ec4c5c5c8150692595c3.
//
// Solidity: event ContractCreated(uint256 indexed operationType, address indexed contractAddress, uint256 indexed value, bytes32 salt)
func (_UniversalProfile *UniversalProfileFilterer) FilterContractCreated(opts *bind.FilterOpts, operationType []*big.Int, contractAddress []common.Address, value []*big.Int) (*UniversalProfileContractCreatedIterator, error) {

	var operationTypeRule []interface{}
	for _, operationTypeItem := range operationType {
		operationTypeRule = append(operationTypeRule, operationTypeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "ContractCreated", operationTypeRule, contractAddressRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileContractCreatedIterator{contract: _UniversalProfile.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xa1fb700aaee2ae4a2ff6f91ce7eba292f89c2f5488b8ec4c5c5c8150692595c3.
//
// Solidity: event ContractCreated(uint256 indexed operationType, address indexed contractAddress, uint256 indexed value, bytes32 salt)
func (_UniversalProfile *UniversalProfileFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *UniversalProfileContractCreated, operationType []*big.Int, contractAddress []common.Address, value []*big.Int) (event.Subscription, error) {

	var operationTypeRule []interface{}
	for _, operationTypeItem := range operationType {
		operationTypeRule = append(operationTypeRule, operationTypeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "ContractCreated", operationTypeRule, contractAddressRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileContractCreated)
				if err := _UniversalProfile.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ParseContractCreated is a log parse operation binding the contract event 0xa1fb700aaee2ae4a2ff6f91ce7eba292f89c2f5488b8ec4c5c5c8150692595c3.
//
// Solidity: event ContractCreated(uint256 indexed operationType, address indexed contractAddress, uint256 indexed value, bytes32 salt)
func (_UniversalProfile *UniversalProfileFilterer) ParseContractCreated(log types.Log) (*UniversalProfileContractCreated, error) {
	event := new(UniversalProfileContractCreated)
	if err := _UniversalProfile.contract.UnpackLog(event, "ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalProfileDataChangedIterator is returned from FilterDataChanged and is used to iterate over the raw logs and unpacked data for DataChanged events raised by the UniversalProfile contract.
type UniversalProfileDataChangedIterator struct {
	Event *UniversalProfileDataChanged // Event containing the contract specifics and raw log

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
func (it *UniversalProfileDataChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileDataChanged)
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
		it.Event = new(UniversalProfileDataChanged)
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
func (it *UniversalProfileDataChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileDataChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileDataChanged represents a DataChanged event raised by the UniversalProfile contract.
type UniversalProfileDataChanged struct {
	DataKey   [32]byte
	DataValue []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDataChanged is a free log retrieval operation binding the contract event 0xece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2.
//
// Solidity: event DataChanged(bytes32 indexed dataKey, bytes dataValue)
func (_UniversalProfile *UniversalProfileFilterer) FilterDataChanged(opts *bind.FilterOpts, dataKey [][32]byte) (*UniversalProfileDataChangedIterator, error) {

	var dataKeyRule []interface{}
	for _, dataKeyItem := range dataKey {
		dataKeyRule = append(dataKeyRule, dataKeyItem)
	}

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "DataChanged", dataKeyRule)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileDataChangedIterator{contract: _UniversalProfile.contract, event: "DataChanged", logs: logs, sub: sub}, nil
}

// WatchDataChanged is a free log subscription operation binding the contract event 0xece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2.
//
// Solidity: event DataChanged(bytes32 indexed dataKey, bytes dataValue)
func (_UniversalProfile *UniversalProfileFilterer) WatchDataChanged(opts *bind.WatchOpts, sink chan<- *UniversalProfileDataChanged, dataKey [][32]byte) (event.Subscription, error) {

	var dataKeyRule []interface{}
	for _, dataKeyItem := range dataKey {
		dataKeyRule = append(dataKeyRule, dataKeyItem)
	}

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "DataChanged", dataKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileDataChanged)
				if err := _UniversalProfile.contract.UnpackLog(event, "DataChanged", log); err != nil {
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

// ParseDataChanged is a log parse operation binding the contract event 0xece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2.
//
// Solidity: event DataChanged(bytes32 indexed dataKey, bytes dataValue)
func (_UniversalProfile *UniversalProfileFilterer) ParseDataChanged(log types.Log) (*UniversalProfileDataChanged, error) {
	event := new(UniversalProfileDataChanged)
	if err := _UniversalProfile.contract.UnpackLog(event, "DataChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalProfileExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the UniversalProfile contract.
type UniversalProfileExecutedIterator struct {
	Event *UniversalProfileExecuted // Event containing the contract specifics and raw log

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
func (it *UniversalProfileExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileExecuted)
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
		it.Event = new(UniversalProfileExecuted)
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
func (it *UniversalProfileExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileExecuted represents a Executed event raised by the UniversalProfile contract.
type UniversalProfileExecuted struct {
	OperationType *big.Int
	Target        common.Address
	Value         *big.Int
	Selector      [4]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x4810874456b8e6487bd861375cf6abd8e1c8bb5858c8ce36a86a04dabfac199e.
//
// Solidity: event Executed(uint256 indexed operationType, address indexed target, uint256 indexed value, bytes4 selector)
func (_UniversalProfile *UniversalProfileFilterer) FilterExecuted(opts *bind.FilterOpts, operationType []*big.Int, target []common.Address, value []*big.Int) (*UniversalProfileExecutedIterator, error) {

	var operationTypeRule []interface{}
	for _, operationTypeItem := range operationType {
		operationTypeRule = append(operationTypeRule, operationTypeItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "Executed", operationTypeRule, targetRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileExecutedIterator{contract: _UniversalProfile.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x4810874456b8e6487bd861375cf6abd8e1c8bb5858c8ce36a86a04dabfac199e.
//
// Solidity: event Executed(uint256 indexed operationType, address indexed target, uint256 indexed value, bytes4 selector)
func (_UniversalProfile *UniversalProfileFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *UniversalProfileExecuted, operationType []*big.Int, target []common.Address, value []*big.Int) (event.Subscription, error) {

	var operationTypeRule []interface{}
	for _, operationTypeItem := range operationType {
		operationTypeRule = append(operationTypeRule, operationTypeItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "Executed", operationTypeRule, targetRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileExecuted)
				if err := _UniversalProfile.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x4810874456b8e6487bd861375cf6abd8e1c8bb5858c8ce36a86a04dabfac199e.
//
// Solidity: event Executed(uint256 indexed operationType, address indexed target, uint256 indexed value, bytes4 selector)
func (_UniversalProfile *UniversalProfileFilterer) ParseExecuted(log types.Log) (*UniversalProfileExecuted, error) {
	event := new(UniversalProfileExecuted)
	if err := _UniversalProfile.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalProfileOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the UniversalProfile contract.
type UniversalProfileOwnershipRenouncedIterator struct {
	Event *UniversalProfileOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *UniversalProfileOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileOwnershipRenounced)
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
		it.Event = new(UniversalProfileOwnershipRenounced)
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
func (it *UniversalProfileOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileOwnershipRenounced represents a OwnershipRenounced event raised by the UniversalProfile contract.
type UniversalProfileOwnershipRenounced struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xd1f66c3d2bc1993a86be5e3d33709d98f0442381befcedd29f578b9b2506b1ce.
//
// Solidity: event OwnershipRenounced()
func (_UniversalProfile *UniversalProfileFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts) (*UniversalProfileOwnershipRenouncedIterator, error) {

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "OwnershipRenounced")
	if err != nil {
		return nil, err
	}
	return &UniversalProfileOwnershipRenouncedIterator{contract: _UniversalProfile.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xd1f66c3d2bc1993a86be5e3d33709d98f0442381befcedd29f578b9b2506b1ce.
//
// Solidity: event OwnershipRenounced()
func (_UniversalProfile *UniversalProfileFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *UniversalProfileOwnershipRenounced) (event.Subscription, error) {

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "OwnershipRenounced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileOwnershipRenounced)
				if err := _UniversalProfile.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xd1f66c3d2bc1993a86be5e3d33709d98f0442381befcedd29f578b9b2506b1ce.
//
// Solidity: event OwnershipRenounced()
func (_UniversalProfile *UniversalProfileFilterer) ParseOwnershipRenounced(log types.Log) (*UniversalProfileOwnershipRenounced, error) {
	event := new(UniversalProfileOwnershipRenounced)
	if err := _UniversalProfile.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalProfileOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the UniversalProfile contract.
type UniversalProfileOwnershipTransferStartedIterator struct {
	Event *UniversalProfileOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *UniversalProfileOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileOwnershipTransferStarted)
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
		it.Event = new(UniversalProfileOwnershipTransferStarted)
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
func (it *UniversalProfileOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the UniversalProfile contract.
type UniversalProfileOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalProfile *UniversalProfileFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalProfileOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileOwnershipTransferStartedIterator{contract: _UniversalProfile.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalProfile *UniversalProfileFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *UniversalProfileOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileOwnershipTransferStarted)
				if err := _UniversalProfile.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalProfile *UniversalProfileFilterer) ParseOwnershipTransferStarted(log types.Log) (*UniversalProfileOwnershipTransferStarted, error) {
	event := new(UniversalProfileOwnershipTransferStarted)
	if err := _UniversalProfile.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalProfileOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniversalProfile contract.
type UniversalProfileOwnershipTransferredIterator struct {
	Event *UniversalProfileOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniversalProfileOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileOwnershipTransferred)
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
		it.Event = new(UniversalProfileOwnershipTransferred)
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
func (it *UniversalProfileOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileOwnershipTransferred represents a OwnershipTransferred event raised by the UniversalProfile contract.
type UniversalProfileOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalProfile *UniversalProfileFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalProfileOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileOwnershipTransferredIterator{contract: _UniversalProfile.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalProfile *UniversalProfileFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniversalProfileOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileOwnershipTransferred)
				if err := _UniversalProfile.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalProfile *UniversalProfileFilterer) ParseOwnershipTransferred(log types.Log) (*UniversalProfileOwnershipTransferred, error) {
	event := new(UniversalProfileOwnershipTransferred)
	if err := _UniversalProfile.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalProfileRenounceOwnershipStartedIterator is returned from FilterRenounceOwnershipStarted and is used to iterate over the raw logs and unpacked data for RenounceOwnershipStarted events raised by the UniversalProfile contract.
type UniversalProfileRenounceOwnershipStartedIterator struct {
	Event *UniversalProfileRenounceOwnershipStarted // Event containing the contract specifics and raw log

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
func (it *UniversalProfileRenounceOwnershipStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileRenounceOwnershipStarted)
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
		it.Event = new(UniversalProfileRenounceOwnershipStarted)
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
func (it *UniversalProfileRenounceOwnershipStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileRenounceOwnershipStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileRenounceOwnershipStarted represents a RenounceOwnershipStarted event raised by the UniversalProfile contract.
type UniversalProfileRenounceOwnershipStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRenounceOwnershipStarted is a free log retrieval operation binding the contract event 0x81b7f830f1f0084db6497c486cbe6974c86488dcc4e3738eab94ab6d6b1653e7.
//
// Solidity: event RenounceOwnershipStarted()
func (_UniversalProfile *UniversalProfileFilterer) FilterRenounceOwnershipStarted(opts *bind.FilterOpts) (*UniversalProfileRenounceOwnershipStartedIterator, error) {

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "RenounceOwnershipStarted")
	if err != nil {
		return nil, err
	}
	return &UniversalProfileRenounceOwnershipStartedIterator{contract: _UniversalProfile.contract, event: "RenounceOwnershipStarted", logs: logs, sub: sub}, nil
}

// WatchRenounceOwnershipStarted is a free log subscription operation binding the contract event 0x81b7f830f1f0084db6497c486cbe6974c86488dcc4e3738eab94ab6d6b1653e7.
//
// Solidity: event RenounceOwnershipStarted()
func (_UniversalProfile *UniversalProfileFilterer) WatchRenounceOwnershipStarted(opts *bind.WatchOpts, sink chan<- *UniversalProfileRenounceOwnershipStarted) (event.Subscription, error) {

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "RenounceOwnershipStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileRenounceOwnershipStarted)
				if err := _UniversalProfile.contract.UnpackLog(event, "RenounceOwnershipStarted", log); err != nil {
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

// ParseRenounceOwnershipStarted is a log parse operation binding the contract event 0x81b7f830f1f0084db6497c486cbe6974c86488dcc4e3738eab94ab6d6b1653e7.
//
// Solidity: event RenounceOwnershipStarted()
func (_UniversalProfile *UniversalProfileFilterer) ParseRenounceOwnershipStarted(log types.Log) (*UniversalProfileRenounceOwnershipStarted, error) {
	event := new(UniversalProfileRenounceOwnershipStarted)
	if err := _UniversalProfile.contract.UnpackLog(event, "RenounceOwnershipStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalProfileUniversalReceiverIterator is returned from FilterUniversalReceiver and is used to iterate over the raw logs and unpacked data for UniversalReceiver events raised by the UniversalProfile contract.
type UniversalProfileUniversalReceiverIterator struct {
	Event *UniversalProfileUniversalReceiver // Event containing the contract specifics and raw log

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
func (it *UniversalProfileUniversalReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalProfileUniversalReceiver)
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
		it.Event = new(UniversalProfileUniversalReceiver)
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
func (it *UniversalProfileUniversalReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalProfileUniversalReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalProfileUniversalReceiver represents a UniversalReceiver event raised by the UniversalProfile contract.
type UniversalProfileUniversalReceiver struct {
	From          common.Address
	Value         *big.Int
	TypeId        [32]byte
	ReceivedData  []byte
	ReturnedValue []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUniversalReceiver is a free log retrieval operation binding the contract event 0x9c3ba68eb5742b8e3961aea0afc7371a71bf433c8a67a831803b64c064a178c2.
//
// Solidity: event UniversalReceiver(address indexed from, uint256 indexed value, bytes32 indexed typeId, bytes receivedData, bytes returnedValue)
func (_UniversalProfile *UniversalProfileFilterer) FilterUniversalReceiver(opts *bind.FilterOpts, from []common.Address, value []*big.Int, typeId [][32]byte) (*UniversalProfileUniversalReceiverIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var typeIdRule []interface{}
	for _, typeIdItem := range typeId {
		typeIdRule = append(typeIdRule, typeIdItem)
	}

	logs, sub, err := _UniversalProfile.contract.FilterLogs(opts, "UniversalReceiver", fromRule, valueRule, typeIdRule)
	if err != nil {
		return nil, err
	}
	return &UniversalProfileUniversalReceiverIterator{contract: _UniversalProfile.contract, event: "UniversalReceiver", logs: logs, sub: sub}, nil
}

// WatchUniversalReceiver is a free log subscription operation binding the contract event 0x9c3ba68eb5742b8e3961aea0afc7371a71bf433c8a67a831803b64c064a178c2.
//
// Solidity: event UniversalReceiver(address indexed from, uint256 indexed value, bytes32 indexed typeId, bytes receivedData, bytes returnedValue)
func (_UniversalProfile *UniversalProfileFilterer) WatchUniversalReceiver(opts *bind.WatchOpts, sink chan<- *UniversalProfileUniversalReceiver, from []common.Address, value []*big.Int, typeId [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var typeIdRule []interface{}
	for _, typeIdItem := range typeId {
		typeIdRule = append(typeIdRule, typeIdItem)
	}

	logs, sub, err := _UniversalProfile.contract.WatchLogs(opts, "UniversalReceiver", fromRule, valueRule, typeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalProfileUniversalReceiver)
				if err := _UniversalProfile.contract.UnpackLog(event, "UniversalReceiver", log); err != nil {
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

// ParseUniversalReceiver is a log parse operation binding the contract event 0x9c3ba68eb5742b8e3961aea0afc7371a71bf433c8a67a831803b64c064a178c2.
//
// Solidity: event UniversalReceiver(address indexed from, uint256 indexed value, bytes32 indexed typeId, bytes receivedData, bytes returnedValue)
func (_UniversalProfile *UniversalProfileFilterer) ParseUniversalReceiver(log types.Log) (*UniversalProfileUniversalReceiver, error) {
	event := new(UniversalProfileUniversalReceiver)
	if err := _UniversalProfile.contract.UnpackLog(event, "UniversalReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
