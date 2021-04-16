package diplomas

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"time"
)

type Skill struct {
	Name  string  `json:"name"`
	Level float64 `json:"level"`
}

type Diploma struct {
	Id         uuid.UUID `bson:"_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BirthDate  string    `json:"birth_date"`
	AlumniDate string    `json:"alumni_date"`
	Level      float64   `json:"level"`
	Skills     []Skill   `json:"skills"`
	Counter    int       `json:"counter"`
}

type VerificationHash struct {
	Id          uuid.UUID `bson:"_id"`
	Tx          *types.Transaction
	StudentHash []byte
	SendTime    time.Time
}
