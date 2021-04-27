package db

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockDatabaseImpl struct {
	mock.Mock
}

func (db *MockDatabaseImpl) Init() error {
	args := db.Called()
	return args.Error(0)
}

func (db *MockDatabaseImpl) InsertOneRetry(dp interface{}) {
	db.Called()
}

func (db *MockDatabaseImpl) InsertOneCheck(check interface{}) {
	db.Called()
}

func (db *MockDatabaseImpl) FindOneRetry(filter interface{}) *mongo.SingleResult {
	args := db.Called(filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (db *MockDatabaseImpl) FindOneCheck(filter interface{}) *mongo.SingleResult {
	args := db.Called(filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (db *MockDatabaseImpl) FindRetry() (*mongo.Cursor, error) {
	args := db.Called()
	return args.Get(0).(*mongo.Cursor), args.Error(0)
}

func (db *MockDatabaseImpl) FindCheck() (*mongo.Cursor, error) {
	args := db.Called()
	return args.Get(0).(*mongo.Cursor), args.Error(0)
}

func (db *MockDatabaseImpl) DeleteOneRetry(dp interface{})  {
	db.Called(dp)
}

func (db *MockDatabaseImpl) DeleteOneCheck(check interface{})  {
	db.Called(check)
}