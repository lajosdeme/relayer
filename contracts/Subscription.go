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

// SubscriptionABI is the input ABI used to generate the binding from.
const SubscriptionABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subscriptionPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"NewSubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_userId\",\"type\":\"string\"}],\"name\":\"isSubscribed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"setSubscriptionPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_userId\",\"type\":\"string\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subscriptionPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Subscription is an auto generated Go binding around an Ethereum contract.
type Subscription struct {
	SubscriptionCaller     // Read-only binding to the contract
	SubscriptionTransactor // Write-only binding to the contract
	SubscriptionFilterer   // Log filterer for contract events
}

// SubscriptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubscriptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscriptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscriptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscriptionSession struct {
	Contract     *Subscription     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubscriptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscriptionCallerSession struct {
	Contract *SubscriptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SubscriptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscriptionTransactorSession struct {
	Contract     *SubscriptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SubscriptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubscriptionRaw struct {
	Contract *Subscription // Generic contract binding to access the raw methods on
}

// SubscriptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscriptionCallerRaw struct {
	Contract *SubscriptionCaller // Generic read-only contract binding to access the raw methods on
}

// SubscriptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscriptionTransactorRaw struct {
	Contract *SubscriptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscription creates a new instance of Subscription, bound to a specific deployed contract.
func NewSubscription(address common.Address, backend bind.ContractBackend) (*Subscription, error) {
	contract, err := bindSubscription(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subscription{SubscriptionCaller: SubscriptionCaller{contract: contract}, SubscriptionTransactor: SubscriptionTransactor{contract: contract}, SubscriptionFilterer: SubscriptionFilterer{contract: contract}}, nil
}

// NewSubscriptionCaller creates a new read-only instance of Subscription, bound to a specific deployed contract.
func NewSubscriptionCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionCaller, error) {
	contract, err := bindSubscription(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionCaller{contract: contract}, nil
}

// NewSubscriptionTransactor creates a new write-only instance of Subscription, bound to a specific deployed contract.
func NewSubscriptionTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionTransactor, error) {
	contract, err := bindSubscription(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionTransactor{contract: contract}, nil
}

// NewSubscriptionFilterer creates a new log filterer instance of Subscription, bound to a specific deployed contract.
func NewSubscriptionFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionFilterer, error) {
	contract, err := bindSubscription(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionFilterer{contract: contract}, nil
}

// bindSubscription binds a generic wrapper to an already deployed contract.
func bindSubscription(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscription *SubscriptionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subscription.Contract.SubscriptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscription *SubscriptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.Contract.SubscriptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscription *SubscriptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscription.Contract.SubscriptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscription *SubscriptionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subscription.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscription *SubscriptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscription *SubscriptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscription.Contract.contract.Transact(opts, method, params...)
}

// IsSubscribed is a free data retrieval call binding the contract method 0xee9c32be.
//
// Solidity: function isSubscribed(string _userId) view returns(bool)
func (_Subscription *SubscriptionCaller) IsSubscribed(opts *bind.CallOpts, _userId string) (bool, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "isSubscribed", _userId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubscribed is a free data retrieval call binding the contract method 0xee9c32be.
//
// Solidity: function isSubscribed(string _userId) view returns(bool)
func (_Subscription *SubscriptionSession) IsSubscribed(_userId string) (bool, error) {
	return _Subscription.Contract.IsSubscribed(&_Subscription.CallOpts, _userId)
}

// IsSubscribed is a free data retrieval call binding the contract method 0xee9c32be.
//
// Solidity: function isSubscribed(string _userId) view returns(bool)
func (_Subscription *SubscriptionCallerSession) IsSubscribed(_userId string) (bool, error) {
	return _Subscription.Contract.IsSubscribed(&_Subscription.CallOpts, _userId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subscription *SubscriptionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subscription *SubscriptionSession) Owner() (common.Address, error) {
	return _Subscription.Contract.Owner(&_Subscription.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subscription *SubscriptionCallerSession) Owner() (common.Address, error) {
	return _Subscription.Contract.Owner(&_Subscription.CallOpts)
}

// SubscriptionPrice is a free data retrieval call binding the contract method 0xbdc8e54c.
//
// Solidity: function subscriptionPrice() view returns(uint256)
func (_Subscription *SubscriptionCaller) SubscriptionPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "subscriptionPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubscriptionPrice is a free data retrieval call binding the contract method 0xbdc8e54c.
//
// Solidity: function subscriptionPrice() view returns(uint256)
func (_Subscription *SubscriptionSession) SubscriptionPrice() (*big.Int, error) {
	return _Subscription.Contract.SubscriptionPrice(&_Subscription.CallOpts)
}

// SubscriptionPrice is a free data retrieval call binding the contract method 0xbdc8e54c.
//
// Solidity: function subscriptionPrice() view returns(uint256)
func (_Subscription *SubscriptionCallerSession) SubscriptionPrice() (*big.Int, error) {
	return _Subscription.Contract.SubscriptionPrice(&_Subscription.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscription *SubscriptionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscription *SubscriptionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subscription.Contract.RenounceOwnership(&_Subscription.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscription *SubscriptionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subscription.Contract.RenounceOwnership(&_Subscription.TransactOpts)
}

// SetSubscriptionPrice is a paid mutator transaction binding the contract method 0x1a5ee779.
//
// Solidity: function setSubscriptionPrice(uint256 _newPrice) returns()
func (_Subscription *SubscriptionTransactor) SetSubscriptionPrice(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "setSubscriptionPrice", _newPrice)
}

// SetSubscriptionPrice is a paid mutator transaction binding the contract method 0x1a5ee779.
//
// Solidity: function setSubscriptionPrice(uint256 _newPrice) returns()
func (_Subscription *SubscriptionSession) SetSubscriptionPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SetSubscriptionPrice(&_Subscription.TransactOpts, _newPrice)
}

// SetSubscriptionPrice is a paid mutator transaction binding the contract method 0x1a5ee779.
//
// Solidity: function setSubscriptionPrice(uint256 _newPrice) returns()
func (_Subscription *SubscriptionTransactorSession) SetSubscriptionPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SetSubscriptionPrice(&_Subscription.TransactOpts, _newPrice)
}

// Subscribe is a paid mutator transaction binding the contract method 0x507e7888.
//
// Solidity: function subscribe(string _userId) payable returns()
func (_Subscription *SubscriptionTransactor) Subscribe(opts *bind.TransactOpts, _userId string) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "subscribe", _userId)
}

// Subscribe is a paid mutator transaction binding the contract method 0x507e7888.
//
// Solidity: function subscribe(string _userId) payable returns()
func (_Subscription *SubscriptionSession) Subscribe(_userId string) (*types.Transaction, error) {
	return _Subscription.Contract.Subscribe(&_Subscription.TransactOpts, _userId)
}

// Subscribe is a paid mutator transaction binding the contract method 0x507e7888.
//
// Solidity: function subscribe(string _userId) payable returns()
func (_Subscription *SubscriptionTransactorSession) Subscribe(_userId string) (*types.Transaction, error) {
	return _Subscription.Contract.Subscribe(&_Subscription.TransactOpts, _userId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscription *SubscriptionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscription *SubscriptionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subscription.Contract.TransferOwnership(&_Subscription.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscription *SubscriptionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subscription.Contract.TransferOwnership(&_Subscription.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Subscription *SubscriptionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Subscription *SubscriptionSession) Withdraw() (*types.Transaction, error) {
	return _Subscription.Contract.Withdraw(&_Subscription.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Subscription *SubscriptionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Subscription.Contract.Withdraw(&_Subscription.TransactOpts)
}

// SubscriptionNewSubscriptionIterator is returned from FilterNewSubscription and is used to iterate over the raw logs and unpacked data for NewSubscription events raised by the Subscription contract.
type SubscriptionNewSubscriptionIterator struct {
	Event *SubscriptionNewSubscription // Event containing the contract specifics and raw log

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
func (it *SubscriptionNewSubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionNewSubscription)
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
		it.Event = new(SubscriptionNewSubscription)
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
func (it *SubscriptionNewSubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionNewSubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionNewSubscription represents a NewSubscription event raised by the Subscription contract.
type SubscriptionNewSubscription struct {
	UserId common.Hash
	Expiry *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewSubscription is a free log retrieval operation binding the contract event 0x46c731fe00969f33592dfb439cbe0dd4655a462556648e4cb6d10290ecc9cbfc.
//
// Solidity: event NewSubscription(string indexed userId, uint256 indexed expiry)
func (_Subscription *SubscriptionFilterer) FilterNewSubscription(opts *bind.FilterOpts, userId []string, expiry []*big.Int) (*SubscriptionNewSubscriptionIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var expiryRule []interface{}
	for _, expiryItem := range expiry {
		expiryRule = append(expiryRule, expiryItem)
	}

	logs, sub, err := _Subscription.contract.FilterLogs(opts, "NewSubscription", userIdRule, expiryRule)
	if err != nil {
		return nil, err
	}
	return &SubscriptionNewSubscriptionIterator{contract: _Subscription.contract, event: "NewSubscription", logs: logs, sub: sub}, nil
}

// WatchNewSubscription is a free log subscription operation binding the contract event 0x46c731fe00969f33592dfb439cbe0dd4655a462556648e4cb6d10290ecc9cbfc.
//
// Solidity: event NewSubscription(string indexed userId, uint256 indexed expiry)
func (_Subscription *SubscriptionFilterer) WatchNewSubscription(opts *bind.WatchOpts, sink chan<- *SubscriptionNewSubscription, userId []string, expiry []*big.Int) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var expiryRule []interface{}
	for _, expiryItem := range expiry {
		expiryRule = append(expiryRule, expiryItem)
	}

	logs, sub, err := _Subscription.contract.WatchLogs(opts, "NewSubscription", userIdRule, expiryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionNewSubscription)
				if err := _Subscription.contract.UnpackLog(event, "NewSubscription", log); err != nil {
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

// ParseNewSubscription is a log parse operation binding the contract event 0x46c731fe00969f33592dfb439cbe0dd4655a462556648e4cb6d10290ecc9cbfc.
//
// Solidity: event NewSubscription(string indexed userId, uint256 indexed expiry)
func (_Subscription *SubscriptionFilterer) ParseNewSubscription(log types.Log) (*SubscriptionNewSubscription, error) {
	event := new(SubscriptionNewSubscription)
	if err := _Subscription.contract.UnpackLog(event, "NewSubscription", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscriptionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Subscription contract.
type SubscriptionOwnershipTransferredIterator struct {
	Event *SubscriptionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SubscriptionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionOwnershipTransferred)
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
		it.Event = new(SubscriptionOwnershipTransferred)
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
func (it *SubscriptionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionOwnershipTransferred represents a OwnershipTransferred event raised by the Subscription contract.
type SubscriptionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subscription *SubscriptionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubscriptionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subscription.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubscriptionOwnershipTransferredIterator{contract: _Subscription.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subscription *SubscriptionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubscriptionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subscription.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionOwnershipTransferred)
				if err := _Subscription.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Subscription *SubscriptionFilterer) ParseOwnershipTransferred(log types.Log) (*SubscriptionOwnershipTransferred, error) {
	event := new(SubscriptionOwnershipTransferred)
	if err := _Subscription.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
