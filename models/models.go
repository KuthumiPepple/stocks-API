package models

import (
	"database/sql"
	"fmt"
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
