package main

import (
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/async"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/db"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/rest"
	"github.com/42School/blockchain-service/src/tools"
	log "github.com/sirupsen/logrus"
	"os"
)

func init()  {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, TimestampFormat : "2006-01-02 15:04:05", FullTimestamp:true, PadLevelText: true})
	log.SetOutput(os.Stdout)
	if tools.Env == "dev" || tools.Env == "DEV" || tools.Env == "Dev" {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	tools.Db = db.NewDatabase()
	if tools.Db == nil {
		log.Fatal("Mongo Connection Failed.")
		return
	}
	err := api.InitApi()
	if err != nil {
		log.WithError(err).Fatal("Intra.42 API Connection Failed.")
		return
	}
	async.RestoreQueue()
	account.Accounts = account.NewAccountsManager()
	metrics.RecordMetrics()
	go async.CheckHash()
	go async.RetryDiploma()
	server := rest.NewServer()
	err = server.ListenAndServe()
	if err != nil {
		log.WithError(err).Fatal("HTTP Server doesn't running.")
	}
}
