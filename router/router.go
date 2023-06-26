package router

import (
	"github.com/gorilla/mux"
	"github.com/kuthumipepple/stocks-api/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stocks/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET")

	return router
}
