package main

import (
	"context"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/async"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/rest"
	"github.com/42School/blockchain-service/src/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

func MongoStart() error {
	url := "mongodb://" + tools.MongoIp + ":" + tools.MongoPort
	credential := options.Credential{
		Username: tools.MongoUser,
		Password: tools.MongoPasswd,
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url).SetAuth(credential))
	if err != err {
		return err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}
	tools.RetryDB = client.Database("queue").Collection("retry")
	tools.ToCheckDB = client.Database("queue").Collection("valide")
	return nil
}

func main() {
	go async.ValideHash()
	go async.RetryDiploma()
	go async.ReadStdin()
	err := MongoStart()
	if err != nil {
		tools.LogsError(err)
		return
	}
	err = api.InitApi()
	if err != nil {
		tools.LogsError(err)
		return
	}
	async.RestoreQueue()
	account.CreateAccountsManager()
	router := rest.InitRouter()
	tools.LogsMsg("Blockchain Service is running !")
	log.Fatal(http.ListenAndServe(":8080", router))
}
