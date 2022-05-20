package contracts

import (
	"bytes"
	"context"
	"fmt"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"strings"

	"math/big"
)

var Blockchain BlockchainFunc

type BlockchainImpl struct {}

// connectEthGetInstance connect the client Ethereum and get the instance of the smart-contract.
func (bc BlockchainImpl) connectEthGetInstance() (*Diploma, *ethclient.Client, error) {
	client, err := ethclient.Dial(tools.NetworkLink)
	if err != nil {
		return nil, nil, err
	}
	addressOfAddress := common.HexToAddress(tools.AddressOfContract)
	instance, err := NewDiploma(addressOfAddress, client)
	if err != nil {
		return nil, nil, err
	}
	return instance, client, nil
}

func (bc BlockchainImpl) getAuth() (*bind.TransactOpts, error) {
	//ethclient.NewClient(rpc.Dial(rpc.DialIPC()))
	client, err := ethclient.Dial(tools.NetworkLink)
	if err != nil {
		log.Debug("Dial")
		return nil, err
	}
	address, privateKey, err := account.Accounts.GetWriter()
	if err != nil {
		log.Debug("GetWriterAccount")
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Debug("ChainID")
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Debug("PendingNonceAt")
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Debug("SuggestGasPrice")
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Debug("NewKeyedTransactorWithChainID")
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(2424242)
	auth.GasPrice = gasPrice
	return auth, nil
}

func (bc BlockchainImpl) getLogs(client *ethclient.Client) ([]types.Log, abi.ABI, error) {
	query := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(tools.AddressOfContract)}}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		tools.LogsError(err)
		return nil, abi.ABI{}, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(DiplomaABI))
	if err != nil {
		tools.LogsError(err)
		return nil, abi.ABI{}, err
	}
	return logs, contractAbi, nil
}

func (bc BlockchainImpl) GetBalance(address common.Address) (int64, error) {
	client, err := ethclient.Dial(tools.NetworkLink)
	if err != nil {
		return 0, err
	}
	balance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		return 0, err
	}
	return balance.Int64(), err
}

func (bc BlockchainImpl) GetRevert(client *ethclient.Client, tx *types.Transaction, receipt *types.Receipt) string {
	address, _, _ := account.Accounts.GetWriter()
	msg := ethereum.CallMsg{
		From:     address,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	res, err := client.CallContract(context.Background(), msg, receipt.BlockNumber)
	if err != nil {
		tools.LogsError(err)
		return ""
	}
	var (
		errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0}
		abiString, _ = abi.NewType("string", "", nil)
	)
	if len(res) < 4 || !bytes.Equal(res[:4], errorSig) {
		tools.LogsError(err)
		return ""
	}
	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(res[4:])
	if err != nil {
		tools.LogsError(err)
		return ""
	}
	return vs[0].(string)
}

func (bc BlockchainImpl) CheckSecurity(client *ethclient.Client, tx *types.Transaction, hash []byte) bool {
	logs, contractAbi, err := bc.getLogs(client)
	if err != nil {
		log.Println(err)
		return true
	}
	for _, vLog := range logs {
		if vLog.TxHash.Hex() == tx.Hash().Hex() {
			var eventHash [32]byte
			eventData, err := contractAbi.Unpack("CreateDiploma", vLog.Data)
			eventHash = eventData[0].([32]byte)
			if err != nil {
				tools.LogsError(err)
				return true
			}
			if common.Bytes2Hex(hash[:]) != common.Bytes2Hex(eventHash[:]) {
				log.Panic("Error: The hash writing in blockchain is not the same of this student !")
				return false
			}
		}
	}
	return true
}

func (bc BlockchainImpl) CallCreateDiploma(level uint64, skills [30]uint64, skillsSlugs [30]string, v uint8, r [32]byte, s [32]byte, hash [32]byte) (*types.Transaction, bool) {
	instance, _, err := bc.connectEthGetInstance()
	if err != nil {
		log.Debug("Instance")
		tools.LogsError(err)
		return nil, false
	}
	auth, err := bc.getAuth()
	if err != nil {
		log.Debug("Auth")
		tools.LogsError(err)
		return nil, false
	}
	tx, err := instance.CreateDiploma(auth, level, skills, skillsSlugs, v, r, s, hash)
	if err != nil {
		log.Debug("Write")
		tools.LogsError(err)
		if strings.Contains(err.Error(), "insufficient funds for gas * price + value") {
			account.Accounts.ChangeWriter()
		}
		return nil, false
	}
	log.WithFields(log.Fields{"tx_hash": tx.Hash().Hex()}).Debug("Transaction Hash")
	return tx, true
}

func (bc BlockchainImpl) CallGetDiploma(hash []byte) (uint64, []FtDiplomaSkill, error) {
	instance, _, err := bc.connectEthGetInstance()
	if err != nil {
		return 0, []FtDiplomaSkill{}, err
	}
	hash32 := [32]byte{}
	copy(hash32[:], hash)
	level, skills, err := instance.GetDiploma(&bind.CallOpts{}, hash32)
	if err != nil {
		return 0, []FtDiplomaSkill{}, err
	}
	if level == 0 {
		return 0, []FtDiplomaSkill{}, fmt.Errorf("the diploma doesnt exist")
	}
	return level, skills, nil
}

func (bc BlockchainImpl) CallGetAllDiploma() ([]FtDiplomaDiplomas, error) {
	instance, _, err := bc.connectEthGetInstance()
	if err != nil {
		return nil, err
	}
	result, err := instance.GetAllDiploma(&bind.CallOpts{From: account.Accounts.GetSign().Address})
	if err != nil {
		return nil, err
	}
	log.Print(result)
	return result, nil
}
