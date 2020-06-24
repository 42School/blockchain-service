package api

import (
	"github.com/gorilla/mux"
	controllers "github.com/lpieri/42-Diploma/src/api/controllers"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/diplomas").Name("Create").HandlerFunc(controllers.Print)
	return router
}
