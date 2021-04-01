package main

import (
	"context"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/async"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/rest"
	"github.com/42School/blockchain-service/src/tools"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"os"
)

func init()  {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{ForceColors: true, TimestampFormat : "2006-01-02 15:04:05", FullTimestamp:true, PadLevelText: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	if tools.Env == "dev" || tools.Env == "DEV" || tools.Env == "Dev" {
		log.SetLevel(log.DebugLevel)
	}
}

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
	go async.CheckHash()
	go async.RetryDiploma()
	err := MongoStart()
	if err != nil {
		log.WithError(err).Fatal("Mongo Connection Failed.")
		return
	}
	err = api.InitApi()
	if err != nil {
		log.WithError(err).Fatal("Intra.42 API Connection Failed.")
		return
	}
	async.RestoreQueue()
	account.CreateAccountsManager()
	router := rest.InitRouter()
	metrics.RecordMetrics()
	log.Info("Blockchain Service is running !")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.WithError(err).Fatal("HTTP Server doesn't running.")
	}


}
