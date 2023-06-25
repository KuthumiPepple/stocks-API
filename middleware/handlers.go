package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kuthumipepple/stocks-api/config"
)

func GetStock(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE")
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 0)
	if err != nil {
		log.Fatalf("unable to convert from string to int. %v", err)
	}

	db := config.GetDB()
	stock, err := db.GetStock(id)
	if err != nil {
		log.Fatalf("unable to get stock. %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}
