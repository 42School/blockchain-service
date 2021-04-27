package async

import (
	"context"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"time"
)

type checkHashDB struct {
	Id          uuid.UUID
	Tx          []byte
	StudentHash []byte
	SendTime    time.Time
}

func restoreRetryQueue() {
	cursor, err := tools.Db.FindRetry()
	if err != nil {
		tools.LogsError(err)
		return
	}
	for cursor.Next(context.TODO()) {
		var dp diplomas.DiplomaImpl
		err = cursor.Decode(&dp)
		if err != nil {
			continue
		}
		tools.RetryQueue.PushBack(dp)
	}
	metrics.GaugeRetryQueue.Set(float64(tools.RetryQueue.Len()))
}

func restoreCheckQueue() {
	cursor, err := tools.Db.FindCheck()
	if err != nil {
		tools.LogsError(err)
		return
	}
	for cursor.Next(context.TODO()) {
		var toGet checkHashDB
		var toCheck diplomas.VerificationHash
		var tx types.Transaction
		err = cursor.Decode(&toGet)
		if err != nil {
			continue
		}
		toCheck.StudentHash = toGet.StudentHash
		err = tx.UnmarshalJSON(toGet.Tx)
		toCheck.Tx = &tx
		toCheck.SendTime = toGet.SendTime
		if err != nil {
			continue
		}
		tools.ToCheckHash.PushBack(toCheck)
	}
	metrics.GaugeCheckQueue.Set(float64(tools.ToCheckHash.Len()))
}

func RestoreQueue() {
	restoreRetryQueue()
	restoreCheckQueue()
}
