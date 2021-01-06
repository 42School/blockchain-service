package main

import (
	"context"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/async"
	"github.com/42School/blockchain-service/src/rest"
	"github.com/42School/blockchain-service/src/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)



func MongoStart() error {
	credential := options.Credential{
		Username: "root",
		Password: "example",
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))
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
	tools.LogsMsg("Blockchain Service is running !")
	account.CreateAccountsManager()
	router := rest.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
