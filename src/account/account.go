package account

import (
	"crypto/ecdsa"
	"github.com/42School/blockchain-service/src/global"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

type Account struct {
	KeyStoreFile	string
	Password		string
}

var AccountsManager *accounts.Manager
var KeyStore *keystore.KeyStore
var CurrentAccount int = 0
var Accounts = []Account{
	{"UTC--2020-07-24T08-19-17.983576000Z--cac03bac6965e6d8ca96537a0344cc506b32c2c7", "password"},
	{"UTC--2020-07-24T08-24-31.985849000Z--fe5ac6a7bb66da6916becb74a4a3e00074cd2599", "password"},
	{"UTC--2020-07-24T08-25-31.194883000Z--aec7bdfb241e56c04acf5e1a2a49f147867b85b7", "password"},
}

func CreateAccountsManager() {
	KeyStore = keystore.NewKeyStore(global.PathKeyStoreSign, keystore.StandardScryptN, keystore.StandardScryptP)
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
	log.Println("pk:", hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))
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