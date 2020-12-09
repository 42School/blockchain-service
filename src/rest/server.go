package rest

import (
	handlers "github.com/42School/blockchain-service/src/rest/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/create-diploma").Name("Create").HandlerFunc(handlers.CreateDiploma)
	router.Methods("POST").Path("/get-diploma").Name("Get").HandlerFunc(handlers.GetDiploma)
	// Testing Route
	router.Methods("POST").Path("/check-request").Name("Check").HandlerFunc(handlers.CheckRouter)
	return router
}
