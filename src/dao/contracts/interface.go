package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainFunc interface {
	GetBalance(address common.Address) (int64, error)
	GetRevert(client *ethclient.Client, tx *types.Transaction, receipt *types.Receipt) string
	CheckSecurity(client *ethclient.Client, tx *types.Transaction, hash []byte) bool
	CallCreateDiploma(level uint64, skills [30]uint64, skillsSlugs [30]string, v uint8, r [32]byte, s [32]byte, hash [32]byte) (*types.Transaction, bool)
	CallGetDiploma(hash []byte) (uint64, []FtDiplomaSkill, error)
	CallGetAllDiploma() ([]FtDiplomaDiplomas, error)
}

func NewBlockchainFunc() BlockchainFunc {
	var i BlockchainFunc
	i = &BlockchainImpl{}
	return i
}