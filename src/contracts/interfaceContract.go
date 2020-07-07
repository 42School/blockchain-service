package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

var networkLink string = os.Getenv("NETWORKLINK")
var addressOfContract string = os.Getenv("ADDRESSCONTRACT")

func connectEthGetInstance() (*Diploma, error) {
	client, errConnection := ethclient.Dial(networkLink)
	if errConnection != nil {
		return nil, errConnection
	}
	log.Println(client)
	addressOfAddress := common.HexToAddress(addressOfContract)
	instance, errInstance := NewDiploma(addressOfAddress, client)
	if errInstance != nil {
		return nil, errInstance
	}
	log.Println("after get instance", instance)
	return instance, nil
}

func CallCreateDiploma() bool {
	instance, err := connectEthGetInstance()
	if err != nil {
		return false
	}
	log.Println(instance)
	return true
}

func CallGetDiploma(hash []byte) bool {
	instance, err := connectEthGetInstance()
	if err != nil {
		return false
	}
	log.Println(instance)
	hash32 := [32]byte{}
	copy(hash32[:], hash)
	result, errGet := instance.GetDiploma(&bind.CallOpts{}, hash32)
	if errGet != nil {
		return false
	}
	log.Print(result)
	return true
}