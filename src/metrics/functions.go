package metrics

import (
	"context"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
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
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path, r.Method, strconv.Itoa(rw.statusCode)))
		timer.ObserveDuration()
	})
}

func prometheusGaugeWallets() {
	bc := contracts.NewBlockchainFunc()
	for {
		for i := 0; i < account.Accounts.GetLenAccounts(); i++ {
			address, _, err := account.Accounts.GetWriterByI(i)
			balance, err := bc.GetBalance(address)
			if err != nil {
				log.WithFields(log.Fields{"address": address.String(), "error": err}).Error("Metrics Records failed - prometheusGaugeWallets.getBalance")
				continue
			}
			GaugeBalanceWallet.WithLabelValues(address.String()).Set(float64(balance))
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
