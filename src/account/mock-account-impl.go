package account

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

type MockAccountsImpl struct {
	mock.Mock
}

func (m *MockAccountsImpl) InitAccounts() {
	m.Called()
}

func (m *MockAccountsImpl) GetLenAccounts() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockAccountsImpl) ChangeWriter() {
	m.Called()
}

func (m *MockAccountsImpl) GetSign() accounts.Account {
	args := m.Called()
	return args.Get(0).(accounts.Account)
}

func (m *MockAccountsImpl) SignHash(hash common.Hash) ([]byte, error) {
	args := m.Called(hash)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockAccountsImpl) GetWriter() (common.Address, *ecdsa.PrivateKey, error) {
	args := m.Called()
	return args.Get(0).(common.Address), args.Get(1).(*ecdsa.PrivateKey), args.Error(1)
}

func (m *MockAccountsImpl) GetWriterByI(i int) (common.Address, *ecdsa.PrivateKey, error) {
	args := m.Called()
	return args.Get(0).(common.Address), args.Get(1).(*ecdsa.PrivateKey), args.Error(1)
}

