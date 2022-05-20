package account

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

type AccountsManager interface {
	InitAccounts()
	GetLenAccounts() int
	ChangeWriter()
	GetSign() accounts.Account
	SignHash(hash common.Hash) ([]byte, error)
	GetWriter() (common.Address, *ecdsa.PrivateKey, error)
	GetWriterByI(i int) (common.Address, *ecdsa.PrivateKey, error)
}

func NewAccountsManager() AccountsManager {
	var i AccountsManager
	impl := new(AccountsImpl)
	impl.InitAccounts()
	i = impl
	return i
}