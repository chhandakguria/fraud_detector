package router

import (
	"github.com/chhandakguria/fraud_detector/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transaction", handlers.CreateTransaction).Methods("POST")
	return r
}
