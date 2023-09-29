package main

import (
	"log"

	"github.com.br/Leodf/walletcore/balances/internal/app/config"
	"github.com.br/Leodf/walletcore/balances/internal/database/connection"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Coudn't load .env file")
	}
	db := connection.InitDatabase()
	defer db.Close()
	config.InitApp()
}
