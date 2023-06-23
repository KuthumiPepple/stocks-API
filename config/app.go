package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kuthumipepple/stocks-api/models"
	_ "github.com/lib/pq"
)

var database *models.StocksDB

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	database = &models.StocksDB{DB: db}
	fmt.Println("Successfully connected to postgres")

}

func GetDB() *models.StocksDB {
	return database
}
