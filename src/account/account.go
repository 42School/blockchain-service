package account

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
)

var pathKeyStore string = os.Getenv("KEYSTORE")
var AccountsManager *accounts.Manager
var KeyStore *keystore.KeyStore

func CreateAccountsManager() {
	KeyStore = keystore.NewKeyStore(pathKeyStore, keystore.StandardScryptN, keystore.StandardScryptP)
	AccountsManager = accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, KeyStore)
}

func GetAccount() accounts.Account {
	return KeyStore.Accounts()[0]
}