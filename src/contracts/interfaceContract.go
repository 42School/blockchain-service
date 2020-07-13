package contracts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lpieri/42-Diploma/src/global"

	//accounts "github.com/lpieri/42-Diploma/src/account"
	"log"
	"math/big"
)

func connectEthGetInstance() (*Diploma, *ethclient.Client, error) {
	client, errConnection := ethclient.Dial(global.NetworkLink)
	if errConnection != nil {
		return nil, nil, errConnection
	}
	log.Println(client)
	addressOfAddress := common.HexToAddress(global.AddressOfContract)
	instance, errInstance := NewDiploma(addressOfAddress, client)
	if errInstance != nil {
		return nil, nil, errInstance
	}
	log.Println("after get instance", instance)
	return instance, client, nil
}

func getAuth() (*bind.TransactOpts, error) {
	//log.Println("enter in getauth...")
	client, errConnection := ethclient.Dial(global.NetworkLink)
	if errConnection != nil {
		return nil, errConnection
	}
	//account := accounts.GetAccount()
	//key, errKey := accounts.GetKey()
	//if errKey != nil  {
	//	return nil, errKey
	//}
	address := common.HexToAddress("0xa0A6AC7843BF9E9A1bb86F5391a2BF5fB1Bd8C78")
	pk, _ := crypto.HexToECDSA("10865bedc583c8c667980202e5918e22ae1fb4c98f64cb1ecbe2381cbdcee27f")
	nonce, errNonce := client.PendingNonceAt(context.Background(), address)
	//nonce, errNonce := client.PendingNonceAt(context.Background(), account.Address)
	if errNonce != nil  {
		return nil, errNonce
	}
	gasPrice, errGas := client.SuggestGasPrice(context.Background())
	if errGas != nil {
		return nil, errGas
	}
	//auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth := bind.NewKeyedTransactor(pk)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(600000)
	auth.GasPrice = gasPrice
	return auth, nil
}

func CallCreateDiploma(level uint64, skills [30]uint64, v uint8, r [32]byte, s [32]byte, hash [32]byte) bool {
	instance, _, err := connectEthGetInstance()
	if err != nil {
		return false
	}
	auth, errAuth := getAuth()
	if errAuth != nil {
		//log.Println("auth", errAuth)
		return false
	}
	tx, errCreate := instance.CreateDiploma(auth, level, skills, v, r, s, hash)
	if errCreate != nil {
		log.Println(errCreate)
		return false
	}
	log.Println("tx", tx)
	return true
}

func CallGetDiploma(hash []byte) (uint64, [30]uint64, error) {
	instance, client, err := connectEthGetInstance()
	if err != nil {
		return 0, [30]uint64{}, err
	}
	log.Println(instance, client)
	hash32 := [32]byte{}
	copy(hash32[:], hash)
	result, errGet := instance.GetDiploma(&bind.CallOpts{}, hash32)
	if errGet != nil {
		log.Println(errGet)
		return 0, [30]uint64{}, errGet
	}
	if result.Level == 0 {
		return 0, [30]uint64{}, fmt.Errorf("the diploma doesnt exist")
	}
	log.Print("result of get:", result)
	return result.Level, result.Skills, nil
}