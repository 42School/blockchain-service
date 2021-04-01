package tools

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
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

	/*
	***	Other Metrics
	 */
	NmbOfRetryQueue = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "blockchain_service",
		Name: "number_of_diploma_in_retry_queue",
		Help: "The total number of diploma in the retry queue.",
	})

	NmbOfCheckQueue = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "blockchain_service",
		Name: "number_of_diploma_in_check_queue",
		Help: "The total number of diploma in the check queue.",
	})
)

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