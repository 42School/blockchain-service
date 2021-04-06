package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

var (
	/*
	***	Metrics for middleware
	 */
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "blockchain_service",
		Name: "http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"url", "code"})

	validationBlockDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "blockchain_service",
		Name: "block_eth_duration_validation_minutes",
		Help: "Duration of the validation block Ethereum (in minutes).",
		Buckets: prometheus.LinearBuckets(1, 2, 15),
	})

	NumberOfRetryDiploma = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "blockchain_service",
		Name: "number_of_retry_per_diploma",
		Help: "Histogram with number of retry per diploma.",
		Buckets: prometheus.LinearBuckets(0, 1, 10),
	})

	/*
	***	Gauge metrics
	 */
	GaugeBalanceWallet = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "blockchain_service",
		Name: "gauge_balance_of_wallet",
		Help: "The level of each wallets",
	}, []string{"address"})

	GaugeRetryQueue = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "blockchain_service",
		Name: "number_of_diploma_in_retry_queue",
		Help: "The total number of diploma in the retry queue.",
	})

	GaugeCheckQueue = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "blockchain_service",
		Name: "number_of_diploma_in_check_queue",
		Help: "The total number of diploma in the check queue.",
	})

	/*
	***	Counter metrics
	 */
	CounterDiplomaSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "blockchain_service",
		Name: "number_of_diploma_success",
		Help: "The total number of diploma completely written on Ethereum.",
	})
)
