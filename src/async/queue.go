package async

import (
	"context"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type verifHashDB struct {
	Id uuid.UUID
	Tx []byte
	StudentHash []byte
	SendTime time.Time
}

func RestoreQueue() {
	cursor, err := tools.RetryDB.Find(context.TODO(), bson.M{})
	if err != nil {
		tools.LogsError(err)
		return
	}
	for cursor.Next(context.TODO()) {
		var dp diplomas.Diploma
		err = cursor.Decode(&dp)
		if err != nil {
			continue
		}
		tools.RetryQueue.PushBack(dp)
	}
	metrics.GaugeRetryQueue.Set(float64(tools.RetryQueue.Len()))
	cursor, err = tools.ToCheckDB.Find(context.TODO(), bson.M{})
	if err != nil {
		tools.LogsError(err)
		return
	}
	for cursor.Next(context.TODO()) {
		var toGet verifHashDB
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
