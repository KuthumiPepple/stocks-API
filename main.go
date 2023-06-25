package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kuthumipepple/stocks-api/config"
	"github.com/kuthumipepple/stocks-api/router"
)

func main() {
	config.ConnectDatabase()
	r := router.Router()
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))

}
