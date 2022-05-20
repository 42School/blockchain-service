package db

import "go.mongodb.org/mongo-driver/mongo"

type Database interface {
	Init() error
	InsertOneRetry(dp interface{})
	InsertOneCheck(check interface{})
	FindOneRetry(filter interface{}) *mongo.SingleResult
	FindOneCheck(filter interface{}) *mongo.SingleResult
	FindRetry() (*mongo.Cursor, error)
	FindCheck() (*mongo.Cursor, error)
	DeleteOneRetry(dp interface{})
	DeleteOneCheck(check interface{})
}

func NewDatabase() Database {
	var i Database
	impl := DatabaseImpl{}
	err := impl.Init()
	if err != nil {
		return i
	}
	i = &impl
	return i
}