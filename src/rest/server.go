package rest

import (
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/rest/handlers"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	server                *http.Server
	handlerCreateDiploma  *handlers.CreateDiplomaHandler
	handlerGetDiploma     *handlers.GetDiplomaHandler
	handlerGetAllDiplomas *handlers.GetAllDiplomaHandler
}

func NewServer() *Server {
	s := &Server{}
	s.handlerCreateDiploma = handlers.NewCreateDiplomaHandler()
	s.handlerGetDiploma = handlers.NewGetDiplomaHandler()
	s.handlerGetAllDiplomas = handlers.NewGetAllDiplomaHandler()
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: s.router(),
	}
	return s
}

func (s *Server) ListenAndServe() error {
	log.Info("Blockchain Service is running !")
	return s.server.ListenAndServe()
}

func (s *Server) router() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(metrics.PrometheusMiddleware)
	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/create-diploma", s.handlerCreateDiploma).Methods(http.MethodPost)
	router.Handle("/get-diploma", s.handlerGetDiploma).Methods(http.MethodPost)
	router.Handle("/get-all-diploma", s.handlerGetAllDiplomas).Methods(http.MethodGet)
	// Testing Route
	if tools.Env == "dev" || tools.Env == "DEV" || tools.Env == "Dev" {
		router.Methods("POST").Path("/check-request").Name("Check").HandlerFunc(handlers.CheckRouter)
	}
	return router
}
