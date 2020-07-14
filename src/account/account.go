package account

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/lpieri/42-Diploma/src/global"
	"io/ioutil"
)

var AccountsManager *accounts.Manager
var KeyStore *keystore.KeyStore

func CreateAccountsManager() {
	KeyStore = keystore.NewKeyStore(global.PathKeyStore, keystore.StandardScryptN, keystore.StandardScryptP)
	AccountsManager = accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, KeyStore)
}

func GetAccount() accounts.Account {
	return KeyStore.Accounts()[0]
}

func GetKey() (*keystore.Key, error) {
	keyjson, errRead := ioutil.ReadFile(global.PathKeyStore + "/" + global.FileKeyStore)
	if errRead != nil {
		return nil, errRead
	}
	key, errDecrypt := keystore.DecryptKey(keyjson, global.PasswordAccount)
	if errDecrypt != nil {
		return nil, errDecrypt
	}
	if key.Address != global.OfficialAddress {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, global.OfficialAddress)
	}
	return key, nil
}