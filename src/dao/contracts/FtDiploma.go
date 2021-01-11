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

// FtDiplomaDiploma is an auto generated low-level Go binding around an user-defined struct.
type FtDiplomaDiploma struct {
	Level     uint64
	Skills    [30]uint64
	Hash      [32]byte
	Signature FtDiplomaSign
}

// FtDiplomaSign is an auto generated low-level Go binding around an user-defined struct.
type FtDiplomaSign struct {
	V uint8
	R [32]byte
	S [32]byte
}

// DiplomaABI is the input ABI used to generate the binding from.
const DiplomaABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"student\",\"type\":\"bytes32\"}],\"name\":\"CreateDiploma\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ftPubAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_link\",\"type\":\"string\"}],\"name\":\"Publish42Diploma\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_level\",\"type\":\"uint64\"},{\"internalType\":\"uint64[30]\",\"name\":\"_skills\",\"type\":\"uint64[30]\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_studentHash\",\"type\":\"bytes32\"}],\"name\":\"createDiploma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftPubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllDiploma\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"level\",\"type\":\"uint64\"},{\"internalType\":\"uint64[30]\",\"name\":\"skills\",\"type\":\"uint64[30]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structFtDiploma.Sign\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structFtDiploma.Diploma[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_studentHash\",\"type\":\"bytes32\"}],\"name\":\"getDiploma\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"level\",\"type\":\"uint64\"},{\"internalType\":\"uint64[30]\",\"name\":\"skills\",\"type\":\"uint64[30]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkOfRepo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DiplomaBin is the compiled bytecode used for deploying new contracts.
var DiplomaBin = "0x6080604052348015620000125760006000fd5b505b7f88f959f217ca7512aff2b6020e52875bfe620c8115d803d8d6dec6e0fd952ecc737e12234e994384a757e2689addb2a463ccd3b47d604051806060016040528060268152602001620016756026913960405162000074929190620000dc565b60405180910390a15b620001b056620001af565b62000093816200012f565b82525b5050565b6000620000a78262000111565b620000b381856200011d565b9350620000c581856020860162000165565b620000d0816200019d565b84019150505b92915050565b6000604082019050620000f3600083018562000088565b81810360208301526200010781846200009a565b90505b9392505050565b6000815190505b919050565b60008282526020820190505b92915050565b60006200013c8262000144565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60005b83811015620001865780820151818401525b60208101905062000168565b8381111562000196576000848401525b505b505050565b6000601f19601f83011690505b919050565b5b6114b580620001c06000396000f3fe60806040523480156100115760006000fd5b50600436106100825760003560e01c80638ebc72c11161005c5780638ebc72c1146100f557806395d89b4114610113578063be55a55614610131578063fbbb6e401461014f57610082565b806306fdde03146100885780632b04cb77146100a65780635679a859146100d757610082565b60006000fd5b61009061016b565b60405161009d9190611105565b60405180910390f35b6100c060048036038101906100bb9190610c0c565b6101a7565b6040516100ce92919061118b565b60405180910390f35b6100df610330565b6040516100ec9190611064565b60405180910390f35b6100fd610348565b60405161010a9190611105565b60405180910390f35b61011b610364565b6040516101289190611105565b60405180910390f35b6101396103a0565b6040516101469190611080565b60405180910390f35b61016960048036038101906101649190610c37565b6106cc565b005b6040518060400160405280600981526020017f343220416c756d6e69000000000000000000000000000000000000000000000081526020015081565b60006101b16109ca565b60006000600050600085600019166000191681526020019081526020016000206000506040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201600050601e806020026040519081016040528092919082601e8015610284576020028201916000905b82829054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001906008019060208260070104928301926001038202915080841161023f5790505b50505050508152602001600982016000505460001916600019168152602001600a82016000506040518060600160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182016000505460001916600019168152602001600282016000505460001916600019168152602001505081526020015050905060008160000151905060008260200151905081819450945050505061032b565050505b915091565b737e12234e994384a757e2689addb2a463ccd3b47d81565b60405180606001604052806026815260200161145a6026913981565b6040518060400160405280600381526020017f343241000000000000000000000000000000000000000000000000000000000081526020015081565b6060737e12234e994384a757e2689addb2a463ccd3b47d73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d9061116a565b60405180910390fd5b600060016000508054905067ffffffffffffffff811115610470577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156104a957816020015b6104966109ed565b81526020019060019003908161048e5790505b5090506000600090505b6001600050805490508110156106be5760006000506000600160005083815481101515610509577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020900160005b5054600019166000191681526020019081526020016000206000506040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201600050601e806020026040519081016040528092919082601e80156105e2576020028201916000905b82829054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001906008019060208260070104928301926001038202915080841161059d5790505b50505050508152602001600982016000505460001916600019168152602001600a82016000506040518060600160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182016000505460001916600019168152602001600282016000505460001916600019168152602001505081526020015050828281518110151561069f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052505b80806106b690611349565b9150506104b3565b50809150506106c956505b90565b737e12234e994384a757e2689addb2a463ccd3b47d73ffffffffffffffffffffffffffffffffffffffff166001828686866040516000815260200160405260405161071a94939291906110bf565b6020604051602081039080840390855afa15801561073d573d600060003e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff1614151561079f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079690611128565b60405180910390fd5b600060006000506000836000191660001916815260200190815260200160002060005060000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16141515610826576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081d90611149565b60405180910390fd5b60405180608001604052808767ffffffffffffffff1681526020018681526020018260001916815260200160405180606001604052808760ff16815260200186600019168152602001856000191681526020015081526020015060006000506000836000191660001916815260200190815260200160002060005060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160010160005090601e6108ea929190610a32565b5060408201518160090160005090600019169055606082015181600a0160005060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160010160005090600019169055604082015181600201600050906000191690555050905050600160005081908060018154018082558091505060019003906000526020600020900160005b90919091909150906000191690557f0b638557d56f5cb5f2923ac2d36850a0c1c0a1944c628abb24fd73fabeeee600816040516109b991906110a3565b60405180910390a15b505050505050565b604051806103c00160405280601e90602082028036833780820191505090505090565b6040518060800160405280600067ffffffffffffffff168152602001610a116109ca565b81526020016000600019168152602001610a29610ae3565b81526020015090565b82601e9090600301600490048101928215610ad25791602002820160005b83821115610a9c57835183826101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509260200192600801602081600701049283019260010302610a50565b8015610ad05782816101000a81549067ffffffffffffffff0219169055600801602081600701049283019260010302610a9c565b505b509050610adf9190610b12565b5090565b6040518060600160405280600060ff168152602001600060001916815260200160006000191681526020015090565b610b17565b80821115610b315760008181506000905550600101610b17565b509056611458565b6000610b4c610b47846111e8565b6111b6565b90508082856020860282011115610b635760006000fd5b60005b85811015610b945781610b798882610be0565b8452602084019350602083019250505b600181019050610b66565b5050505b9392505050565b600082601f8301121515610bb35760006000fd5b601e610bc0848285610b39565b9150505b92915050565b600081359050610bd981611407565b5b92915050565b600081359050610bef81611422565b5b92915050565b600081359050610c058161143d565b5b92915050565b600060208284031215610c1f5760006000fd5b6000610c2d84828501610bca565b9150505b92915050565b6000600060006000600060006104608789031215610c555760006000fd5b6000610c6389828a01610be0565b9650506020610c7489828a01610b9f565b9550506103e0610c8689828a01610bf6565b945050610400610c9889828a01610bca565b935050610420610caa89828a01610bca565b925050610440610cbc89828a01610bca565b9150505b9295509295509295565b6000610cd68383610f88565b610460830190505b92915050565b6000610cf08383611024565b6020830190505b92915050565b610d06816112a7565b82525b5050565b6000610d188261122b565b610d22818561126b565b9350610d2d8361120f565b8060005b83811015610d5f578151610d458882610cca565b9750610d508361124f565b9250505b600181019050610d31565b508593505050505b92915050565b610d7681611237565b610d80818461127d565b9250610d8b82611220565b8060005b83811015610dbd578151610da38782610ce4565b9650610dae8361125d565b9250505b600181019050610d8f565b505050505b5050565b610dcf81611237565b610dd98184611289565b9250610de482611220565b8060005b83811015610e16578151610dfc8782610ce4565b9650610e078361125d565b9250505b600181019050610de8565b505050505b5050565b610e28816112ba565b82525b5050565b610e38816112ba565b82525b5050565b6000610e4a82611243565b610e548185611295565b9350610e64818560208601611314565b610e6d816113f5565b84019150505b92915050565b6000610e86602783611295565b91507f46744469706c6f6d613a204973206e6f74203432207369676e2074686973206460008301527f69706c6f6d612e0000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000610eed602683611295565b91507f46744469706c6f6d613a20546865206469706c6f6d6120616c7265616479206560008301527f78697374732e000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000610f54601583611295565b91507f46744469706c6f6d613a204973206e6f742034322e000000000000000000000060008301526020820190505b919050565b61046082016000820151610f9f6000850182611024565b506020820151610fb26020850182610d6d565b506040820151610fc66103e0850182610e1f565b506060820151610fda610400850182610fe1565b50505b5050565b606082016000820151610ff76000850182611044565b50602082015161100a6020850182610e1f565b50604082015161101d6040850182610e1f565b50505b5050565b61102d816112f1565b82525b5050565b61103d816112f1565b82525b5050565b61104d81611306565b82525b5050565b61105d81611306565b82525b5050565b60006020820190506110796000830184610cfd565b5b92915050565b6000602082019050818103600083015261109a8184610d0d565b90505b92915050565b60006020820190506110b86000830184610e2f565b5b92915050565b60006080820190506110d46000830187610e2f565b6110e16020830186611054565b6110ee6040830185610e2f565b6110fb6060830184610e2f565b5b95945050505050565b6000602082019050818103600083015261111f8184610e3f565b90505b92915050565b6000602082019050818103600083015261114181610e79565b90505b919050565b6000602082019050818103600083015261116281610ee0565b90505b919050565b6000602082019050818103600083015261118381610f47565b90505b919050565b60006103e0820190506111a16000830185611034565b6111ae6020830184610dc6565b5b9392505050565b6000604051905081810181811067ffffffffffffffff821117156111dd576111dc6113c4565b5b80604052505b919050565b600067ffffffffffffffff821115611203576112026113c4565b5b6020820290505b919050565b60008190506020820190505b919050565b60008190505b919050565b6000815190505b919050565b6000601e90505b919050565b6000815190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60008282526020820190505b92915050565b60008190505b92915050565b60008190505b92915050565b60008282526020820190505b92915050565b60006112b2826112c5565b90505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600067ffffffffffffffff821690505b919050565b600060ff821690505b919050565b60005b838110156113335780820151818401525b602081019050611317565b83811115611342576000848401525b505b505050565b6000611354826112e6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561138757611386611393565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b611410816112ba565b8114151561141e5760006000fd5b5b50565b61142b816112f1565b811415156114395760006000fd5b5b50565b61144681611306565b811415156114545760006000fd5b5b50565bfe6769746875622e636f6d2f34325363686f6f6c2f626c6f636b636861696e2d73657276696365a2646970667358221220a1bcc9eaab056a73365221e4563cac5e7513fdae69d632bd2985c50bf6d8a15164736f6c634300080000336769746875622e636f6d2f34325363686f6f6c2f626c6f636b636861696e2d73657276696365"

// DeployDiploma deploys a new Ethereum contract, binding an instance of Diploma to it.
func DeployDiploma(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Diploma, error) {
	parsed, err := abi.JSON(strings.NewReader(DiplomaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DiplomaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Diploma{DiplomaCaller: DiplomaCaller{contract: contract}, DiplomaTransactor: DiplomaTransactor{contract: contract}, DiplomaFilterer: DiplomaFilterer{contract: contract}}, nil
}

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

// GetAllDiploma is a free data retrieval call binding the contract method 0xbe55a556.
//
// Solidity: function getAllDiploma() view returns((uint64,uint64[30],bytes32,(uint8,bytes32,bytes32))[])
func (_Diploma *DiplomaCaller) GetAllDiploma(opts *bind.CallOpts) ([]FtDiplomaDiploma, error) {
	var (
		ret0 = new([]FtDiplomaDiploma)
	)
	out := ret0
	err := _Diploma.contract.Call(opts, out, "getAllDiploma")
	return *ret0, err
}

// GetAllDiploma is a free data retrieval call binding the contract method 0xbe55a556.
//
// Solidity: function getAllDiploma() view returns((uint64,uint64[30],bytes32,(uint8,bytes32,bytes32))[])
func (_Diploma *DiplomaSession) GetAllDiploma() ([]FtDiplomaDiploma, error) {
	return _Diploma.Contract.GetAllDiploma(&_Diploma.CallOpts)
}

// GetAllDiploma is a free data retrieval call binding the contract method 0xbe55a556.
//
// Solidity: function getAllDiploma() view returns((uint64,uint64[30],bytes32,(uint8,bytes32,bytes32))[])
func (_Diploma *DiplomaCallerSession) GetAllDiploma() ([]FtDiplomaDiploma, error) {
	return _Diploma.Contract.GetAllDiploma(&_Diploma.CallOpts)
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
	Student [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateDiploma is a free log retrieval operation binding the contract event 0x0b638557d56f5cb5f2923ac2d36850a0c1c0a1944c628abb24fd73fabeeee600.
//
// Solidity: event CreateDiploma(bytes32 student)
func (_Diploma *DiplomaFilterer) FilterCreateDiploma(opts *bind.FilterOpts) (*DiplomaCreateDiplomaIterator, error) {

	logs, sub, err := _Diploma.contract.FilterLogs(opts, "CreateDiploma")
	if err != nil {
		return nil, err
	}
	return &DiplomaCreateDiplomaIterator{contract: _Diploma.contract, event: "CreateDiploma", logs: logs, sub: sub}, nil
}

// WatchCreateDiploma is a free log subscription operation binding the contract event 0x0b638557d56f5cb5f2923ac2d36850a0c1c0a1944c628abb24fd73fabeeee600.
//
// Solidity: event CreateDiploma(bytes32 student)
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

// ParseCreateDiploma is a log parse operation binding the contract event 0x0b638557d56f5cb5f2923ac2d36850a0c1c0a1944c628abb24fd73fabeeee600.
//
// Solidity: event CreateDiploma(bytes32 student)
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
