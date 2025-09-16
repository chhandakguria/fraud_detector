package router

import (
	"github.com/gorilla/mux"
	"fraud-detector/handlers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transaction", handlers.CreateTransaction).Methods("POST")
	return r
}
