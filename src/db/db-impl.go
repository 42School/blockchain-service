package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

// Mongo Variable
var (
	MongoIp = os.Getenv("MONGO_IP")
	MongoPort = os.Getenv("MONGO_PORT")
	MongoUser = os.Getenv("MONGO_USER")
	MongoPasswd = os.Getenv("MONGO_PASSWD")
)

type DatabaseImpl struct {
	RetryDB *mongo.Collection
	ToCheckDB *mongo.Collection
}

// Init connect the blockchain-service with a MongoDB and create 1 database 'queue' with 2 table 'retry' & 'check'
func (db *DatabaseImpl) Init() error {
	url := "mongodb://" + MongoIp + ":" + MongoPort
	credential := options.Credential{
		Username: MongoUser,
		Password: MongoPasswd,
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url).SetAuth(credential))
	if err != err {
		return err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}
	db.RetryDB = client.Database("queue").Collection("retry")
	db.ToCheckDB = client.Database("queue").Collection("check")
	return nil
}

func (db *DatabaseImpl) InsertOneRetry(dp interface{}) {
	db.RetryDB.InsertOne(context.TODO(), dp)
}

func (db *DatabaseImpl) InsertOneCheck(check interface{}) {
	db.ToCheckDB.InsertOne(context.TODO(), check)
}

func (db *DatabaseImpl) FindOneRetry(filter interface{}) *mongo.SingleResult {
	result := db.RetryDB.FindOne(context.TODO(), filter)
	return result
}

func (db *DatabaseImpl) FindOneCheck(filter interface{}) *mongo.SingleResult {
	result := db.ToCheckDB.FindOne(context.TODO(), filter)
	return result
}

func (db *DatabaseImpl) FindRetry() (*mongo.Cursor, error) {
	cursor, err := db.RetryDB.Find(context.TODO(), bson.M{})
	return cursor, err
}

func (db *DatabaseImpl) FindCheck() (*mongo.Cursor, error) {
	cursor, err := db.ToCheckDB.Find(context.TODO(), bson.M{})
	return cursor, err
}

func (db *DatabaseImpl) DeleteOneRetry(dp interface{}) {
	db.RetryDB.DeleteOne(context.TODO(), dp)
}

func (db *DatabaseImpl) DeleteOneCheck(check interface{}) {
	db.ToCheckDB.DeleteOne(context.TODO(), check)
}