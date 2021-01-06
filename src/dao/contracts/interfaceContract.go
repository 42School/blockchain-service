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
	"log"
	"strings"

	"math/big"
)

func connectEthGetInstance() (*Diploma, *ethclient.Client, error) {
	client, errConnection := ethclient.Dial(tools.NetworkLink)
	if errConnection != nil {
		return nil, nil, errConnection
	}
	addressOfAddress := common.HexToAddress(tools.AddressOfContract)
	instance, errInstance := NewDiploma(addressOfAddress, client)
	if errInstance != nil {
		return nil, nil, errInstance
	}
	return instance, client, nil
}

func getAuth() (*bind.TransactOpts, error) {
	client, errConnection := ethclient.Dial(tools.NetworkLink)
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
	auth.GasLimit = uint64(604758)
	auth.GasPrice = gasPrice
	return auth, nil
}

func getLogs(client *ethclient.Client) (logs []types.Log, contractAbi abi.ABI, error error){
	query := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(tools.AddressOfContract)}}
	logs, errLogs := client.FilterLogs(context.Background(), query)
	if errLogs != nil {
		tools.LogsError(errLogs)
		return nil, abi.ABI{}, errLogs
	}
	contractAbi, errAbi := abi.JSON(strings.NewReader(DiplomaABI))
	if errAbi != nil {
		tools.LogsError(errAbi)
		return nil, abi.ABI{}, errAbi
	}
	return logs, contractAbi, nil
}

func GetRevert(client *ethclient.Client, tx *types.Transaction, receipt *types.Receipt) string {
	address, _, _ := account.GetWriterAccount()
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
		errorSig            = []byte{0x08, 0xc3, 0x79, 0xa0}
		abiString, _        = abi.NewType("string", "", nil)
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

func CheckSecurity(client *ethclient.Client, tx *types.Transaction, hash []byte) bool {
	logs, contractAbi, errLogs := getLogs(client)
	if errLogs != nil {
		log.Println(errLogs)
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
				tools.SendMail("Security Alert", "")
				tools.SecuritySystem = true
				return false
			}
		}
	}
	return true
}

func CallCreateDiploma(level uint64, skills [30]uint64, v uint8, r [32]byte, s [32]byte, hash [32]byte) (*types.Transaction, bool) {
	instance, _, err := connectEthGetInstance()
	auth, errAuth := getAuth()
	if err != nil || errAuth != nil {
		return nil, false
	}
	tx, errCreate := instance.CreateDiploma(auth, level, skills, v, r, s, hash)
	if errCreate != nil {
		tools.LogsError(errCreate)
		if strings.Contains(errCreate.Error(), "insufficient funds for gas * price + value") {
			account.ChangeAccount()
		}
		return nil, false
	}
	tools.LogsDev("Transation Hash: " + tx.Hash().Hex())
	return tx, true
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

func CallGetAllDiploma() ([]FtDiplomaDiploma, error) {
	instance, _, err := connectEthGetInstance()
	if err != nil {
		return nil, err
	}
	result, err := instance.GetAllDiploma(&bind.CallOpts{From: account.GetSignAccount().Address})
	if err != nil {
		return nil, err
	}
	log.Print(result)
	return result, nil
}