package api

import (
	"github.com/gorilla/mux"
)

// control := controllers.new()

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.Methods("GET").Path("/diplomas").Name("Create").HandlerFunc(control.test)
	return router
}
