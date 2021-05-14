package account

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ac_1 := Account{, "password"}
var ac_2 := Account{, "password"}
var ac_3 := Account{, "password"}
var account_test = AccountsImpl{KeyStore: &keystore.KeyStore{}, CurrentAccount: 0, Accounts: []Account{ac_1, ac_2, ac_3}}

func TestAccountsImpl_ChangeWriter(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, account_test.CurrentAccount, "Value of Current Account at start")
	account_test.ChangeWriter()
	a.Equal(1, account_test.CurrentAccount, "Value of Current Account doesn't equal at 1")
	account_test.ChangeWriter()
	a.Equal(2, account_test.CurrentAccount, "Value of Current Account doesn't equal at 2")
	account_test.ChangeWriter()
	a.Equal(0, account_test.CurrentAccount, "Value of Current Account doesn't equal at 0")
	account_test.ChangeWriter()
	a.Equal(1, account_test.CurrentAccount, "Value of Current Account doesn't equal at 1")
}

func TestAccountsImpl_GetLenAccounts(t *testing.T) {
	a := assert.New(t)
	length := account_test.GetLenAccounts()
	a.Equal(3, length, "The Accounts length doesn't equal to 3")
}

func TestAccountsImpl_GetWriterByI(t *testing.T) {
	a := assert.New(t)
	address, pk, error := account_test.GetWriterByI(0)
	a.Equal(common.Address{}, address, "")
	a.Equal(ecdsa.PrivateKey{}, pk, "")
	a.Equal(nil, error)
}