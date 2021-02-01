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

func getAuth() (*bind.TransactOpts, error) {
	client, err := ethclient.Dial(tools.NetworkLink)
	if err != nil {
		return nil, err
	}
	address, privateKey, err := account.GetWriterAccount()
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil  {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(604758)
	auth.GasPrice = gasPrice
	return auth, nil
}

func getLogs(client *ethclient.Client) ([]types.Log, abi.ABI, error) {
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
	logs, contractAbi, err := getLogs(client)
	if err != nil {
		log.Println(err)
		return true
	}
	for _, vLog := range logs {
		if vLog.TxHash.Hex() == tx.Hash().Hex() {
			event := struct {
				Student   [32]byte
			}{}
			_, err := contractAbi.Unpack("CreateDiploma", vLog.Data)
			if err != nil {
				tools.LogsError(err)
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
	if err != nil {
		tools.LogsError(err)
		return nil, false
	}
	auth, err := getAuth()
	if err != nil {
		tools.LogsError(err)
		return nil, false
	}
	tx, err := instance.CreateDiploma(auth, level, skills, v, r, s, hash)
	if err != nil {
		tools.LogsError(err)
		if strings.Contains(err.Error(), "insufficient funds for gas * price + value") {
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
	result, err := instance.GetDiploma(&bind.CallOpts{}, hash32)
	if err != nil {
		return 0, [30]uint64{}, err
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