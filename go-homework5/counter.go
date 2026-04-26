// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// SimpleCounterMetaData contains all meta data concerning the SimpleCounter contract.
var SimpleCounterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"setCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104038061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80632baeceb71461004e578063a87d942c14610058578063d09de08a14610076578063d14e62b814610080575b5f5ffd5b61005661009c565b005b610060610133565b60405161006d91906101ea565b60405180910390f35b61007e61013b565b005b61009a60048036038101906100959190610231565b61018f565b005b5f5f54116100df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d6906102b6565b60405180910390fd5b60015f5f8282546100f09190610301565b925050819055507fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f5433604051610129929190610373565b60405180910390a1565b5f5f54905090565b60015f5f82825461014c919061039a565b925050819055507fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f5433604051610185929190610373565b60405180910390a1565b805f819055507fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f54336040516101c7929190610373565b60405180910390a150565b5f819050919050565b6101e4816101d2565b82525050565b5f6020820190506101fd5f8301846101db565b92915050565b5f5ffd5b610210816101d2565b811461021a575f5ffd5b50565b5f8135905061022b81610207565b92915050565b5f6020828403121561024657610245610203565b5b5f6102538482850161021d565b91505092915050565b5f82825260208201905092915050565b7f436f756e742063616e6e6f74206265206e6567617469766500000000000000005f82015250565b5f6102a060188361025c565b91506102ab8261026c565b602082019050919050565b5f6020820190508181035f8301526102cd81610294565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61030b826101d2565b9150610316836101d2565b925082820390508181111561032e5761032d6102d4565b5b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61035d82610334565b9050919050565b61036d81610353565b82525050565b5f6040820190506103865f8301856101db565b6103936020830184610364565b9392505050565b5f6103a4826101d2565b91506103af836101d2565b92508282019050808211156103c7576103c66102d4565b5b9291505056fea264697066735822122035d0aca782bcc0619ac810187907abc91e477ce45362266a182a6dfa8add152364736f6c63430008220033",
}

// SimpleCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleCounterMetaData.ABI instead.
var SimpleCounterABI = SimpleCounterMetaData.ABI

// SimpleCounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleCounterMetaData.Bin instead.
var SimpleCounterBin = SimpleCounterMetaData.Bin

// DeploySimpleCounter deploys a new Ethereum contract, binding an instance of SimpleCounter to it.
func DeploySimpleCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleCounter, error) {
	parsed, err := SimpleCounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleCounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleCounter{SimpleCounterCaller: SimpleCounterCaller{contract: contract}, SimpleCounterTransactor: SimpleCounterTransactor{contract: contract}, SimpleCounterFilterer: SimpleCounterFilterer{contract: contract}}, nil
}

// SimpleCounter is an auto generated Go binding around an Ethereum contract.
type SimpleCounter struct {
	SimpleCounterCaller     // Read-only binding to the contract
	SimpleCounterTransactor // Write-only binding to the contract
	SimpleCounterFilterer   // Log filterer for contract events
}

// SimpleCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleCounterSession struct {
	Contract     *SimpleCounter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCounterCallerSession struct {
	Contract *SimpleCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleCounterTransactorSession struct {
	Contract     *SimpleCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleCounterRaw struct {
	Contract *SimpleCounter // Generic contract binding to access the raw methods on
}

// SimpleCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCounterCallerRaw struct {
	Contract *SimpleCounterCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleCounterTransactorRaw struct {
	Contract *SimpleCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleCounter creates a new instance of SimpleCounter, bound to a specific deployed contract.
func NewSimpleCounter(address common.Address, backend bind.ContractBackend) (*SimpleCounter, error) {
	contract, err := bindSimpleCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleCounter{SimpleCounterCaller: SimpleCounterCaller{contract: contract}, SimpleCounterTransactor: SimpleCounterTransactor{contract: contract}, SimpleCounterFilterer: SimpleCounterFilterer{contract: contract}}, nil
}

// NewSimpleCounterCaller creates a new read-only instance of SimpleCounter, bound to a specific deployed contract.
func NewSimpleCounterCaller(address common.Address, caller bind.ContractCaller) (*SimpleCounterCaller, error) {
	contract, err := bindSimpleCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCounterCaller{contract: contract}, nil
}

// NewSimpleCounterTransactor creates a new write-only instance of SimpleCounter, bound to a specific deployed contract.
func NewSimpleCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleCounterTransactor, error) {
	contract, err := bindSimpleCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCounterTransactor{contract: contract}, nil
}

// NewSimpleCounterFilterer creates a new log filterer instance of SimpleCounter, bound to a specific deployed contract.
func NewSimpleCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleCounterFilterer, error) {
	contract, err := bindSimpleCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleCounterFilterer{contract: contract}, nil
}

// bindSimpleCounter binds a generic wrapper to an already deployed contract.
func bindSimpleCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCounter *SimpleCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCounter.Contract.SimpleCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCounter *SimpleCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCounter.Contract.SimpleCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCounter *SimpleCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCounter.Contract.SimpleCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCounter *SimpleCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCounter *SimpleCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCounter *SimpleCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCounter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_SimpleCounter *SimpleCounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleCounter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_SimpleCounter *SimpleCounterSession) GetCount() (*big.Int, error) {
	return _SimpleCounter.Contract.GetCount(&_SimpleCounter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_SimpleCounter *SimpleCounterCallerSession) GetCount() (*big.Int, error) {
	return _SimpleCounter.Contract.GetCount(&_SimpleCounter.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_SimpleCounter *SimpleCounterTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCounter.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_SimpleCounter *SimpleCounterSession) Decrement() (*types.Transaction, error) {
	return _SimpleCounter.Contract.Decrement(&_SimpleCounter.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_SimpleCounter *SimpleCounterTransactorSession) Decrement() (*types.Transaction, error) {
	return _SimpleCounter.Contract.Decrement(&_SimpleCounter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_SimpleCounter *SimpleCounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCounter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_SimpleCounter *SimpleCounterSession) Increment() (*types.Transaction, error) {
	return _SimpleCounter.Contract.Increment(&_SimpleCounter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_SimpleCounter *SimpleCounterTransactorSession) Increment() (*types.Transaction, error) {
	return _SimpleCounter.Contract.Increment(&_SimpleCounter.TransactOpts)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_SimpleCounter *SimpleCounterTransactor) SetCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _SimpleCounter.contract.Transact(opts, "setCount", _count)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_SimpleCounter *SimpleCounterSession) SetCount(_count *big.Int) (*types.Transaction, error) {
	return _SimpleCounter.Contract.SetCount(&_SimpleCounter.TransactOpts, _count)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_SimpleCounter *SimpleCounterTransactorSession) SetCount(_count *big.Int) (*types.Transaction, error) {
	return _SimpleCounter.Contract.SetCount(&_SimpleCounter.TransactOpts, _count)
}

// SimpleCounterCountChangedIterator is returned from FilterCountChanged and is used to iterate over the raw logs and unpacked data for CountChanged events raised by the SimpleCounter contract.
type SimpleCounterCountChangedIterator struct {
	Event *SimpleCounterCountChanged // Event containing the contract specifics and raw log

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
func (it *SimpleCounterCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCounterCountChanged)
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
		it.Event = new(SimpleCounterCountChanged)
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
func (it *SimpleCounterCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCounterCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCounterCountChanged represents a CountChanged event raised by the SimpleCounter contract.
type SimpleCounterCountChanged struct {
	NewCount  *big.Int
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCountChanged is a free log retrieval operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_SimpleCounter *SimpleCounterFilterer) FilterCountChanged(opts *bind.FilterOpts) (*SimpleCounterCountChangedIterator, error) {

	logs, sub, err := _SimpleCounter.contract.FilterLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return &SimpleCounterCountChangedIterator{contract: _SimpleCounter.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged is a free log subscription operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_SimpleCounter *SimpleCounterFilterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *SimpleCounterCountChanged) (event.Subscription, error) {

	logs, sub, err := _SimpleCounter.contract.WatchLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCounterCountChanged)
				if err := _SimpleCounter.contract.UnpackLog(event, "CountChanged", log); err != nil {
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

// ParseCountChanged is a log parse operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_SimpleCounter *SimpleCounterFilterer) ParseCountChanged(log types.Log) (*SimpleCounterCountChanged, error) {
	event := new(SimpleCounterCountChanged)
	if err := _SimpleCounter.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
