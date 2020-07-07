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

// DiplomaABI is the input ABI used to generate the binding from.
const DiplomaABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"student\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diplomaId\",\"type\":\"uint256\"}],\"name\":\"CreateDiploma\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ftPubAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_link\",\"type\":\"string\"}],\"name\":\"Publish42Diploma\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_level\",\"type\":\"uint64\"},{\"internalType\":\"uint64[30]\",\"name\":\"_skills\",\"type\":\"uint64[30]\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_studentHash\",\"type\":\"bytes32\"}],\"name\":\"createDiploma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftPubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_studentHash\",\"type\":\"bytes32\"}],\"name\":\"getDiploma\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"level\",\"type\":\"uint64\"},{\"internalType\":\"uint64[30]\",\"name\":\"skills\",\"type\":\"uint64[30]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashToDiploma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkOfRepo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Diploma is an auto generated Go binding around an Ethereum contract.
type Diploma struct {
	DiplomaCaller     // Read-only binding to the contract
	DiplomaTransactor // Write-only binding to the contract
	DiplomaFilterer   // Log filterer for contract events
}

// DiplomaCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiplomaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiplomaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiplomaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiplomaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiplomaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiplomaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiplomaSession struct {
	Contract     *Diploma          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiplomaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiplomaCallerSession struct {
	Contract *DiplomaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DiplomaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiplomaTransactorSession struct {
	Contract     *DiplomaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiplomaRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiplomaRaw struct {
	Contract *Diploma // Generic contract binding to access the raw methods on
}

// DiplomaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiplomaCallerRaw struct {
	Contract *DiplomaCaller // Generic read-only contract binding to access the raw methods on
}

// DiplomaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiplomaTransactorRaw struct {
	Contract *DiplomaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiploma creates a new instance of Diploma, bound to a specific deployed contract.
func NewDiploma(address common.Address, backend bind.ContractBackend) (*Diploma, error) {
	contract, err := bindDiploma(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Diploma{DiplomaCaller: DiplomaCaller{contract: contract}, DiplomaTransactor: DiplomaTransactor{contract: contract}, DiplomaFilterer: DiplomaFilterer{contract: contract}}, nil
}

// NewDiplomaCaller creates a new read-only instance of Diploma, bound to a specific deployed contract.
func NewDiplomaCaller(address common.Address, caller bind.ContractCaller) (*DiplomaCaller, error) {
	contract, err := bindDiploma(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiplomaCaller{contract: contract}, nil
}

// NewDiplomaTransactor creates a new write-only instance of Diploma, bound to a specific deployed contract.
func NewDiplomaTransactor(address common.Address, transactor bind.ContractTransactor) (*DiplomaTransactor, error) {
	contract, err := bindDiploma(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiplomaTransactor{contract: contract}, nil
}

// NewDiplomaFilterer creates a new log filterer instance of Diploma, bound to a specific deployed contract.
func NewDiplomaFilterer(address common.Address, filterer bind.ContractFilterer) (*DiplomaFilterer, error) {
	contract, err := bindDiploma(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiplomaFilterer{contract: contract}, nil
}

// bindDiploma binds a generic wrapper to an already deployed contract.
func bindDiploma(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiplomaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diploma *DiplomaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Diploma.Contract.DiplomaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diploma *DiplomaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diploma.Contract.DiplomaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diploma *DiplomaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diploma.Contract.DiplomaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diploma *DiplomaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Diploma.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diploma *DiplomaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diploma.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diploma *DiplomaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diploma.Contract.contract.Transact(opts, method, params...)
}

// FtPubAddress is a free data retrieval call binding the contract method 0x5679a859.
//
// Solidity: function ftPubAddress() view returns(address)
func (_Diploma *DiplomaCaller) FtPubAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Diploma.contract.Call(opts, out, "ftPubAddress")
	return *ret0, err
}

// FtPubAddress is a free data retrieval call binding the contract method 0x5679a859.
//
// Solidity: function ftPubAddress() view returns(address)
func (_Diploma *DiplomaSession) FtPubAddress() (common.Address, error) {
	return _Diploma.Contract.FtPubAddress(&_Diploma.CallOpts)
}

// FtPubAddress is a free data retrieval call binding the contract method 0x5679a859.
//
// Solidity: function ftPubAddress() view returns(address)
func (_Diploma *DiplomaCallerSession) FtPubAddress() (common.Address, error) {
	return _Diploma.Contract.FtPubAddress(&_Diploma.CallOpts)
}

// GetDiploma is a free data retrieval call binding the contract method 0x2b04cb77.
//
// Solidity: function getDiploma(bytes32 _studentHash) view returns(uint64 level, uint64[30] skills)
func (_Diploma *DiplomaCaller) GetDiploma(opts *bind.CallOpts, _studentHash [32]byte) (struct {
	Level  uint64
	Skills [30]uint64
}, error) {
	ret := new(struct {
		Level  uint64
		Skills [30]uint64
	})
	out := ret
	err := _Diploma.contract.Call(opts, out, "getDiploma", _studentHash)
	return *ret, err
}

// GetDiploma is a free data retrieval call binding the contract method 0x2b04cb77.
//
// Solidity: function getDiploma(bytes32 _studentHash) view returns(uint64 level, uint64[30] skills)
func (_Diploma *DiplomaSession) GetDiploma(_studentHash [32]byte) (struct {
	Level  uint64
	Skills [30]uint64
}, error) {
	return _Diploma.Contract.GetDiploma(&_Diploma.CallOpts, _studentHash)
}

// GetDiploma is a free data retrieval call binding the contract method 0x2b04cb77.
//
// Solidity: function getDiploma(bytes32 _studentHash) view returns(uint64 level, uint64[30] skills)
func (_Diploma *DiplomaCallerSession) GetDiploma(_studentHash [32]byte) (struct {
	Level  uint64
	Skills [30]uint64
}, error) {
	return _Diploma.Contract.GetDiploma(&_Diploma.CallOpts, _studentHash)
}

// HashToDiploma is a free data retrieval call binding the contract method 0x61262f77.
//
// Solidity: function hashToDiploma(bytes32 ) view returns(uint256)
func (_Diploma *DiplomaCaller) HashToDiploma(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Diploma.contract.Call(opts, out, "hashToDiploma", arg0)
	return *ret0, err
}

// HashToDiploma is a free data retrieval call binding the contract method 0x61262f77.
//
// Solidity: function hashToDiploma(bytes32 ) view returns(uint256)
func (_Diploma *DiplomaSession) HashToDiploma(arg0 [32]byte) (*big.Int, error) {
	return _Diploma.Contract.HashToDiploma(&_Diploma.CallOpts, arg0)
}

// HashToDiploma is a free data retrieval call binding the contract method 0x61262f77.
//
// Solidity: function hashToDiploma(bytes32 ) view returns(uint256)
func (_Diploma *DiplomaCallerSession) HashToDiploma(arg0 [32]byte) (*big.Int, error) {
	return _Diploma.Contract.HashToDiploma(&_Diploma.CallOpts, arg0)
}

// LinkOfRepo is a free data retrieval call binding the contract method 0x8ebc72c1.
//
// Solidity: function linkOfRepo() view returns(string)
func (_Diploma *DiplomaCaller) LinkOfRepo(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Diploma.contract.Call(opts, out, "linkOfRepo")
	return *ret0, err
}

// LinkOfRepo is a free data retrieval call binding the contract method 0x8ebc72c1.
//
// Solidity: function linkOfRepo() view returns(string)
func (_Diploma *DiplomaSession) LinkOfRepo() (string, error) {
	return _Diploma.Contract.LinkOfRepo(&_Diploma.CallOpts)
}

// LinkOfRepo is a free data retrieval call binding the contract method 0x8ebc72c1.
//
// Solidity: function linkOfRepo() view returns(string)
func (_Diploma *DiplomaCallerSession) LinkOfRepo() (string, error) {
	return _Diploma.Contract.LinkOfRepo(&_Diploma.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Diploma *DiplomaCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Diploma.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Diploma *DiplomaSession) Name() (string, error) {
	return _Diploma.Contract.Name(&_Diploma.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Diploma *DiplomaCallerSession) Name() (string, error) {
	return _Diploma.Contract.Name(&_Diploma.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Diploma *DiplomaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Diploma.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Diploma *DiplomaSession) Symbol() (string, error) {
	return _Diploma.Contract.Symbol(&_Diploma.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Diploma *DiplomaCallerSession) Symbol() (string, error) {
	return _Diploma.Contract.Symbol(&_Diploma.CallOpts)
}

// CreateDiploma is a paid mutator transaction binding the contract method 0xfbbb6e40.
//
// Solidity: function createDiploma(uint64 _level, uint64[30] _skills, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) returns()
func (_Diploma *DiplomaTransactor) CreateDiploma(opts *bind.TransactOpts, _level uint64, _skills [30]uint64, _v uint8, _r [32]byte, _s [32]byte, _studentHash [32]byte) (*types.Transaction, error) {
	return _Diploma.contract.Transact(opts, "createDiploma", _level, _skills, _v, _r, _s, _studentHash)
}

// CreateDiploma is a paid mutator transaction binding the contract method 0xfbbb6e40.
//
// Solidity: function createDiploma(uint64 _level, uint64[30] _skills, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) returns()
func (_Diploma *DiplomaSession) CreateDiploma(_level uint64, _skills [30]uint64, _v uint8, _r [32]byte, _s [32]byte, _studentHash [32]byte) (*types.Transaction, error) {
	return _Diploma.Contract.CreateDiploma(&_Diploma.TransactOpts, _level, _skills, _v, _r, _s, _studentHash)
}

// CreateDiploma is a paid mutator transaction binding the contract method 0xfbbb6e40.
//
// Solidity: function createDiploma(uint64 _level, uint64[30] _skills, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) returns()
func (_Diploma *DiplomaTransactorSession) CreateDiploma(_level uint64, _skills [30]uint64, _v uint8, _r [32]byte, _s [32]byte, _studentHash [32]byte) (*types.Transaction, error) {
	return _Diploma.Contract.CreateDiploma(&_Diploma.TransactOpts, _level, _skills, _v, _r, _s, _studentHash)
}

// DiplomaCreateDiplomaIterator is returned from FilterCreateDiploma and is used to iterate over the raw logs and unpacked data for CreateDiploma events raised by the Diploma contract.
type DiplomaCreateDiplomaIterator struct {
	Event *DiplomaCreateDiploma // Event containing the contract specifics and raw log

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
func (it *DiplomaCreateDiplomaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiplomaCreateDiploma)
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
		it.Event = new(DiplomaCreateDiploma)
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
func (it *DiplomaCreateDiplomaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiplomaCreateDiplomaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiplomaCreateDiploma represents a CreateDiploma event raised by the Diploma contract.
type DiplomaCreateDiploma struct {
	Student   [32]byte
	DiplomaId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateDiploma is a free log retrieval operation binding the contract event 0xb0883e564de552c164f168e24578da0ae7e189df9cb6e8634e5d81948f80b0e4.
//
// Solidity: event CreateDiploma(bytes32 student, uint256 diplomaId)
func (_Diploma *DiplomaFilterer) FilterCreateDiploma(opts *bind.FilterOpts) (*DiplomaCreateDiplomaIterator, error) {

	logs, sub, err := _Diploma.contract.FilterLogs(opts, "CreateDiploma")
	if err != nil {
		return nil, err
	}
	return &DiplomaCreateDiplomaIterator{contract: _Diploma.contract, event: "CreateDiploma", logs: logs, sub: sub}, nil
}

// WatchCreateDiploma is a free log subscription operation binding the contract event 0xb0883e564de552c164f168e24578da0ae7e189df9cb6e8634e5d81948f80b0e4.
//
// Solidity: event CreateDiploma(bytes32 student, uint256 diplomaId)
func (_Diploma *DiplomaFilterer) WatchCreateDiploma(opts *bind.WatchOpts, sink chan<- *DiplomaCreateDiploma) (event.Subscription, error) {

	logs, sub, err := _Diploma.contract.WatchLogs(opts, "CreateDiploma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiplomaCreateDiploma)
				if err := _Diploma.contract.UnpackLog(event, "CreateDiploma", log); err != nil {
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

// ParseCreateDiploma is a log parse operation binding the contract event 0xb0883e564de552c164f168e24578da0ae7e189df9cb6e8634e5d81948f80b0e4.
//
// Solidity: event CreateDiploma(bytes32 student, uint256 diplomaId)
func (_Diploma *DiplomaFilterer) ParseCreateDiploma(log types.Log) (*DiplomaCreateDiploma, error) {
	event := new(DiplomaCreateDiploma)
	if err := _Diploma.contract.UnpackLog(event, "CreateDiploma", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DiplomaPublish42DiplomaIterator is returned from FilterPublish42Diploma and is used to iterate over the raw logs and unpacked data for Publish42Diploma events raised by the Diploma contract.
type DiplomaPublish42DiplomaIterator struct {
	Event *DiplomaPublish42Diploma // Event containing the contract specifics and raw log

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
func (it *DiplomaPublish42DiplomaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiplomaPublish42Diploma)
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
		it.Event = new(DiplomaPublish42Diploma)
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
func (it *DiplomaPublish42DiplomaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiplomaPublish42DiplomaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiplomaPublish42Diploma represents a Publish42Diploma event raised by the Diploma contract.
type DiplomaPublish42Diploma struct {
	FtPubAddress common.Address
	Link         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPublish42Diploma is a free log retrieval operation binding the contract event 0x88f959f217ca7512aff2b6020e52875bfe620c8115d803d8d6dec6e0fd952ecc.
//
// Solidity: event Publish42Diploma(address ftPubAddress, string _link)
func (_Diploma *DiplomaFilterer) FilterPublish42Diploma(opts *bind.FilterOpts) (*DiplomaPublish42DiplomaIterator, error) {

	logs, sub, err := _Diploma.contract.FilterLogs(opts, "Publish42Diploma")
	if err != nil {
		return nil, err
	}
	return &DiplomaPublish42DiplomaIterator{contract: _Diploma.contract, event: "Publish42Diploma", logs: logs, sub: sub}, nil
}

// WatchPublish42Diploma is a free log subscription operation binding the contract event 0x88f959f217ca7512aff2b6020e52875bfe620c8115d803d8d6dec6e0fd952ecc.
//
// Solidity: event Publish42Diploma(address ftPubAddress, string _link)
func (_Diploma *DiplomaFilterer) WatchPublish42Diploma(opts *bind.WatchOpts, sink chan<- *DiplomaPublish42Diploma) (event.Subscription, error) {

	logs, sub, err := _Diploma.contract.WatchLogs(opts, "Publish42Diploma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiplomaPublish42Diploma)
				if err := _Diploma.contract.UnpackLog(event, "Publish42Diploma", log); err != nil {
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

// ParsePublish42Diploma is a log parse operation binding the contract event 0x88f959f217ca7512aff2b6020e52875bfe620c8115d803d8d6dec6e0fd952ecc.
//
// Solidity: event Publish42Diploma(address ftPubAddress, string _link)
func (_Diploma *DiplomaFilterer) ParsePublish42Diploma(log types.Log) (*DiplomaPublish42Diploma, error) {
	event := new(DiplomaPublish42Diploma)
	if err := _Diploma.contract.UnpackLog(event, "Publish42Diploma", log); err != nil {
		return nil, err
	}
	return event, nil
}
