package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	accounts "github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/global"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"

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
	var address common.Address
	var privateKey *ecdsa.PrivateKey
	var nonce uint64
	var errNonce error
	client, errConnection := ethclient.Dial(global.NetworkLink)
	if errConnection != nil {
		return nil, errConnection
	}
	if global.Env == "Dev" {
		address = common.HexToAddress(global.DevAddress)
		privateKey, _ = crypto.HexToECDSA(global.DevPrivateKey)
	} else {
		account := accounts.GetAccount()
		address = account.Address
		ks, errKey := accounts.GetKey()
		if errKey != nil  {
			return nil, errKey
		}
		privateKey = ks.PrivateKey
	}
	nonce, errNonce = client.PendingNonceAt(context.Background(), address)
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

func CallCreateDiploma(level uint64, skills [30]uint64, v uint8, r [32]byte, s [32]byte, hash [32]byte) bool {
	instance, _, err := connectEthGetInstance()
	if err != nil {
		tools.LogsError(err)
		return false
	}
	auth, errAuth := getAuth()
	if errAuth != nil {
		tools.LogsError(errAuth)
		return false
	}
	tx, errCreate := instance.CreateDiploma(auth, level, skills, v, r, s, hash)
	log.Println(tx.Data(), string(tx.Data()))
	if errCreate != nil {
		tools.LogsError(errCreate)
		return false
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