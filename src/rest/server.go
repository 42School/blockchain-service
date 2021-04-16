package rest

import (
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/rest/handlers"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(metrics.PrometheusMiddleware)
	router.Handle("/metrics", promhttp.Handler())
	router.Methods("POST").Path("/create-diploma").Name("Create").HandlerFunc(handlers.CreateDiploma)
	router.Methods("POST").Path("/get-diploma").Name("Get").HandlerFunc(handlers.GetDiploma)
	router.Methods("GET").Path("/get-all-diploma").Name("GetAll").HandlerFunc(handlers.GetAllDiplomas)
	// Testing Route
	if tools.Env == "dev" || tools.Env == "DEV" || tools.Env == "Dev" {
		router.Methods("POST").Path("/check-request").Name("Check").HandlerFunc(handlers.CheckRouter)
	}
	return router
}
