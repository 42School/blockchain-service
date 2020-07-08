package account

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"os"
)

var pathKeyStore string = os.Getenv("KEYSTORE")
var passwordAccount string = os.Getenv("KEYPASSWD")
var addr common.Address = common.HexToAddress("0x8A21Dc0aeC762cD85de81B2bcd396a9d5676cFD7")
var AccountsManager *accounts.Manager
var KeyStore *keystore.KeyStore

func CreateAccountsManager() {
	KeyStore = keystore.NewKeyStore(pathKeyStore, keystore.StandardScryptN, keystore.StandardScryptP)
	AccountsManager = accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, KeyStore)
}

func GetAccount() accounts.Account {
	return KeyStore.Accounts()[0]
}

func GetKey() (*keystore.Key, error) {
	keyjson, errRead := ioutil.ReadFile(pathKeyStore + "/UTC--2020-06-29T11-29-15.108416000Z--8a21dc0aec762cd85de81b2bcd396a9d5676cfd7")
	if errRead != nil {
		return nil, errRead
	}
	key, errDecrypt := keystore.DecryptKey(keyjson, passwordAccount)
	if errDecrypt != nil {
		return nil, errDecrypt
	}
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, addr)
	}
	return key, nil
}