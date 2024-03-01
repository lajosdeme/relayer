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

// LSP6ABI is the input ABI used to generate the binding from.
const LSP6ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"PermissionsVerified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validityTimestamps\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"executeRelayCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"validityTimestamps\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"}],\"name\":\"executeRelayCallBatch\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"channelId\",\"type\":\"uint128\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"returnedStatus\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"lsp20VerifyCall\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"lsp20VerifyCallResult\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LSP6 is an auto generated Go binding around an Ethereum contract.
type LSP6 struct {
	LSP6Caller     // Read-only binding to the contract
	LSP6Transactor // Write-only binding to the contract
	LSP6Filterer   // Log filterer for contract events
}

// LSP6Caller is an auto generated read-only Go binding around an Ethereum contract.
type LSP6Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LSP6Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LSP6Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LSP6Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LSP6Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LSP6Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LSP6Session struct {
	Contract     *LSP6             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LSP6CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LSP6CallerSession struct {
	Contract *LSP6Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LSP6TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LSP6TransactorSession struct {
	Contract     *LSP6Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LSP6Raw is an auto generated low-level Go binding around an Ethereum contract.
type LSP6Raw struct {
	Contract *LSP6 // Generic contract binding to access the raw methods on
}

// LSP6CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LSP6CallerRaw struct {
	Contract *LSP6Caller // Generic read-only contract binding to access the raw methods on
}

// LSP6TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LSP6TransactorRaw struct {
	Contract *LSP6Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLSP6 creates a new instance of LSP6, bound to a specific deployed contract.
func NewLSP6(address common.Address, backend bind.ContractBackend) (*LSP6, error) {
	contract, err := bindLSP6(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LSP6{LSP6Caller: LSP6Caller{contract: contract}, LSP6Transactor: LSP6Transactor{contract: contract}, LSP6Filterer: LSP6Filterer{contract: contract}}, nil
}

// NewLSP6Caller creates a new read-only instance of LSP6, bound to a specific deployed contract.
func NewLSP6Caller(address common.Address, caller bind.ContractCaller) (*LSP6Caller, error) {
	contract, err := bindLSP6(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LSP6Caller{contract: contract}, nil
}

// NewLSP6Transactor creates a new write-only instance of LSP6, bound to a specific deployed contract.
func NewLSP6Transactor(address common.Address, transactor bind.ContractTransactor) (*LSP6Transactor, error) {
	contract, err := bindLSP6(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LSP6Transactor{contract: contract}, nil
}

// NewLSP6Filterer creates a new log filterer instance of LSP6, bound to a specific deployed contract.
func NewLSP6Filterer(address common.Address, filterer bind.ContractFilterer) (*LSP6Filterer, error) {
	contract, err := bindLSP6(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LSP6Filterer{contract: contract}, nil
}

// bindLSP6 binds a generic wrapper to an already deployed contract.
func bindLSP6(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LSP6ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LSP6 *LSP6Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LSP6.Contract.LSP6Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LSP6 *LSP6Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LSP6.Contract.LSP6Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LSP6 *LSP6Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LSP6.Contract.LSP6Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LSP6 *LSP6CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LSP6.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LSP6 *LSP6TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LSP6.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LSP6 *LSP6TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LSP6.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_LSP6 *LSP6Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LSP6.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_LSP6 *LSP6Session) VERSION() (string, error) {
	return _LSP6.Contract.VERSION(&_LSP6.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_LSP6 *LSP6CallerSession) VERSION() (string, error) {
	return _LSP6.Contract.VERSION(&_LSP6.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xb44581d9.
//
// Solidity: function getNonce(address from, uint128 channelId) view returns(uint256)
func (_LSP6 *LSP6Caller) GetNonce(opts *bind.CallOpts, from common.Address, channelId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LSP6.contract.Call(opts, &out, "getNonce", from, channelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xb44581d9.
//
// Solidity: function getNonce(address from, uint128 channelId) view returns(uint256)
func (_LSP6 *LSP6Session) GetNonce(from common.Address, channelId *big.Int) (*big.Int, error) {
	return _LSP6.Contract.GetNonce(&_LSP6.CallOpts, from, channelId)
}

// GetNonce is a free data retrieval call binding the contract method 0xb44581d9.
//
// Solidity: function getNonce(address from, uint128 channelId) view returns(uint256)
func (_LSP6 *LSP6CallerSession) GetNonce(from common.Address, channelId *big.Int) (*big.Int, error) {
	return _LSP6.Contract.GetNonce(&_LSP6.CallOpts, from, channelId)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4 returnedStatus)
func (_LSP6 *LSP6Caller) IsValidSignature(opts *bind.CallOpts, dataHash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _LSP6.contract.Call(opts, &out, "isValidSignature", dataHash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4 returnedStatus)
func (_LSP6 *LSP6Session) IsValidSignature(dataHash [32]byte, signature []byte) ([4]byte, error) {
	return _LSP6.Contract.IsValidSignature(&_LSP6.CallOpts, dataHash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4 returnedStatus)
func (_LSP6 *LSP6CallerSession) IsValidSignature(dataHash [32]byte, signature []byte) ([4]byte, error) {
	return _LSP6.Contract.IsValidSignature(&_LSP6.CallOpts, dataHash, signature)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LSP6 *LSP6Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LSP6.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LSP6 *LSP6Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LSP6.Contract.SupportsInterface(&_LSP6.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LSP6 *LSP6CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LSP6.Contract.SupportsInterface(&_LSP6.CallOpts, interfaceId)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_LSP6 *LSP6Caller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LSP6.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_LSP6 *LSP6Session) Target() (common.Address, error) {
	return _LSP6.Contract.Target(&_LSP6.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_LSP6 *LSP6CallerSession) Target() (common.Address, error) {
	return _LSP6.Contract.Target(&_LSP6.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes payload) payable returns(bytes)
func (_LSP6 *LSP6Transactor) Execute(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _LSP6.contract.Transact(opts, "execute", payload)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes payload) payable returns(bytes)
func (_LSP6 *LSP6Session) Execute(payload []byte) (*types.Transaction, error) {
	return _LSP6.Contract.Execute(&_LSP6.TransactOpts, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes payload) payable returns(bytes)
func (_LSP6 *LSP6TransactorSession) Execute(payload []byte) (*types.Transaction, error) {
	return _LSP6.Contract.Execute(&_LSP6.TransactOpts, payload)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xbf0176ff.
//
// Solidity: function executeBatch(uint256[] values, bytes[] payloads) payable returns(bytes[])
func (_LSP6 *LSP6Transactor) ExecuteBatch(opts *bind.TransactOpts, values []*big.Int, payloads [][]byte) (*types.Transaction, error) {
	return _LSP6.contract.Transact(opts, "executeBatch", values, payloads)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xbf0176ff.
//
// Solidity: function executeBatch(uint256[] values, bytes[] payloads) payable returns(bytes[])
func (_LSP6 *LSP6Session) ExecuteBatch(values []*big.Int, payloads [][]byte) (*types.Transaction, error) {
	return _LSP6.Contract.ExecuteBatch(&_LSP6.TransactOpts, values, payloads)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xbf0176ff.
//
// Solidity: function executeBatch(uint256[] values, bytes[] payloads) payable returns(bytes[])
func (_LSP6 *LSP6TransactorSession) ExecuteBatch(values []*big.Int, payloads [][]byte) (*types.Transaction, error) {
	return _LSP6.Contract.ExecuteBatch(&_LSP6.TransactOpts, values, payloads)
}

// ExecuteRelayCall is a paid mutator transaction binding the contract method 0x4c8a4e74.
//
// Solidity: function executeRelayCall(bytes signature, uint256 nonce, uint256 validityTimestamps, bytes payload) payable returns(bytes)
func (_LSP6 *LSP6Transactor) ExecuteRelayCall(opts *bind.TransactOpts, signature []byte, nonce *big.Int, validityTimestamps *big.Int, payload []byte) (*types.Transaction, error) {
	return _LSP6.contract.Transact(opts, "executeRelayCall", signature, nonce, validityTimestamps, payload)
}

// ExecuteRelayCall is a paid mutator transaction binding the contract method 0x4c8a4e74.
//
// Solidity: function executeRelayCall(bytes signature, uint256 nonce, uint256 validityTimestamps, bytes payload) payable returns(bytes)
func (_LSP6 *LSP6Session) ExecuteRelayCall(signature []byte, nonce *big.Int, validityTimestamps *big.Int, payload []byte) (*types.Transaction, error) {
	return _LSP6.Contract.ExecuteRelayCall(&_LSP6.TransactOpts, signature, nonce, validityTimestamps, payload)
}

// ExecuteRelayCall is a paid mutator transaction binding the contract method 0x4c8a4e74.
//
// Solidity: function executeRelayCall(bytes signature, uint256 nonce, uint256 validityTimestamps, bytes payload) payable returns(bytes)
func (_LSP6 *LSP6TransactorSession) ExecuteRelayCall(signature []byte, nonce *big.Int, validityTimestamps *big.Int, payload []byte) (*types.Transaction, error) {
	return _LSP6.Contract.ExecuteRelayCall(&_LSP6.TransactOpts, signature, nonce, validityTimestamps, payload)
}

// ExecuteRelayCallBatch is a paid mutator transaction binding the contract method 0xa20856a5.
//
// Solidity: function executeRelayCallBatch(bytes[] signatures, uint256[] nonces, uint256[] validityTimestamps, uint256[] values, bytes[] payloads) payable returns(bytes[])
func (_LSP6 *LSP6Transactor) ExecuteRelayCallBatch(opts *bind.TransactOpts, signatures [][]byte, nonces []*big.Int, validityTimestamps []*big.Int, values []*big.Int, payloads [][]byte) (*types.Transaction, error) {
	return _LSP6.contract.Transact(opts, "executeRelayCallBatch", signatures, nonces, validityTimestamps, values, payloads)
}

// ExecuteRelayCallBatch is a paid mutator transaction binding the contract method 0xa20856a5.
//
// Solidity: function executeRelayCallBatch(bytes[] signatures, uint256[] nonces, uint256[] validityTimestamps, uint256[] values, bytes[] payloads) payable returns(bytes[])
func (_LSP6 *LSP6Session) ExecuteRelayCallBatch(signatures [][]byte, nonces []*big.Int, validityTimestamps []*big.Int, values []*big.Int, payloads [][]byte) (*types.Transaction, error) {
	return _LSP6.Contract.ExecuteRelayCallBatch(&_LSP6.TransactOpts, signatures, nonces, validityTimestamps, values, payloads)
}

// ExecuteRelayCallBatch is a paid mutator transaction binding the contract method 0xa20856a5.
//
// Solidity: function executeRelayCallBatch(bytes[] signatures, uint256[] nonces, uint256[] validityTimestamps, uint256[] values, bytes[] payloads) payable returns(bytes[])
func (_LSP6 *LSP6TransactorSession) ExecuteRelayCallBatch(signatures [][]byte, nonces []*big.Int, validityTimestamps []*big.Int, values []*big.Int, payloads [][]byte) (*types.Transaction, error) {
	return _LSP6.Contract.ExecuteRelayCallBatch(&_LSP6.TransactOpts, signatures, nonces, validityTimestamps, values, payloads)
}

// Lsp20VerifyCall is a paid mutator transaction binding the contract method 0xde928f14.
//
// Solidity: function lsp20VerifyCall(address , address targetContract, address caller, uint256 msgValue, bytes callData) returns(bytes4)
func (_LSP6 *LSP6Transactor) Lsp20VerifyCall(opts *bind.TransactOpts, arg0 common.Address, targetContract common.Address, caller common.Address, msgValue *big.Int, callData []byte) (*types.Transaction, error) {
	return _LSP6.contract.Transact(opts, "lsp20VerifyCall", arg0, targetContract, caller, msgValue, callData)
}

// Lsp20VerifyCall is a paid mutator transaction binding the contract method 0xde928f14.
//
// Solidity: function lsp20VerifyCall(address , address targetContract, address caller, uint256 msgValue, bytes callData) returns(bytes4)
func (_LSP6 *LSP6Session) Lsp20VerifyCall(arg0 common.Address, targetContract common.Address, caller common.Address, msgValue *big.Int, callData []byte) (*types.Transaction, error) {
	return _LSP6.Contract.Lsp20VerifyCall(&_LSP6.TransactOpts, arg0, targetContract, caller, msgValue, callData)
}

// Lsp20VerifyCall is a paid mutator transaction binding the contract method 0xde928f14.
//
// Solidity: function lsp20VerifyCall(address , address targetContract, address caller, uint256 msgValue, bytes callData) returns(bytes4)
func (_LSP6 *LSP6TransactorSession) Lsp20VerifyCall(arg0 common.Address, targetContract common.Address, caller common.Address, msgValue *big.Int, callData []byte) (*types.Transaction, error) {
	return _LSP6.Contract.Lsp20VerifyCall(&_LSP6.TransactOpts, arg0, targetContract, caller, msgValue, callData)
}

// Lsp20VerifyCallResult is a paid mutator transaction binding the contract method 0xd3fc45d3.
//
// Solidity: function lsp20VerifyCallResult(bytes32 , bytes ) returns(bytes4)
func (_LSP6 *LSP6Transactor) Lsp20VerifyCallResult(opts *bind.TransactOpts, arg0 [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _LSP6.contract.Transact(opts, "lsp20VerifyCallResult", arg0, arg1)
}

// Lsp20VerifyCallResult is a paid mutator transaction binding the contract method 0xd3fc45d3.
//
// Solidity: function lsp20VerifyCallResult(bytes32 , bytes ) returns(bytes4)
func (_LSP6 *LSP6Session) Lsp20VerifyCallResult(arg0 [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _LSP6.Contract.Lsp20VerifyCallResult(&_LSP6.TransactOpts, arg0, arg1)
}

// Lsp20VerifyCallResult is a paid mutator transaction binding the contract method 0xd3fc45d3.
//
// Solidity: function lsp20VerifyCallResult(bytes32 , bytes ) returns(bytes4)
func (_LSP6 *LSP6TransactorSession) Lsp20VerifyCallResult(arg0 [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _LSP6.Contract.Lsp20VerifyCallResult(&_LSP6.TransactOpts, arg0, arg1)
}

// LSP6PermissionsVerifiedIterator is returned from FilterPermissionsVerified and is used to iterate over the raw logs and unpacked data for PermissionsVerified events raised by the LSP6 contract.
type LSP6PermissionsVerifiedIterator struct {
	Event *LSP6PermissionsVerified // Event containing the contract specifics and raw log

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
func (it *LSP6PermissionsVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LSP6PermissionsVerified)
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
		it.Event = new(LSP6PermissionsVerified)
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
func (it *LSP6PermissionsVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LSP6PermissionsVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LSP6PermissionsVerified represents a PermissionsVerified event raised by the LSP6 contract.
type LSP6PermissionsVerified struct {
	Signer   common.Address
	Value    *big.Int
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPermissionsVerified is a free log retrieval operation binding the contract event 0xc0a62328f6bf5e3172bb1fcb2019f54b2c523b6a48e3513a2298fbf0150b781e.
//
// Solidity: event PermissionsVerified(address indexed signer, uint256 indexed value, bytes4 indexed selector)
func (_LSP6 *LSP6Filterer) FilterPermissionsVerified(opts *bind.FilterOpts, signer []common.Address, value []*big.Int, selector [][4]byte) (*LSP6PermissionsVerifiedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _LSP6.contract.FilterLogs(opts, "PermissionsVerified", signerRule, valueRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &LSP6PermissionsVerifiedIterator{contract: _LSP6.contract, event: "PermissionsVerified", logs: logs, sub: sub}, nil
}

// WatchPermissionsVerified is a free log subscription operation binding the contract event 0xc0a62328f6bf5e3172bb1fcb2019f54b2c523b6a48e3513a2298fbf0150b781e.
//
// Solidity: event PermissionsVerified(address indexed signer, uint256 indexed value, bytes4 indexed selector)
func (_LSP6 *LSP6Filterer) WatchPermissionsVerified(opts *bind.WatchOpts, sink chan<- *LSP6PermissionsVerified, signer []common.Address, value []*big.Int, selector [][4]byte) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _LSP6.contract.WatchLogs(opts, "PermissionsVerified", signerRule, valueRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LSP6PermissionsVerified)
				if err := _LSP6.contract.UnpackLog(event, "PermissionsVerified", log); err != nil {
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

// ParsePermissionsVerified is a log parse operation binding the contract event 0xc0a62328f6bf5e3172bb1fcb2019f54b2c523b6a48e3513a2298fbf0150b781e.
//
// Solidity: event PermissionsVerified(address indexed signer, uint256 indexed value, bytes4 indexed selector)
func (_LSP6 *LSP6Filterer) ParsePermissionsVerified(log types.Log) (*LSP6PermissionsVerified, error) {
	event := new(LSP6PermissionsVerified)
	if err := _LSP6.contract.UnpackLog(event, "PermissionsVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
