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

// FtDiplomaDiplomas is an auto generated low-level Go binding around an user-defined struct.
type FtDiplomaDiplomas struct {
	Level     uint64
	Skills    [30]FtDiplomaSkill
	Hash      [32]byte
	Signature FtDiplomaSign
}

// FtDiplomaSign is an auto generated low-level Go binding around an user-defined struct.
type FtDiplomaSign struct {
	V uint8
	R [32]byte
	S [32]byte
}

// FtDiplomaSkill is an auto generated low-level Go binding around an user-defined struct.
type FtDiplomaSkill struct {
	Slug  string
	Level uint64
}

// DiplomaABI is the input ABI used to generate the binding from.
const DiplomaABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"student\",\"type\":\"bytes32\"}],\"name\":\"CreateDiploma\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ftPubAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_link\",\"type\":\"string\"}],\"name\":\"Publish42Diploma\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_level\",\"type\":\"uint64\"},{\"internalType\":\"uint64[30]\",\"name\":\"_skillLevel\",\"type\":\"uint64[30]\"},{\"internalType\":\"string[30]\",\"name\":\"_skillSlug\",\"type\":\"string[30]\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_studentHash\",\"type\":\"bytes32\"}],\"name\":\"createDiploma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftPubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllDiploma\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"level\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"slug\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"level\",\"type\":\"uint64\"}],\"internalType\":\"structFtDiploma.Skill[30]\",\"name\":\"skills\",\"type\":\"tuple[30]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structFtDiploma.Sign\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structFtDiploma.Diplomas[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_studentHash\",\"type\":\"bytes32\"}],\"name\":\"getDiploma\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"slug\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"level\",\"type\":\"uint64\"}],\"internalType\":\"structFtDiploma.Skill[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkOfRepo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DiplomaBin is the compiled bytecode used for deploying new contracts.
var DiplomaBin = "0x6080604052348015620000125760006000fd5b505b7f88f959f217ca7512aff2b6020e52875bfe620c8115d803d8d6dec6e0fd952ecc737e12234e994384a757e2689addb2a463ccd3b47d60405180606001604052806026815260200162001d726026913960405162000074929190620000dc565b60405180910390a15b620001b056620001af565b62000093816200012f565b82525b5050565b6000620000a78262000111565b620000b381856200011d565b9350620000c581856020860162000165565b620000d0816200019d565b84019150505b92915050565b6000604082019050620000f3600083018562000088565b81810360208301526200010781846200009a565b90505b9392505050565b6000815190505b919050565b60008282526020820190505b92915050565b60006200013c8262000144565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60005b83811015620001865780820151818401525b60208101905062000168565b8381111562000196576000848401525b505b505050565b6000601f19601f83011690505b919050565b5b611bb280620001c06000396000f3fe60806040523480156100115760006000fd5b50600436106100825760003560e01c80635679a8591161005c5780635679a859146100f35780638ebc72c11461011157806395d89b411461012f578063be55a5561461014d57610082565b806306fdde03146100885780632b04cb77146100a65780633a59ea31146100d757610082565b60006000fd5b61009061016b565b60405161009d91906116b9565b60405180910390f35b6100c060048036038101906100bb91906110c4565b6101a7565b6040516100ce92919061173f565b60405180910390f35b6100f160048036038101906100ec91906110ef565b610405565b005b6100fb610815565b6040516101089190611618565b60405180910390f35b61011961082d565b60405161012691906116b9565b60405180910390f35b610137610849565b60405161014491906116b9565b60405180910390f35b610155610885565b6040516101629190611634565b60405180910390f35b6040518060400160405280600981526020017f343220416c756d6e69000000000000000000000000000000000000000000000081526020015081565b600060606000600060005060008560001916600019168152602001908152602001600020600050905060008160000160009054906101000a900467ffffffffffffffff1690506000601e67ffffffffffffffff811115610230577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561026957816020015b610256610d87565b81526020019060019003908161024e5790505b5090506000600090505b601e8167ffffffffffffffff1610156103ee578360010160005060008267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000506040518060400160405290816000820160005080546102d3906119ae565b80601f01602080910402602001604051908101604052809291908181526020018280546102ff906119ae565b801561034c5780601f106103215761010080835404028352916020019161034c565b820191906000526020600020905b81548152906001019060200180831161032f57829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015050828267ffffffffffffffff168151811015156103cf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052505b80806103e690611a2d565b915050610273565b50818194509450505050610400565050505b915091565b737e12234e994384a757e2689addb2a463ccd3b47d73ffffffffffffffffffffffffffffffffffffffff16600182868686604051600081526020016040526040516104539493929190611673565b6020604051602081039080840390855afa158015610476573d600060003e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff161415156104d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104cf906116dc565b60405180910390fd5b600060006000506000836000191660001916815260200190815260200160002060005060000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614151561055f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610556906116fd565b60405180910390fd5b60006000600050600083600019166000191681526020019081526020016000206000509050878160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555081816002016000508190906000191690555060405180606001604052808660ff1681526020018560001916815260200184600019168152602001508160030160005060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160010160005090600019169055604082015181600201600050906000191690559050506000600090505b601e8167ffffffffffffffff16101561079e576040518060400160405280888367ffffffffffffffff16601e811015156106a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201518152602001898367ffffffffffffffff16601e811015156106f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015167ffffffffffffffff168152602001508260010160005060008367ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000506000820151816000016000509080519060200190610757929190610dae565b5060208201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050505b808061079690611a2d565b915050610642565b50600160005082908060018154018082558091505060019003906000526020600020900160005b90919091909150906000191690557f0b638557d56f5cb5f2923ac2d36850a0c1c0a1944c628abb24fd73fabeeee600826040516108029190611657565b60405180910390a1505b50505050505050565b737e12234e994384a757e2689addb2a463ccd3b47d81565b604051806060016040528060268152602001611b576026913981565b6040518060400160405280600381526020017f343241000000000000000000000000000000000000000000000000000000000081526020015081565b6060737e12234e994384a757e2689addb2a463ccd3b47d73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561090b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109029061171e565b60405180910390fd5b600060016000508054905067ffffffffffffffff811115610955577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561098e57816020015b61097b610e39565b8152602001906001900390816109735790505b5090506000600090505b600160005080549050811015610d79576000600060005060006001600050848154811015156109f0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020900160005b50546000191660001916815260200190815260200160002060005090508060000160009054906101000a900467ffffffffffffffff168383815181101515610a6f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000019067ffffffffffffffff16908167ffffffffffffffff16815260200150506000600090505b601e8167ffffffffffffffff161015610c5d578160010160005060008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600050604051806040016040529081600082016000508054610aff906119ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2b906119ae565b8015610b785780601f10610b4d57610100808354040283529160200191610b78565b820191906000526020600020905b815481529060010190602001808311610b5b57829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200150508484815181101515610bf1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001518267ffffffffffffffff16601e81101515610c41577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201819052505b8080610c5590611a2d565b915050610a9f565b5080600201600050548383815181101515610ca1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151604001906000191690816000191681526020015050806003016000506040518060600160405290816000820160009054906101000a900460ff1660ff1660ff16815260200160018201600050546000191660001916815260200160028201600050546000191660001916815260200150508383815181101515610d55577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160600181905250505b8080610d71906119e3565b915050610998565b5080915050610d8456505b90565b604051806040016040528060608152602001600067ffffffffffffffff1681526020015090565b828054610dba906119ae565b90600052602060002090601f016020900481019282610ddc5760008555610e28565b82601f10610df557805160ff1916838001178555610e28565b82800160010185558215610e28579182015b82811115610e275782518260005090905591602001919060010190610e07565b5b509050610e359190610e7e565b5090565b6040518060800160405280600067ffffffffffffffff168152602001610e5d610ea1565b81526020016000600019168152602001610e75610ecf565b81526020015090565b610e83565b80821115610e9d5760008181506000905550600101610e83565b5090565b604051806103c00160405280601e905b610eb9610d87565b815260200190600190039081610eb15790505090565b6040518060600160405280600060ff16815260200160006000191681526020016000600019168152602001509056611b55565b6000610f15610f10846117a2565b611770565b9050808260005b85811015610f4d5781358501610f32888261106a565b8452602084019350602083019250505b600181019050610f1c565b5050505b9392505050565b6000610f6b610f66846117c9565b611770565b90508082856020860282011115610f825760006000fd5b60005b85811015610fb35781610f988882611098565b8452602084019350602083019250505b600181019050610f85565b5050505b9392505050565b6000610fd1610fcc846117f0565b611770565b905082815260208101848484011115610fea5760006000fd5b610ff5848285611969565b505b9392505050565b600082601f83011215156110125760006000fd5b601e61101f848285610f02565b9150505b92915050565b600082601f830112151561103d5760006000fd5b601e61104a848285610f58565b9150505b92915050565b60008135905061106381611b04565b5b92915050565b600082601f830112151561107e5760006000fd5b813561108e848260208601610fbe565b9150505b92915050565b6000813590506110a781611b1f565b5b92915050565b6000813590506110bd81611b3a565b5b92915050565b6000602082840312156110d75760006000fd5b60006110e584828501611054565b9150505b92915050565b6000600060006000600060006000610480888a03121561110f5760006000fd5b600061111d8a828b01611098565b975050602061112e8a828b01611029565b9650506103e088013567ffffffffffffffff81111561114d5760006000fd5b6111598a828b01610ffe565b95505061040061116b8a828b016110ae565b94505061042061117d8a828b01611054565b93505061044061118f8a828b01611054565b9250506104606111a18a828b01611054565b9150505b92959891949750929550565b60006111bd83836114f3565b90505b92915050565b60006111d2838361159a565b90505b92915050565b6111e4816118fc565b82525b5050565b60006111f68261184e565b61120081856118a8565b93508360208202850161121285611821565b8060005b8581101561124f578484038952815161122f85826111b1565b945061123a8361187e565b925060208a019950505b600181019050611216565b5082975087955050505050505b92915050565b600061126d8261185a565b61127781856118ba565b93508360208202850161128985611832565b8060005b858110156112c657848403895281516112a685826111c6565b94506112b18361188c565b925060208a019950505b60018101905061128d565b5082975087955050505050505b92915050565b60006112e482611866565b6112ee81856118c6565b9350836020820285016113008561183d565b8060005b8581101561133d578484038952815161131d85826111c6565b94506113288361189a565b925060208a019950505b600181019050611304565b5082975087955050505050505b92915050565b6113598161190f565b82525b5050565b6113698161190f565b82525b5050565b600061137b82611872565b61138581856118d8565b9350611395818560208601611979565b61139e81611af2565b84019150505b92915050565b60006113b582611872565b6113bf81856118ea565b93506113cf818560208601611979565b6113d881611af2565b84019150505b92915050565b60006113f16027836118ea565b91507f46744469706c6f6d613a204973206e6f74203432207369676e2074686973206460008301527f69706c6f6d612e0000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006114586026836118ea565b91507f46744469706c6f6d613a20546865206469706c6f6d6120616c7265616479206560008301527f78697374732e000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006114bf6015836118ea565b91507f46744469706c6f6d613a204973206e6f742034322e000000000000000000000060008301526020820190505b919050565b600060c08301600083015161150b60008601826115d8565b50602083015184820360208601526115238282611262565b91505060408301516115386040860182611350565b50606083015161154b6060860182611557565b50809150505b92915050565b60608201600082015161156d60008501826115f8565b5060208201516115806020850182611350565b5060408201516115936040850182611350565b50505b5050565b600060408301600083015184820360008601526115b78282611370565b91505060208301516115cc60208601826115d8565b50809150505b92915050565b6115e181611946565b82525b5050565b6115f181611946565b82525b5050565b6116018161195b565b82525b5050565b6116118161195b565b82525b5050565b600060208201905061162d60008301846111db565b5b92915050565b6000602082019050818103600083015261164e81846111eb565b90505b92915050565b600060208201905061166c6000830184611360565b5b92915050565b60006080820190506116886000830187611360565b6116956020830186611608565b6116a26040830185611360565b6116af6060830184611360565b5b95945050505050565b600060208201905081810360008301526116d381846113aa565b90505b92915050565b600060208201905081810360008301526116f5816113e4565b90505b919050565b600060208201905081810360008301526117168161144b565b90505b919050565b60006020820190508181036000830152611737816114b2565b90505b919050565b600060408201905061175460008301856115e8565b818103602083015261176681846112d9565b90505b9392505050565b6000604051905081810181811067ffffffffffffffff8211171561179757611796611ac1565b5b80604052505b919050565b600067ffffffffffffffff8211156117bd576117bc611ac1565b5b6020820290505b919050565b600067ffffffffffffffff8211156117e4576117e3611ac1565b5b6020820290505b919050565b600067ffffffffffffffff82111561180b5761180a611ac1565b5b601f19601f83011690506020810190505b919050565b60008190506020820190505b919050565b60008190505b919050565b60008190506020820190505b919050565b6000815190505b919050565b6000601e90505b919050565b6000815190505b919050565b6000815190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60008282526020820190505b92915050565b60008190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60006119078261191a565b90505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600067ffffffffffffffff821690505b919050565b600060ff821690505b919050565b828183376000838301525b505050565b60005b838110156119985780820151818401525b60208101905061197c565b838111156119a7576000848401525b505b505050565b6000600282049050600182168015156119c857607f821691505b602082108114156119dc576119db611a90565b5b505b919050565b60006119ee8261193b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611a2157611a20611a5f565b5b6001820190505b919050565b6000611a3882611946565b915067ffffffffffffffff821415611a5357611a52611a5f565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b611b0d8161190f565b81141515611b1b5760006000fd5b5b50565b611b2881611946565b81141515611b365760006000fd5b5b50565b611b438161195b565b81141515611b515760006000fd5b5b50565bfe6769746875622e636f6d2f34325363686f6f6c2f626c6f636b636861696e2d73657276696365a2646970667358221220b5ae8c44410b03e1083b28ba667995a3b45cc4bb0cd2d0ba5168018b3e581a8264736f6c634300080000336769746875622e636f6d2f34325363686f6f6c2f626c6f636b636861696e2d73657276696365"

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
func (_Diploma *DiplomaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Diploma *DiplomaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _Diploma.contract.Call(opts, &out, "ftPubAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
// Solidity: function getAllDiploma() view returns((uint64,(string,uint64)[30],bytes32,(uint8,bytes32,bytes32))[])
func (_Diploma *DiplomaCaller) GetAllDiploma(opts *bind.CallOpts) ([]FtDiplomaDiplomas, error) {
	var out []interface{}
	err := _Diploma.contract.Call(opts, &out, "getAllDiploma")

	if err != nil {
		return *new([]FtDiplomaDiplomas), err
	}

	out0 := *abi.ConvertType(out[0], new([]FtDiplomaDiplomas)).(*[]FtDiplomaDiplomas)

	return out0, err

}

// GetAllDiploma is a free data retrieval call binding the contract method 0xbe55a556.
//
// Solidity: function getAllDiploma() view returns((uint64,(string,uint64)[30],bytes32,(uint8,bytes32,bytes32))[])
func (_Diploma *DiplomaSession) GetAllDiploma() ([]FtDiplomaDiplomas, error) {
	return _Diploma.Contract.GetAllDiploma(&_Diploma.CallOpts)
}

// GetAllDiploma is a free data retrieval call binding the contract method 0xbe55a556.
//
// Solidity: function getAllDiploma() view returns((uint64,(string,uint64)[30],bytes32,(uint8,bytes32,bytes32))[])
func (_Diploma *DiplomaCallerSession) GetAllDiploma() ([]FtDiplomaDiplomas, error) {
	return _Diploma.Contract.GetAllDiploma(&_Diploma.CallOpts)
}

// GetDiploma is a free data retrieval call binding the contract method 0x2b04cb77.
//
// Solidity: function getDiploma(bytes32 _studentHash) view returns(uint64, (string,uint64)[])
func (_Diploma *DiplomaCaller) GetDiploma(opts *bind.CallOpts, _studentHash [32]byte) (uint64, []FtDiplomaSkill, error) {
	var out []interface{}
	err := _Diploma.contract.Call(opts, &out, "getDiploma", _studentHash)

	if err != nil {
		return *new(uint64), *new([]FtDiplomaSkill), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new([]FtDiplomaSkill)).(*[]FtDiplomaSkill)

	return out0, out1, err

}

// GetDiploma is a free data retrieval call binding the contract method 0x2b04cb77.
//
// Solidity: function getDiploma(bytes32 _studentHash) view returns(uint64, (string,uint64)[])
func (_Diploma *DiplomaSession) GetDiploma(_studentHash [32]byte) (uint64, []FtDiplomaSkill, error) {
	return _Diploma.Contract.GetDiploma(&_Diploma.CallOpts, _studentHash)
}

// GetDiploma is a free data retrieval call binding the contract method 0x2b04cb77.
//
// Solidity: function getDiploma(bytes32 _studentHash) view returns(uint64, (string,uint64)[])
func (_Diploma *DiplomaCallerSession) GetDiploma(_studentHash [32]byte) (uint64, []FtDiplomaSkill, error) {
	return _Diploma.Contract.GetDiploma(&_Diploma.CallOpts, _studentHash)
}

// LinkOfRepo is a free data retrieval call binding the contract method 0x8ebc72c1.
//
// Solidity: function linkOfRepo() view returns(string)
func (_Diploma *DiplomaCaller) LinkOfRepo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Diploma.contract.Call(opts, &out, "linkOfRepo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

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
	var out []interface{}
	err := _Diploma.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

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
	var out []interface{}
	err := _Diploma.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

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

// CreateDiploma is a paid mutator transaction binding the contract method 0x3a59ea31.
//
// Solidity: function createDiploma(uint64 _level, uint64[30] _skillLevel, string[30] _skillSlug, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) returns()
func (_Diploma *DiplomaTransactor) CreateDiploma(opts *bind.TransactOpts, _level uint64, _skillLevel [30]uint64, _skillSlug [30]string, _v uint8, _r [32]byte, _s [32]byte, _studentHash [32]byte) (*types.Transaction, error) {
	return _Diploma.contract.Transact(opts, "createDiploma", _level, _skillLevel, _skillSlug, _v, _r, _s, _studentHash)
}

// CreateDiploma is a paid mutator transaction binding the contract method 0x3a59ea31.
//
// Solidity: function createDiploma(uint64 _level, uint64[30] _skillLevel, string[30] _skillSlug, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) returns()
func (_Diploma *DiplomaSession) CreateDiploma(_level uint64, _skillLevel [30]uint64, _skillSlug [30]string, _v uint8, _r [32]byte, _s [32]byte, _studentHash [32]byte) (*types.Transaction, error) {
	return _Diploma.Contract.CreateDiploma(&_Diploma.TransactOpts, _level, _skillLevel, _skillSlug, _v, _r, _s, _studentHash)
}

// CreateDiploma is a paid mutator transaction binding the contract method 0x3a59ea31.
//
// Solidity: function createDiploma(uint64 _level, uint64[30] _skillLevel, string[30] _skillSlug, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) returns()
func (_Diploma *DiplomaTransactorSession) CreateDiploma(_level uint64, _skillLevel [30]uint64, _skillSlug [30]string, _v uint8, _r [32]byte, _s [32]byte, _studentHash [32]byte) (*types.Transaction, error) {
	return _Diploma.Contract.CreateDiploma(&_Diploma.TransactOpts, _level, _skillLevel, _skillSlug, _v, _r, _s, _studentHash)
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}
