package api

import (
	"controllers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/diplomas").Name("Create").HandlerFunc(controllers.DiplomasCreate)
}
