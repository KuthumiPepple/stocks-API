package models

import (
	"database/sql"
	"fmt"
	"log"
)

type StocksDB struct {
	*sql.DB
}

type Stock struct {
	StockID int64  `json:"stockid"`
	Symbol  string `json:"symbol"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}

func (db *StocksDB) GetStock(id int64) (Stock, error) {
	var stock Stock
	sqlQuery := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.QueryRow(sqlQuery, id)
	err := row.Scan(&stock.StockID, &stock.Symbol, &stock.Price, &stock.Company)
	if err == sql.ErrNoRows {
		fmt.Println("no rows were returned")
		return stock, nil
	}
	return stock, err
}

func (db *StocksDB) InsertStock(stock Stock) int64 {
	sqlQuery := `INSERT INTO stocks(symbol, price, company) VALUES($1, $2, $3) RETURNING stockid`
	var id int64
	err := db.QueryRow(sqlQuery, stock.Symbol, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("unable to execute the query. %v", err)
	}
	fmt.Printf("inserted a record: %v\n", id)
	return id
}

func (db *StocksDB) GetAllStocks() []Stock {
	var stocks []Stock
	sqlQuery := `SELECT * FROM stocks`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}
	for rows.Next(){
		var stock Stock
		err = rows.Scan(&stock.StockID, &stock.Symbol, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("unable to scan the row. %v", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks
}
