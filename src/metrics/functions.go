package metrics

import (
	"context"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path, strconv.Itoa(rw.statusCode)))
		timer.ObserveDuration()
	})
}

func prometheusGaugeWallets() {
	for {
		for i := 0; i < len(account.Accounts); i++ {
			keyjson, err := ioutil.ReadFile(tools.PathKeyStore + "/" + account.Accounts[i].KeyStoreFile)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Error("Metrics Records failed - prometheusGaugeWallets.ReadFile (keystore)")
				continue
			}
			key, err := keystore.DecryptKey(keyjson, account.Accounts[i].Password)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Error("Metrics Records failed - prometheusGaugeWallets.DecryptKey")
				continue
			}
			balance, err := contracts.GetBalance(key.Address)
			if err != nil {
				log.WithFields(log.Fields{"address": key.Address.String(), "error": err}).Error("Metrics Records failed - prometheusGaugeWallets.getBalance")
				continue
			}
			GaugeBalanceWallet.WithLabelValues(key.Address.String()).Set(float64(balance))
		}
		time.Sleep(10 * time.Minute)
	}
}

func PrometheusBlockDuration(blockHash common.Hash, txSendTime time.Time) {
	client, _ := ethclient.Dial(tools.NetworkLink)
	block, err := client.BlockByHash(context.TODO(), blockHash)
	if err != nil {
		return
	}
	metricsTime := block.ReceivedAt.Minute() - txSendTime.Minute()
	validationBlockDuration.Observe(float64(metricsTime))
}

func RecordMetrics() {
	go prometheusGaugeWallets()
}