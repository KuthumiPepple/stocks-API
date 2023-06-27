package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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
	for rows.Next() {
		var stock Stock
		err = rows.Scan(&stock.StockID, &stock.Symbol, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("unable to scan the row. %v", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks
}

func (db *StocksDB) UpdateStock(id int64, stockUpdate Stock) int64 {
	originalStock, err := db.GetStock(id)
	if err != nil {
		log.Fatalf("unable to get stock. %v", err)
	}

	sqlQuery, params := generateUpdateQuery(id, originalStock, stockUpdate)

	res, err := db.Exec(sqlQuery, params...)
	if err != nil {
		log.Fatalf("unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("error while checking the affected rows. %v", err)
	}
	fmt.Println("number of rows affected:", rowsAffected)
	return rowsAffected
}

func generateUpdateQuery(id int64, originalStock, stockUpdate Stock) (string, []any) {
	var sb strings.Builder
	sb.WriteString("UPDATE stocks SET ")
	var params []any
	if stockUpdate.Symbol != "" {
		sb.WriteString(fmt.Sprintf("symbol=$%d,", len(params)+1))
		params = append(params, stockUpdate.Symbol)
	}
	if stockUpdate.Price != 0 {
		sb.WriteString(fmt.Sprintf("price=$%d,", len(params)+1))
		params = append(params, stockUpdate.Price)
	}
	if stockUpdate.Company != "" {
		sb.WriteString(fmt.Sprintf("company=$%d,", len(params)+1))
		params = append(params, stockUpdate.Company)
	}

	updateQuery := strings.TrimSuffix(sb.String(), ",")
	sb.Reset()
	sb.WriteString(updateQuery)
	sb.WriteString(fmt.Sprintf(" WHERE stockid=$%d", len(params)+1))
	params = append(params, id)
	updateQuery = sb.String()

	return updateQuery, params
}
