package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kuthumipepple/stocks-api/config"
	"github.com/kuthumipepple/stocks-api/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

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

func CreateStock(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		return
	}

	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("unable to decode the request body. %v", err)
	}

	db := config.GetDB()
	newID := db.InsertStock(stock)
	res := response{
		ID:      newID,
		Message: "stock created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	stocks := db.GetAllStocks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 0)
	if err != nil {
		log.Fatalf("unable to convert from string to int. %v", err)
	}

	var stockUpdate models.Stock
	json.NewDecoder(r.Body).Decode(&stockUpdate)
	if err != nil {
		log.Fatalf("unable to decode the request body. %v", err)
	}

	updatedRows := config.GetDB().UpdateStock(id, stockUpdate)
	msg := fmt.Sprintf("stock updated successfully. Total records affected: %v", updatedRows)
	res := response{
		ID:      id,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 0)
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	db := config.GetDB()
	rowsDeleted := db.DeleteStock(id)
	msg := fmt.Sprintf("Stock deleted. Total records affected: %v", rowsDeleted)
	res := response{
		ID:      id,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
