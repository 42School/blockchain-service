package contracts

import (
	"context"
	"fmt"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/global"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"

	"math/big"
)

func connectEthGetInstance() (*Diploma, *ethclient.Client, error) {
	client, errConnection := ethclient.Dial(global.NetworkLink)
	if errConnection != nil {
		return nil, nil, errConnection
	}
	addressOfAddress := common.HexToAddress(global.AddressOfContract)
	instance, errInstance := NewDiploma(addressOfAddress, client)
	if errInstance != nil {
		return nil, nil, errInstance
	}
	return instance, client, nil
}

func getAuth() (*bind.TransactOpts, error) {
	client, errConnection := ethclient.Dial(global.NetworkLink)
	if errConnection != nil {
		return nil, errConnection
	}
	address, privateKey, errGet := account.GetWriterAccount()
	if errGet != nil {
		return nil, errGet
	}
	nonce, errNonce := client.PendingNonceAt(context.Background(), address)
	if errNonce != nil  {
		return nil, errNonce
	}
	gasPrice, errGas := client.SuggestGasPrice(context.Background())
	if errGas != nil {
		return nil, errGas
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(600000)
	auth.GasPrice = gasPrice
	return auth, nil
}

func getLogs(client *ethclient.Client) (logs []types.Log, contractAbi abi.ABI, error error){
	query := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(global.AddressOfContract)}}
	logs, errLogs := client.FilterLogs(context.Background(), query)
	if errLogs != nil {
		tools.LogsError(errLogs)
		return nil, abi.ABI{}, errLogs
	}
	contractAbi, errAbi := abi.JSON(strings.NewReader(string(DiplomaABI)))
	if errAbi != nil {
		tools.LogsError(errAbi)
		return nil, abi.ABI{}, errAbi
	}
	return logs, contractAbi, nil
}

func CallCreateDiploma(level uint64, skills [30]uint64, v uint8, r [32]byte, s [32]byte, hash [32]byte) bool {
	instance, client, err := connectEthGetInstance()
	auth, errAuth := getAuth()
	if err != nil || errAuth != nil {
		return false
	}
	tx, errCreate := instance.CreateDiploma(auth, level, skills, v, r, s, hash)
	if errCreate != nil {
		tools.LogsError(errCreate)
		if strings.Contains(errCreate.Error(), "FtDiploma: The diploma already exists.") {
			return true
		}
		if strings.Contains(errCreate.Error(), "sender doesn't have enough funds to send tx.") {
			account.ChangeAccount()
		}
		return false
	}
	logs, contractAbi, errLogs := getLogs(client)
	if errLogs != nil {
		tools.LogsError(errLogs)
		return true
	}
	for _, vLog := range logs {
		if vLog.TxHash.Hex() == tx.Hash().Hex() {
			event := struct {
				Student   [32]byte
			}{}
			errUnpack := contractAbi.Unpack(&event, "CreateDiploma", vLog.Data)
			if errUnpack != nil {
				tools.LogsError(errUnpack)
				return true
			}
			if common.Bytes2Hex(hash[:]) != common.Bytes2Hex(event.Student[:]) {
				tools.LogsMsg("Error: The hash writing in blockchain is not the same of this student !")
				tools.SendMail("Security Alert", "bocal@email", "")
				global.SecuritySystem = true
			}
		}
	}
	return true
}

func CallGetDiploma(hash []byte) (uint64, [30]uint64, error) {
	instance, _, err := connectEthGetInstance()
	if err != nil {
		return 0, [30]uint64{}, err
	}
	hash32 := [32]byte{}
	copy(hash32[:], hash)
	result, errGet := instance.GetDiploma(&bind.CallOpts{}, hash32)
	if errGet != nil {
		return 0, [30]uint64{}, errGet
	}
	if result.Level == 0 {
		return 0, [30]uint64{}, fmt.Errorf("the diploma doesnt exist")
	}
	return result.Level, result.Skills, nil
}