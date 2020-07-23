package account

import (
	"crypto/ecdsa"
	"github.com/42School/blockchain-service/src/global"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
)

type Account struct {
	KeyStoreFile	string
	Password		string
}

var AccountsManager *accounts.Manager
var KeyStore *keystore.KeyStore
var CurrentAccount int = 0
var Accounts = []Account{{"UTC--2020-07-16T13-52-10.535505000Z--7e12234e994384a757e2689addb2a463ccd3b47d", "password"}, {"File2", "pwd2"}}


func CreateAccountsManager() {
	KeyStore = keystore.NewKeyStore(global.PathKeyStore, keystore.StandardScryptN, keystore.StandardScryptP)
	AccountsManager = accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, KeyStore)
}

func GetAccount() accounts.Account {
	return KeyStore.Accounts()[0]
}

func GetWriterAccount() (common.Address, *ecdsa.PrivateKey, error) {
	keyjson, errRead := ioutil.ReadFile(global.PathKeyStore + "/" + Accounts[CurrentAccount].KeyStoreFile)
	if errRead != nil {
		return common.Address{}, nil, errRead
	}
	key, errDecrypt := keystore.DecryptKey(keyjson, Accounts[CurrentAccount].Password)
	if errDecrypt != nil {
		return common.Address{}, nil, errDecrypt
	}
	return key.Address, key.PrivateKey, nil
}

func ChangeAccount() {
	// Send Mail of Current Account 
	if CurrentAccount + 1 == len(Accounts) {
		CurrentAccount = 0
	} else {
		CurrentAccount = CurrentAccount + 1
	}
}