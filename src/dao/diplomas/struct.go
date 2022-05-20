package diplomas

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"time"
)

type VerificationHash struct {
	Id          uuid.UUID `bson:"_id"`
	Tx          *types.Transaction
	StudentHash []byte
	SendTime    time.Time
}
